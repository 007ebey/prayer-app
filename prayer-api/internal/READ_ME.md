# Prayer API

Go backend for the Prayer App.

The backend handles:

- Clerk authentication
- Application users
- Roles and permissions
- Prayer groups
- Prayer-group access
- Prayer sessions
- Authorization rules

---

# 1. Architecture

The backend follows this dependency direction:

    HTTP / Clerk
        |
        v
    Application Services
        |
        v
    Domain
        |
        v
    Repository Interfaces
        |
        +---- Memory Repository
        |
        +---- PostgreSQL Repository (planned)

The domain must not depend on Clerk, HTTP, or database implementations.

---

# 2. Current Project Structure

    prayer-api/
    |
    +-- cmd/
    |   +-- api/
    |       +-- main.go
    |
    +-- internal/
    |   |
    |   +-- application/
    |   |   +-- auth/
    |   |       +-- login.go
    |   |       +-- login_test.go
    |   |
    |   +-- auth/
    |   |   +-- identity.go
    |   |   +-- clerk.go
    |   |
    |   +-- config/
    |   |   +-- config.go
    |   |
    |   +-- domain/
    |   |   +-- user/
    |   |   +-- role/
    |   |   +-- prayergroup/
    |   |
    |   +-- http/
    |   |   +-- auth_handler.go
    |   |   +-- auth_handler_test.go
    |   |   +-- router.go
    |   |   +-- router_test.go
    |   |
    |   +-- repository/
    |       +-- memory/
    |
    +-- .env
    +-- go.mod
    +-- go.sum
    +-- README.md

---

# 3. Core Domain Model

## User

A User represents an application user.

A user is NOT the same thing as a Clerk user.

Clerk establishes identity.

The Prayer API owns application-specific information such as:

- status
- roles
- prayer-group access

Conceptually:

    User
    |
    +-- ID
    +-- ExternalID (Clerk user ID)
    +-- Name
    +-- Status
    +-- RoleIDs

Possible status:

    active
    blocked

---

# 4. Authentication

Clerk is the external identity provider.

Authentication flow:

    React App
        |
        | Clerk Session Token
        v
    Prayer API
        |
        v
    Clerk Middleware
        |
        | verified
        v
    Session Claims
        |
        v
    Clerk User ID
        |
        v
    Identity Provider
        |
        v
    Application Login Service

The client must NOT send a user ID and claim that identity.

The authenticated Clerk token determines the identity.

---

# 5. First Login / User Provisioning

When an authenticated Clerk user logs into the Prayer API:

    Clerk Identity
          |
          v
    Find Application User
          |
       +--+--+
       |     |
     Found  Missing
       |     |
       |     v
       |   Create User
       |     |
       |     +-- Status = Active
       |     |
       |     +-- Role = Members
       |     |
       |     +-- Prayer Group = Visitor
       |     
       +--------+
            |
            v
       Login Response

Every new user therefore starts with:

    Role
        Members

    Prayer Group Access
        Visitor

Visitor is the default prayer-group access.

---

# 6. Roles

Role is a separate domain.

A user can have multiple roles.

Example:

    User
    |
    +-- Members
    |
    +-- Host

Roles contain permissions.

Example:

    Members
    |
    +-- View Prayer Sessions
    +-- Join Prayer Sessions

Another role could contain:

    Prayer Host
    |
    +-- View Prayer Sessions
    +-- Join Prayer Sessions
    +-- Host Prayer Sessions

Authorization should depend on permissions rather than hard-coded role names.

Avoid:

    if role == "admin"

Prefer:

    user has permission "manage_users"

This allows new roles to be created without rewriting authorization logic.

---

# 7. Prayer Groups

A user can have access to multiple prayer groups.

Example:

    User
    |
    +-- Visitor
    |
    +-- Youth Prayer
    |
    +-- Leaders Prayer

Prayer-group access is separate from roles.

Roles answer:

    WHAT can the user do?

Prayer groups answer:

    WHERE can the user do it?

Example:

    Role:
        Prayer Host

    Prayer Groups:
        Youth Prayer
        Sunday Prayer

Meaning:

    The user has hosting capabilities,
    but only within groups they can access.

---

# 8. Prayer Group Access State

Prayer-group access should support at least three conceptual states:

## No Association

The user has no explicit access to the prayer group.

## Active

The user has access.

## Blocked

The user is explicitly denied access.

These are intentionally different.

    NONE
      |
      | grant
      v
    ACTIVE
      |
      | block
      v
    BLOCKED
      |
      | unblock
      v
    ACTIVE

Removing access:

    ACTIVE
      |
      | remove
      v
    NONE

Blocked should NOT mean the same thing as removed.

This allows future automatic access rules without accidentally restoring access to explicitly blocked users.

---

# 9. Currently Implemented API

## Health

    GET /api/health

Authentication:

    Not required

Response:

    HTTP 200

Example:

    {
        "status": "ok"
    }

---

## Login

    POST /api/auth/login

Authentication:

    Clerk Bearer token required.

Header:

    Authorization: Bearer <CLERK_SESSION_TOKEN>

Request body:

    None

The user identity comes from the verified Clerk session.

### New User

Response:

    HTTP 201 Created

Example:

    {
        "user": {
            "id": "user_1",
            "name": "Anna Mary",
            "status": "active",
            "roles": [
                "role_members"
            ]
        },
        "prayerGroups": [
            {
                "groupId": "visitor",
                "status": "active"
            }
        ],
        "created": true
    }

### Existing User

Response:

    HTTP 200 OK

Example:

    {
        "user": {
            "id": "user_1",
            "name": "Anna Mary",
            "status": "active",
            "roles": [
                "role_members"
            ]
        },
        "prayerGroups": [
            {
                "groupId": "visitor",
                "status": "active"
            }
        ],
        "created": false
    }

---

# 10. Planned User Access API

These endpoints are NOT implemented yet.

They represent the intended API contract.

## Get User

    GET /api/users/{userID}

Expected result:

    User
    + Roles
    + Prayer Groups

---

## Assign Role

    POST /api/users/{userID}/roles/{roleID}

Example:

    POST /api/users/user_123/roles/role_host

Guard:

    Actor must have permission to manage users/roles.

---

## Remove Role

    DELETE /api/users/{userID}/roles/{roleID}

Guard:

    Actor must have permission to manage users/roles.

---

# 11. Planned Prayer Group Access API

## Grant Access

    POST /api/users/{userID}/prayer-groups/{groupID}

Example:

    POST /api/users/user_123/prayer-groups/youth

State transition:

    NONE -> ACTIVE

---

## Remove Access

    DELETE /api/users/{userID}/prayer-groups/{groupID}

State transition:

    ACTIVE -> NONE

---

## Block Access

    POST /api/users/{userID}/prayer-groups/{groupID}/block

State transition:

    ACTIVE -> BLOCKED

---

## Unblock Access

    POST /api/users/{userID}/prayer-groups/{groupID}/unblock

State transition:

    BLOCKED -> ACTIVE

---

# 12. Planned Prayer Group API

Prayer groups will eventually expose operations such as:

    GET    /api/prayer-groups
    GET    /api/prayer-groups/{groupID}
    POST   /api/prayer-groups
    PUT    /api/prayer-groups/{groupID}
    DELETE /api/prayer-groups/{groupID}

Prayer-group data includes:

    PrayerGroup
    |
    +-- ID
    +-- Name
    +-- Description
    +-- Type
    +-- Status
    +-- Sessions

The frontend Prayer Groups screen should consume these APIs rather than owning prayer-group business rules.

---

# 13. Planned Prayer Sessions

Prayer groups can contain multiple prayer sessions.

Conceptually:

    PrayerGroup
        |
        +-- Session
        |
        +-- Session
        |
        +-- Session

Session management will eventually include:

- upcoming sessions
- active sessions
- joining a session
- ending a session
- hosting permissions
- prayer wall
- active participants

The session domain should remain separate from PrayerGroup.

PrayerGroup should reference session IDs rather than absorb all session behavior.

---

# 14. Authorization Model

Authentication and authorization are different.

Authentication:

    Who are you?

Handled by:

    Clerk

Authorization:

    What are you allowed to do?

Handled by:

    Prayer API Domain

Example protected operation:

    POST /api/users/user_123/roles/role_host
                    |
                    v
            Clerk Authentication
                    |
                    v
               Actor User
                    |
                    v
               Actor Roles
                    |
                    v
              Permissions
                    |
             +------+------+
             |             |
          Allowed        Denied
             |             |
             v             v
          Execute       HTTP 403

Authorization rules must be enforced by the backend.

The React frontend hiding a button is NOT authorization.

---

# 15. HTTP Status Conventions

Use:

    200 OK
        Existing resource/action succeeded.

    201 Created
        New resource created.

    400 Bad Request
        Invalid request.

    401 Unauthorized
        Authentication missing or invalid.

    403 Forbidden
        Authenticated but not permitted.

    404 Not Found
        Requested resource does not exist.

    409 Conflict
        Requested transition conflicts with current state.

    500 Internal Server Error
        Internal application/repository failure.

    502 Bad Gateway
        External identity provider failure.

---

# 16. Current Persistence

Current repositories are IN MEMORY.

Located under:

    internal/repository/memory/

Actual data exists only in process RAM.

Example:

    memory.UserRepository
            |
            v
    map[user.ID]*user.User

Restarting the server destroys this data.

This is intentional during early domain development.

Planned:

    Repository Interface
           |
       +---+---+
       |       |
     Memory  PostgreSQL
     Tests   Production

Application services should not need to change when PostgreSQL is introduced.

---

# 17. Configuration

Local development uses:

    .env

Example:

    PORT=8080
    CLERK_SECRET_KEY=<secret>

Never commit the real `.env`.

Use:

    .env.example

for documenting required variables.

---

# 18. Running the API

From the Go project root:

    go run ./cmd/api

Expected:

    Prayer API listening on :8080

Test health:

    GET http://localhost:8080/api/health

---

# 19. Running Tests

Run everything:

    go test ./...

Verbose:

    go test -v ./...

Application tests:

    go test -v ./internal/application/auth

Domain tests:

    go test -v ./internal/domain/...

HTTP/API tests:

    go test -v ./internal/http

---

# 20. Current Test Coverage

Currently verified behavior includes:

## User

- New user is active
- Default Members role
- Block user
- Unblock user
- Duplicate roles prevented

## Role

- Permissions
- Invalid permission rejection
- Duplicate permission handling

## Prayer Group

- Visitor group
- Active access
- Block access
- Unblock access
- Session replacement

## Login Application Service

- New user creation
- Members role assignment
- Visitor access assignment
- Existing user reused
- Blocked user rejected
- Duplicate Visitor access prevented

## HTTP API

- New login returns 201
- Existing login returns 200
- Missing authentication returns 401
- Identity provider failure returns 502
- Health endpoint returns 200

---

# 21. Current API Status

| API | Status |
|---|---|
| GET /api/health | Implemented |
| POST /api/auth/login | Implemented |
| GET /api/users/{id} | Implemented |
| POST user role | Implemented |
| DELETE user role | Planned |
| POST prayer-group access | Planned |
| DELETE prayer-group access | Planned |
| Block prayer-group access | Planned |
| Unblock prayer-group access | Planned |
| GET prayer groups | Planned |
| GET prayer group | Planned |
| Create prayer group | Planned |
| Update prayer group | Planned |
| Delete prayer group | Planned |
| Prayer session APIs | Planned |

---

# 22. Development Rule

Every new backend capability should follow:

    Requirement
        |
        v
    Domain Rule
        |
        v
    State / Event / Guard
        |
        v
    Application Service
        |
        v
    Repository Contract
        |
        v
    HTTP API
        |
        v
    Tests

Do not start from the HTTP endpoint and push business logic downward.

The HTTP layer is an adapter.

The domain owns the rules.