# Prayer App Setup and Integration Guide

This guide explains how to set up and use the prayer-app with the prayer-api backend.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Using the API](#using-the-api)
5. [Custom Hooks](#custom-hooks)
6. [Examples](#examples)

---

## Prerequisites

- Node.js 18+ and npm/yarn
- Clerk account with publishable key
- Prayer API running locally or accessible at a remote URL

---

## Installation

### 1. Install Dependencies

```bash
cd prayer-app
npm install
```

This will install the new dependencies:
- `@clerk/clerk-react` - Authentication provider
- `axios` - HTTP client for API requests
- `react-router-dom` - Routing support

### 2. Set Up Environment Variables

Create a `.env.local` file in the prayer-app root directory:

```env
# Prayer API Configuration
VITE_API_BASE_URL=http://localhost:8080

# Clerk Configuration
VITE_CLERK_PUBLISHABLE_KEY=your_clerk_publishable_key_here
```

Replace `your_clerk_publishable_key_here` with your actual Clerk publishable key.

### 3. Run the Development Server

```bash
npm run dev
```

The app will start on `http://localhost:5173` (or the next available port).

---

## Configuration

### Clerk Setup

The prayer-app uses Clerk for authentication. The ClerkProvider is already configured in `src/main.tsx`.

To get your Clerk publishable key:

1. Go to [Clerk Dashboard](https://dashboard.clerk.com)
2. Create a new application or select existing one
3. Copy the **Publishable Key** from the API Keys section
4. Add it to your `.env.local` file

### Prayer API Configuration

The prayer-api endpoint is configured via `VITE_API_BASE_URL`. By default, it points to `http://localhost:8080`.

To run the prayer-api locally:

```bash
cd prayer-api
go run ./cmd/api
```

---

## Using the API

### API Service (`src/services/api.ts`)

The API service provides a configured axios client with helper methods:

```typescript
import { prayerApi, setAuthToken } from '@/services/api';

// Prayer Groups
await prayerApi.listPrayerGroups();
await prayerApi.getPrayerGroup('group-123');
await prayerApi.createPrayerGroup({ name: 'New Group', description: '...' });
await prayerApi.updatePrayerGroup('group-123', { name: 'Updated' });
await prayerApi.deletePrayerGroup('group-123');

// User Management
await prayerApi.getProfile('user-123');
await prayerApi.assignRole('user-123', 'role_admin');
await prayerApi.removeRole('user-123', 'role_admin');

// Group Access
await prayerApi.assignUserToPrayerGroup('group-123', 'user-456');
await prayerApi.removeUserFromPrayerGroup('user-456', 'group-123');
await prayerApi.blockUserFromPrayerGroup('user-456', 'group-123');
```

### Setting Auth Token

The `setAuthToken` function is called automatically by the `useAuth` hook. To manually set or clear the token:

```typescript
import { setAuthToken } from '@/services/api';

setAuthToken('your_clerk_token');
// OR
setAuthToken(null); // Clear token
```

---

## Custom Hooks

### `useAuth()`

Manages authentication with Clerk and syncs with prayer-api.

```typescript
import { useAuth } from '@/hooks';

function MyComponent() {
  const { isLoaded, isSignedIn, user, getToken } = useAuth();

  if (!isLoaded) return <div>Loading...</div>;
  if (!isSignedIn) return <div>Please sign in</div>;

  return <div>Welcome, {user?.firstName}!</div>;
}
```

**Returns:**
- `isLoaded: boolean` - Authentication state is loaded
- `isSignedIn: boolean` - User is authenticated
- `user: User | null` - Clerk user object
- `getToken: () => Promise<string>` - Get current auth token

### `usePrayerGroups()`

Manages prayer groups and group operations.

```typescript
import { usePrayerGroups } from '@/hooks';

function GroupsManager() {
  const {
    groups,
    loading,
    error,
    createGroup,
    updateGroup,
    deleteGroup,
    assignUserToGroup,
    removeUserFromGroup,
    blockUserFromGroup,
  } = usePrayerGroups();

  if (loading) return <div>Loading groups...</div>;
  if (error) return <div>Error: {error}</div>;

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

**Returns:**
- `groups: PrayerGroup[]` - List of prayer groups
- `loading: boolean` - Loading state
- `error: string | null` - Error message if any
- `createGroup(data)` - Create a new prayer group
- `updateGroup(groupId, data)` - Update a prayer group
- `deleteGroup(groupId)` - Delete a prayer group
- `assignUserToGroup(groupId, userId)` - Add user to group
- `removeUserFromGroup(userId, groupId)` - Remove user from group
- `blockUserFromGroup(userId, groupId)` - Block user from group
- `refetch()` - Manually refresh groups

### `useUserProfile(userId)`

Manages user profile information.

```typescript
import { useUserProfile } from '@/hooks';

function UserProfile({ userId }) {
  const { profile, loading, error, assignRole, removeRole } = useUserProfile(userId);

  if (loading) return <div>Loading profile...</div>;
  if (error) return <div>Error: {error}</div>;
  if (!profile) return <div>No profile found</div>;

  return (
    <div>
      <h2>{profile.name}</h2>
      <p>Status: {profile.status}</p>
      <p>Roles: {profile.roles.map((r) => r.name).join(', ')}</p>
    </div>
  );
}
```

**Returns:**
- `profile: UserProfile | null` - User profile data
- `loading: boolean` - Loading state
- `error: string | null` - Error message if any
- `refetch()` - Manually refresh profile
- `assignRole(roleId)` - Assign a role to user
- `removeRole(roleId)` - Remove a role from user

---

## Examples

### Example 1: List Prayer Groups

```typescript
import { useEffect, useState } from 'react';
import { usePrayerGroups } from '@/hooks';

export function GroupsList() {
  const { groups, loading, error } = usePrayerGroups();

  if (loading) return <p>Loading groups...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <div>
      <h2>Prayer Groups</h2>
      <ul>
        {groups.map((group) => (
          <li key={group.id}>
            {group.name} ({group.status})
          </li>
        ))}
      </ul>
    </div>
  );
}
```

### Example 2: Create Prayer Group

```typescript
import { useState } from 'react';
import { usePrayerGroups } from '@/hooks';

export function CreateGroupForm() {
  const { createGroup } = usePrayerGroups();
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    try {
      await createGroup({
        name,
        description,
        type: 'regular',
      });
      setName('');
      setDescription('');
      alert('Group created successfully!');
    } catch (error) {
      alert('Failed to create group: ' + error.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        placeholder="Group name"
        value={name}
        onChange={(e) => setName(e.target.value)}
        required
      />
      <textarea
        placeholder="Description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      />
      <button type="submit" disabled={loading}>
        {loading ? 'Creating...' : 'Create Group'}
      </button>
    </form>
  );
}
```

### Example 3: User Authentication

```typescript
import { useAuth } from '@/hooks';
import { SignOutButton } from '@clerk/clerk-react';

export function Header() {
  const { isLoaded, isSignedIn, user } = useAuth();

  if (!isLoaded) return <p>Loading...</p>;

  if (!isSignedIn) {
    return <div>Please sign in</div>;
  }

  return (
    <header>
      <h1>Prayer App</h1>
      <div>
        <p>Welcome, {user?.firstName}</p>
        <SignOutButton />
      </div>
    </header>
  );
}
```

### Example 4: Display User Profile with Roles

```typescript
import { useAuth } from '@/hooks';
import { useUserProfile } from '@/hooks';

export function UserProfilePage() {
  const { user } = useAuth();
  const { profile, loading, error } = useUserProfile(user?.id);

  if (loading) return <p>Loading profile...</p>;
  if (error) return <p>Error: {error}</p>;
  if (!profile) return <p>No profile found</p>;

  return (
    <div>
      <h2>{profile.name}</h2>
      <p>Status: {profile.status}</p>
      
      <h3>Roles</h3>
      <ul>
        {profile.roles.map((role) => (
          <li key={role.id}>
            <strong>{role.name}</strong>
            <p>Permissions: {role.permissions.join(', ')}</p>
          </li>
        ))}
      </ul>

      <h3>Prayer Groups</h3>
      <ul>
        {profile.prayerGroups.map((group) => (
          <li key={group.id}>
            {group.name} {group.blocked ? '(Blocked)' : ''}
          </li>
        ))}
      </ul>
    </div>
  );
}
```

---

## Error Handling

### API Errors

All hook functions throw errors that can be caught:

```typescript
try {
  await createGroup({ name: 'New Group' });
} catch (error) {
  if (error.response?.status === 403) {
    console.error('Insufficient permissions');
  } else if (error.response?.status === 409) {
    console.error('Group already exists');
  } else {
    console.error('Error:', error.message);
  }
}
```

### Common HTTP Status Codes

| Status | Meaning |
|--------|---------|
| 200 | Success |
| 201 | Created |
| 204 | No Content |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 409 | Conflict |
| 500 | Server Error |

---

## Development Workflow

1. **Start prayer-api**:
   ```bash
   cd prayer-api
   go run ./cmd/api
   ```

2. **Start prayer-app**:
   ```bash
   cd prayer-app
   npm run dev
   ```

3. **Make changes** and test in the browser

4. **View API calls** in browser DevTools → Network tab

5. **Check prayer-api logs** for server-side errors

---

## Troubleshooting

### "Cannot find module '@clerk/clerk-react'"

Run `npm install` to ensure all dependencies are installed.

### "VITE_CLERK_PUBLISHABLE_KEY is undefined"

Make sure your `.env.local` file has the correct key and you've restarted the dev server.

### "Failed to connect to prayer-api"

- Verify prayer-api is running on `http://localhost:8080`
- Check `VITE_API_BASE_URL` in `.env.local`
- Check browser console for CORS errors

### "Authorization failed"

- Verify your Clerk publishable key is correct
- Ensure you're signed in with a valid Clerk user
- Check that the prayer-api `CLERK_SECRET_KEY` matches your Clerk secret key

---

## Next Steps

- Explore the existing components in `src/pages/` and `src/components/`
- Create custom components that use the hooks
- Add routing with react-router-dom
- Style components using Tailwind CSS
- Add form validation and error handling

---

## Resources

- [Clerk Documentation](https://clerk.com/docs)
- [Prayer API Integration Guide](../prayer-api/INTEGRATION_GUIDE.md)
- [Prayer API Architecture](../prayer-api/ARCHITECTURE.md)
- [Axios Documentation](https://axios-http.com/)
- [React Hooks Guide](https://react.dev/reference/react)
