# Prayer API Architecture and Go Usage

## Why the code is organized this way

The `prayer-api` service uses a clean layered architecture to keep responsibilities separated and to make the codebase easy to understand, test, and extend.

### 1. `cmd/api` is the entry point

- `cmd/api/main.go` is the only package that composes and wires the application.
- It builds concrete implementations for repository interfaces and application services.
- This keeps wiring separate from business logic.

### 2. `internal` enforces boundaries

- The `internal/` directory means packages under it are private to this module.
- External code cannot import `prayer-api/internal/...`, which protects implementation details.
- This is a standard Go convention for encapsulation.

### 3. Layered dependency direction

The project follows a one-way dependency flow:

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
        +---- In-memory repository implementation

This means:

- HTTP only depends on application services, not domain internals.
- Application services depend on abstract repository interfaces, not concrete storage.
- Domain model packages do not depend on HTTP, Clerk, or repository implementations.

### 4. Separation of concerns

- `internal/http/` handles request parsing, response writing, and auth middleware integration.
- `internal/application/` contains use cases and business workflows.
- `internal/domain/` contains entity invariants, validation, and pure business rules.
- `internal/repository/memory/` provides a concrete in-memory database for runtime behavior.

## Key Go features used effectively

### Packages and import aliases

- Imports use named aliases for clarity, e.g.:
  - `appprayergroup "prayer-api/internal/application/prayergroup"`
  - `domainprayergroup "prayer-api/internal/domain/prayergroup"`
- This avoids collisions and makes the source intent clear.

### Interfaces for abstraction and testing

- Services depend on interfaces, not concrete implementations.
- Example:
  - `type UserRepository interface { FindByExternalID(...) ... }`
  - `type PrayerGroupRepository interface { FindByID(...) Save(...) }`
- This makes it easy to swap implementations and to test services in isolation.

### Constructor functions and dependency injection

- Packages expose `New...` constructors.
- Example: `NewCreateService(users, roles, groups, ids)`.
- `main.go` wires objects together explicitly.
- This pattern avoids global state and keeps dependencies visible.

### Context propagation

- Handler and service methods accept `context.Context`.
- Example: `Create(ctx context.Context, command CreateCommand)`.
- This allows request-scoped cancellation, deadlines, and cross-cutting values.

### Error handling

- Go-style error returns are used consistently.
- Domain packages expose sentinel errors such as `ErrUnauthorized` and `ErrPrayerGroupExists`.
- HTTP handlers translate these errors into appropriate status codes.

### Typed constants for domain values

- Domain packages define typed constants:
  - `TypeRegular`, `TypeVisitor` in `internal/domain/prayergroup`
  - `StatusActive`, `StatusBlocked` in `internal/domain/user`
- Typed constants keep values safe and self-documenting.

### Composition through methods

- Domain entities expose methods with behavior:
  - `PrayerGroup.Rename()`
  - `User.AssignRole()`
  - `Role.HasPermission()`
- This keeps business rules close to the entities they govern.

### In-memory repository with synchronization

- `internal/repository/memory` uses `sync.RWMutex` to protect maps.
- This is a simple runnable storage layer suitable for local development and tests.
- It implements the repository interfaces used by application services.

### HTTP routing and handler adaptation

- `internal/http/router.go` creates an `http.Handler` using `http.NewServeMux()`.
- It wires routes to handler methods and applies Clerk auth middleware.
- Handler packages adapt HTTP requests into application commands and responses.

## How this organization helps

- **Testability**: application logic can be tested independently of HTTP and DB.
- **Maintainability**: changes in one layer do not ripple across unrelated layers.
- **Clarity**: each package has a focused responsibility.
- **Extensibility**: new repository implementations can be added without changing services.
- **Go idioms**: the code leverages common Go patterns like packages, interfaces, constructors, typed constants, and `context`.

## Practical examples from the codebase

- `cmd/api/main.go`
  - Wires `memory.NewUserRepository()` and `appauth.NewService(...)`.
  - Creates `httpapi.NewRouter(...)` and starts the HTTP server.

- `internal/application/prayergroup/create.go`
  - Implements a create use case as a service.
  - Validates authorization, checks uniqueness, instantiates the domain object, and persists it.

- `internal/http/prayer_group_handler.go`
  - Decodes JSON request bodies.
  - Calls `CreateService.Create(...)`.
  - Converts service errors into HTTP responses.

- `internal/domain/prayergroup/prayer_group.go`
  - Encapsulates entity invariants and state transitions.
  - Keeps business rules in the domain layer.

- `internal/repository/memory/prayer_group_repository.go`
  - Implements repository behavior with in-memory storage.
  - Uses mutexes to protect concurrent access.

## Recommended Go workflows

- Run `go fmt ./...` to keep formatting consistent.
- Run `go test ./...` to validate packages and catch compile-time issues.
- Keep `internal/` package boundaries strict to preserve architecture.
