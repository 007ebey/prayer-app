# Prayer API Integration Guide

This guide explains how to integrate the **prayer-app** (React frontend) with the **prayer-api** (Go backend) service.

---

## Table of Contents

1. [Overview](#overview)
2. [Authentication](#authentication)
3. [API Endpoints](#api-endpoints)
4. [Implementation Steps](#implementation-steps)
5. [Error Handling](#error-handling)
6. [Development and Testing](#development-and-testing)

---

## Overview

The Prayer API is a Go backend service that manages:

- **User Management**: User profiles, statuses, and permissions
- **Role-Based Access Control**: Assigning and managing user roles with permissions
- **Prayer Groups**: Creating, updating, and managing prayer groups
- **Group Access**: Assigning users to prayer groups and blocking access
- **Authentication**: Integration with Clerk for identity verification

### Architecture

The API follows a layered architecture:

```
HTTP / Clerk Authentication
    ↓
Application Services (Business Logic)
    ↓
Domain Model (Core Rules)
    ↓
Repository (Data Persistence)
```

All requests are authenticated via **Clerk** and require a valid session token in the `Authorization` header.

---

## Authentication

### Clerk Setup

The Prayer API uses **Clerk** for identity management. Every authenticated request must include:

```
Authorization: Bearer <session_token>
```

The session token is obtained from Clerk during user login/signup.

### Frontend Setup

In **prayer-app**, configure Clerk:

```typescript
import { ClerkProvider } from '@clerk/clerk-react';

export default function App() {
  return (
    <ClerkProvider publishableKey={YOUR_PUBLISHABLE_KEY}>
      {/* Your app components */}
    </ClerkProvider>
  );
}
```

### API Client Setup

Create an API client utility to handle authentication:

```typescript
async function apiCall(endpoint: string, options: RequestInit = {}) {
  const { getToken } = useAuth();
  const token = await getToken();

  const response = await fetch(`/api${endpoint}`, {
    ...options,
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
      ...options.headers,
    },
  });

  if (!response.ok) {
    throw new Error(`API Error: ${response.status}`);
  }

  return response.json();
}
```

---

## API Endpoints

### Health Check

**GET** `/api/health`

No authentication required. Returns API status.

```bash
curl http://localhost:8080/api/health
```

Response:
```json
{
  "status": "ok"
}
```

### Authentication

#### Login / Register

**POST** `/api/auth/login`

Authenticates a user with Clerk and creates/updates the user in the Prayer API.

**Required Headers:**
```
Authorization: Bearer <clerk_session_token>
```

**Response:**
```json
{
  "user": {
    "id": "user-123",
    "externalID": "clerk-user-id",
    "name": "John Doe",
    "status": "active",
    "roles": ["role_members"],
    "prayerGroups": []
  }
}
```

---

### User Management

#### Get User Profile

**GET** `/api/users/{id}`

Retrieves the authenticated user's profile.

**Required Headers:**
```
Authorization: Bearer <session_token>
```

**Response:**
```json
{
  "user": {
    "id": "user-123",
    "name": "John Doe",
    "status": "active",
    "roles": [
      {
        "id": "role_members",
        "name": "Members",
        "permissions": ["view_prayer_sessions", "join_prayer_sessions"]
      }
    ],
    "prayerGroups": [
      {
        "id": "group-1",
        "name": "Prayer Group A",
        "blocked": false
      }
    ]
  }
}
```

---

### Role Management

#### Assign Role to User

**POST** `/api/users/{userID}/roles/{roleID}`

Assigns a role to a user (admin only).

**Required Headers:**
```
Authorization: Bearer <session_token>
```

**Available Roles:**
- `role_admin`: Administrative access
- `role_members`: Standard member access

**Response (200 OK):**
```json
{
  "user": {
    "id": "user-456",
    "name": "Jane Doe",
    "roles": ["role_members", "role_admin"]
  },
  "role": {
    "id": "role_admin",
    "name": "Administrator"
  },
  "assigned": true
}
```

**Errors:**
- `401 Unauthorized`: Invalid session
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: User or role not found

#### Remove Role from User

**DELETE** `/api/users/{userID}/roles/{roleID}`

Removes a role from a user (admin only).

**Required Headers:**
```
Authorization: Bearer <session_token>
```

**Response (200 OK):**
```json
{
  "user": { ... },
  "role": { ... },
  "removed": true
}
```

**Errors:**
- `401 Unauthorized`: Invalid session
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: User or role not found
- `409 Conflict`: Cannot remove the protected members role

---

### Prayer Group Management

#### Create Prayer Group

**POST** `/api/prayer-groups`

Creates a new prayer group (admin only).

**Required Headers:**
```
Authorization: Bearer <session_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Sunday Prayer Circle",
  "description": "Weekly prayer meeting on Sundays",
  "type": "regular"
}
```

**Response (201 Created):**
```json
{
  "id": "group-123",
  "name": "Sunday Prayer Circle",
  "description": "Weekly prayer meeting on Sundays",
  "type": "regular",
  "status": "active",
  "createdBy": "user-123",
  "createdAt": "2024-01-15T10:30:00Z"
}
```

**Errors:**
- `400 Bad Request`: Invalid input
- `401 Unauthorized`: Invalid session
- `403 Forbidden`: Insufficient permissions
- `409 Conflict`: Duplicate prayer group

#### List Prayer Groups

**GET** `/api/prayer-groups`

Lists all available prayer groups.

**Response (200 OK):**
```json
{
  "groups": [
    {
      "id": "group-1",
      "name": "Prayer Group A",
      "description": "Morning prayers",
      "type": "regular",
      "status": "active"
    },
    {
      "id": "group-2",
      "name": "Prayer Group B",
      "description": "Evening prayers",
      "type": "visitor",
      "status": "active"
    }
  ]
}
```

#### Get Prayer Group Details

**GET** `/api/prayer-groups/{groupID}`

Retrieves details of a specific prayer group.

**Required Headers:**
```
Authorization: Bearer <session_token>
```

**Response (200 OK):**
```json
{
  "id": "group-1",
  "name": "Prayer Group A",
  "description": "Morning prayers",
  "type": "regular",
  "status": "active",
  "members": [
    {
      "id": "user-1",
      "name": "John Doe",
      "role": "member"
    }
  ]
}
```

#### Update Prayer Group

**PATCH** `/api/prayer-groups/{groupID}`

Updates a prayer group (admin only).

**Required Headers:**
```
Authorization: Bearer <session_token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Updated Prayer Group Name",
  "description": "Updated description"
}
```

**Response (200 OK):**
```json
{
  "id": "group-1",
  "name": "Updated Prayer Group Name",
  "description": "Updated description",
  "status": "active"
}
```

#### Delete Prayer Group

**DELETE** `/api/prayer-groups/{groupID}`

Deletes a prayer group (admin only).

**Required Headers:**
```
Authorization: Bearer <session_token>
```

**Response (204 No Content)**

---

### Prayer Group Access Management

#### Assign User to Prayer Group

**PUT** `/api/prayer-groups/{groupID}/users/{userID}`

Assigns a user to a prayer group (admin only).

**Required Headers:**
```
Authorization: Bearer <session_token>
```

**Response (201 Created):**
```json
{
  "message": "User assigned to prayer group"
}
```

**Errors:**
- `401 Unauthorized`: Invalid session
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: User or prayer group not found

#### Remove User from Prayer Group

**DELETE** `/api/users/{userID}/prayer-groups/{groupID}`

Removes a user from a prayer group (admin only).

**Required Headers:**
```
Authorization: Bearer <session_token>
```

**Response (204 No Content)**

#### Block User from Prayer Group

**POST** `/api/users/{userID}/prayer-groups/{groupID}/block`

Blocks a user from accessing a prayer group.

**Response (204 No Content)**

**Errors:**
- `400 Bad Request`: Missing required parameters
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: User or prayer group not found
- `409 Conflict`: User is not assigned to the prayer group

---

## Implementation Steps

### 1. Setup Environment

**Create `.env` file in prayer-app:**

```env
REACT_APP_API_BASE_URL=http://localhost:8080
REACT_APP_CLERK_PUBLISHABLE_KEY=your_clerk_publishable_key
```

### 2. Install Dependencies

```bash
npm install @clerk/clerk-react axios
```

### 3. Create API Service

Create `src/services/api.ts`:

```typescript
import axios from 'axios';
import { useAuth } from '@clerk/clerk-react';

const apiClient = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL,
});

// Request interceptor to add Clerk token
apiClient.interceptors.request.use(async (config) => {
  const { getToken } = useAuth();
  const token = await getToken();
  
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  
  return config;
});

// Response interceptor for error handling
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Handle unauthorized
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default apiClient;
```

### 4. Implement User Login

Create `src/pages/Login.tsx`:

```typescript
import { useAuth, useUser } from '@clerk/clerk-react';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import apiClient from '../services/api';

export function Login() {
  const { isLoaded, isSignedIn } = useUser();
  const navigate = useNavigate();

  useEffect(() => {
    if (isLoaded && isSignedIn) {
      // Call prayer-api login endpoint
      apiClient.post('/api/auth/login').then(() => {
        navigate('/dashboard');
      });
    }
  }, [isLoaded, isSignedIn, navigate]);

  return <div>Logging in...</div>;
}
```

### 5. Implement Prayer Group List

Create `src/pages/PrayerGroups.tsx`:

```typescript
import { useEffect, useState } from 'react';
import apiClient from '../services/api';

interface PrayerGroup {
  id: string;
  name: string;
  description: string;
  status: string;
}

export function PrayerGroups() {
  const [groups, setGroups] = useState<PrayerGroup[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    apiClient
      .get('/api/prayer-groups')
      .then((res) => setGroups(res.data.groups))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div>Loading...</div>;

  return (
    <div>
      {groups.map((group) => (
        <div key={group.id}>
          <h3>{group.name}</h3>
          <p>{group.description}</p>
        </div>
      ))}
    </div>
  );
}
```

---

## Error Handling

### Common HTTP Status Codes

| Status | Meaning | Action |
|--------|---------|--------|
| 200 | OK | Request succeeded |
| 201 | Created | Resource created successfully |
| 204 | No Content | Request succeeded, no content to return |
| 400 | Bad Request | Invalid request parameters |
| 401 | Unauthorized | Missing or invalid session token |
| 403 | Forbidden | Insufficient permissions |
| 404 | Not Found | Resource not found |
| 409 | Conflict | Resource conflict (e.g., duplicate) |
| 500 | Server Error | Internal server error |

### Example Error Response

```json
{
  "error": "user not found"
}
```

### Frontend Error Handling

```typescript
async function handleRequest() {
  try {
    const response = await apiClient.post('/api/prayer-groups', {
      name: 'New Group',
    });
    console.log('Success:', response.data);
  } catch (error) {
    if (error.response?.status === 401) {
      console.error('Authentication required');
    } else if (error.response?.status === 403) {
      console.error('Permission denied');
    } else if (error.response?.status === 409) {
      console.error('Conflict:', error.response.data.error);
    } else {
      console.error('Error:', error.message);
    }
  }
}
```

---

## Development and Testing

### Running the API Locally

```bash
cd prayer-api
go run ./cmd/api
```

The API will start on `http://localhost:8080`.

### Environment Variables

Create a `.env` file in the `prayer-api` directory:

```env
CLERK_SECRET_KEY=your_clerk_secret_key
```

### Testing Endpoints with cURL

```bash
# Health check
curl http://localhost:8080/api/health

# Login (replace token with valid Clerk session)
curl -X POST http://localhost:8080/api/auth/login \
  -H "Authorization: Bearer YOUR_CLERK_TOKEN"

# List prayer groups
curl http://localhost:8080/api/prayer-groups

# Get user profile
curl -X GET http://localhost:8080/api/users/user-123 \
  -H "Authorization: Bearer YOUR_CLERK_TOKEN"
```

### Running Tests

```bash
cd prayer-api
go test -v ./...
```

### Development Workflow

1. **Make changes** to prayer-api code
2. **Run tests** to verify functionality
3. **Test endpoints** with cURL or Postman
4. **Update prayer-app** to consume new endpoints
5. **Test frontend** with the running prayer-api

---

## Next Steps

- Implement WebSocket support for real-time prayer session updates
- Add PostgreSQL repository for production data persistence
- Implement advanced role-based access control (RBAC)
- Add prayer session scheduling and notifications
- Integrate email notifications for prayer group activities

---

## Support

For issues or questions about the Prayer API, refer to:
- [ARCHITECTURE.md](./ARCHITECTURE.md) - Detailed architecture documentation
- [internal/READ_ME.md](./internal/READ_ME.md) - Core domain model documentation
- Go test files for usage examples
