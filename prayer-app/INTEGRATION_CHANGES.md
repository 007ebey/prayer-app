# Prayer App Integration Changes Summary

This document summarizes all the changes made to prayer-app to integrate with prayer-api.

## Overview

The prayer-app has been updated to integrate with the prayer-api backend service. This includes:
- Authentication via Clerk
- API client setup with axios
- Custom React hooks for prayer groups, user profiles, and authentication
- Environment configuration
- Comprehensive setup and integration guides

---

## Files Added

### Configuration Files

#### `.env.example`
Example environment variables file. Copy to `.env.local` and fill in your actual values.

**Contents:**
- `VITE_API_BASE_URL` - Prayer API server URL
- `VITE_CLERK_PUBLISHABLE_KEY` - Clerk authentication key

#### `.env.local.example`
Same as `.env.example`, provides a template for local development.

### Services

#### `src/services/api.ts`
Configured axios HTTP client with prayer-api endpoints.

**Exports:**
- `prayerApi` - Object with all API methods
- `setAuthToken()` - Function to set/clear authorization token
- `apiClient` - Raw axios instance

**API Methods:**
- Authentication: `login()`
- User management: `getProfile()`, `assignRole()`, `removeRole()`
- Prayer groups: `createPrayerGroup()`, `listPrayerGroups()`, `getPrayerGroup()`, `updatePrayerGroup()`, `deletePrayerGroup()`
- Group access: `assignUserToPrayerGroup()`, `removeUserFromPrayerGroup()`, `blockUserFromPrayerGroup()`
- Health: `health()`

### Custom Hooks

#### `src/hooks/useAuth.ts`
Manages Clerk authentication and syncs with prayer-api.

**Features:**
- Automatically obtains and sets Clerk token
- Calls prayer-api login endpoint on sign-in
- Clears auth token on sign-out
- Handles token refreshing

**Usage:**
```typescript
const { isLoaded, isSignedIn, user, getToken } = useAuth();
```

#### `src/hooks/usePrayerGroups.ts`
Manages prayer group operations and state.

**Features:**
- Fetch all prayer groups
- Create, update, and delete prayer groups
- Assign/remove users from groups
- Block users from groups
- Error handling and loading states

**Usage:**
```typescript
const { groups, loading, error, createGroup, updateGroup, ... } = usePrayerGroups();
```

#### `src/hooks/useUserProfile.ts`
Manages user profile data and role management.

**Features:**
- Fetch user profile by ID
- Display user roles and permissions
- Assign/remove roles
- Show prayer group access
- Error handling and loading states

**Usage:**
```typescript
const { profile, loading, error, assignRole, removeRole } = useUserProfile(userId);
```

#### `src/hooks/index.ts`
Barrel export file for all hooks, making imports cleaner.

**Usage:**
```typescript
import { useAuth, usePrayerGroups, useUserProfile } from '@/hooks';
```

### Documentation

#### `SETUP_GUIDE.md`
Comprehensive guide for setting up and using prayer-app with prayer-api.

**Sections:**
- Prerequisites and installation
- Environment configuration
- Using the API service
- Custom hooks documentation with examples
- Error handling guide
- Troubleshooting
- Development workflow

---

## Files Modified

### `package.json`
Updated with new dependencies:

**Added:**
```json
"@clerk/clerk-react": "^6.x",
"axios": "^1.x",
"react-router-dom": "^7.x"
```

### `src/main.tsx`
Wrapped app with ClerkProvider for authentication.

**Changes:**
- Import ClerkProvider from @clerk/clerk-react
- Wrap ThemeProvider with ClerkProvider
- Load VITE_CLERK_PUBLISHABLE_KEY from environment

**Before:**
```typescript
<ThemeProvider>
  <App />
</ThemeProvider>
```

**After:**
```typescript
<ClerkProvider publishableKey={publishableKey}>
  <ThemeProvider>
    <App />
  </ThemeProvider>
</ClerkProvider>
```

---

## Architecture

### Data Flow

```
Prayer App (React)
    ↓
useAuth() → Clerk Authentication
    ↓
setAuthToken() → Store token in axios headers
    ↓
usePrayerGroups() / useUserProfile() → Call prayerApi methods
    ↓
prayerApi (axios client) → HTTP requests
    ↓
Prayer API (Go backend)
    ↓
Clerk verification → Business logic → Database
```

### Hook Composition

```
Component
    ↓
useAuth() [Auto-syncs with Clerk & Prayer API]
    ↓
usePrayerGroups() or useUserProfile()
    ↓
prayerApi methods
    ↓
HTTP requests with Authorization header
```

---

## Usage Examples

### Basic Setup

1. Copy `.env.example` to `.env.local`
2. Fill in Clerk publishable key
3. Ensure prayer-api is running
4. Run `npm install`
5. Run `npm run dev`

### Using Prayer Groups Hook

```typescript
import { usePrayerGroups } from '@/hooks';

function GroupsList() {
  const { groups, loading, error } = usePrayerGroups();

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <ul>
      {groups.map(group => (
        <li key={group.id}>{group.name}</li>
      ))}
    </ul>
  );
}
```

### Using User Profile Hook

```typescript
import { useAuth } from '@/hooks';
import { useUserProfile } from '@/hooks';

function ProfilePage() {
  const { user } = useAuth();
  const { profile, loading } = useUserProfile(user?.id);

  if (loading) return <p>Loading...</p>;

  return <div>{profile?.name}</div>;
}
```

---

## Configuration

### Environment Variables

Required environment variables (create `.env.local`):

```env
VITE_API_BASE_URL=http://localhost:8080
VITE_CLERK_PUBLISHABLE_KEY=your_key_here
```

### Prayer API Connection

The app connects to prayer-api at the URL specified in `VITE_API_BASE_URL`. Make sure:

1. Prayer API is running on that URL
2. CORS is properly configured (if on different domain)
3. Both Clerk keys (app's publishable + API's secret) are valid

---

## Testing

### Manual Testing

1. Start prayer-api: `cd prayer-api && go run ./cmd/api`
2. Start prayer-app: `cd prayer-app && npm run dev`
3. Open browser DevTools → Network tab
4. Sign in with Clerk
5. Navigate and watch API calls in Network tab

### Debugging

- **Browser Console**: Shows authentication and hook errors
- **Network Tab**: Shows all API requests and responses
- **Prayer API Logs**: Server-side errors and business logic issues

---

## Next Steps

1. **Update Components**: Update existing components to use the new hooks
2. **Add Forms**: Create form components for prayer group management
3. **Add Routing**: Use react-router-dom to add pages
4. **Add Styling**: Enhance UI with existing primitives
5. **Add Tests**: Write tests for hooks and components
6. **Error UI**: Create user-friendly error messages and handling
7. **Loading States**: Add loading spinners and skeleton screens

---

## Troubleshooting

### Dependencies Not Installing

```bash
rm -rf node_modules package-lock.json
npm install
```

### Clerk Key Issues

- Verify key in `.env.local`
- Check Clerk dashboard for correct key
- Restart dev server after changing `.env.local`

### API Connection Issues

- Verify prayer-api is running: `curl http://localhost:8080/api/health`
- Check `VITE_API_BASE_URL` in `.env.local`
- Check browser console for CORS errors

### Authentication Failures

- Sign out and sign back in
- Check Clerk dashboard for user status
- Verify prayer-api `CLERK_SECRET_KEY` matches

---

## File Structure

```
prayer-app/
├── .env.example               # Environment template
├── .env.local.example         # Local env template
├── SETUP_GUIDE.md            # Setup and usage guide
├── package.json              # Updated with new dependencies
├── src/
│   ├── main.tsx              # Updated with ClerkProvider
│   ├── App.tsx               # (Existing - ready for updates)
│   ├── services/
│   │   └── api.ts           # NEW: Prayer API client
│   └── hooks/
│       ├── index.ts         # NEW: Hook exports
│       ├── useAuth.ts       # NEW: Authentication hook
│       ├── usePrayerGroups.ts # NEW: Prayer groups hook
│       └── useUserProfile.ts  # NEW: User profile hook
```

---

## Related Documentation

- [Prayer API Integration Guide](../prayer-api/INTEGRATION_GUIDE.md) - Detailed API endpoint documentation
- [Prayer API Architecture](../prayer-api/ARCHITECTURE.md) - Backend architecture details
- [Clerk Documentation](https://clerk.com/docs) - Authentication provider docs
- [Axios Documentation](https://axios-http.com/) - HTTP client docs

---

## Summary

The prayer-app now has:

✅ Clerk authentication integration  
✅ Prayer API HTTP client setup  
✅ Custom hooks for business logic  
✅ Environment configuration system  
✅ Comprehensive documentation  
✅ Error handling and loading states  
✅ Type-safe TypeScript interfaces  

Ready to build features using `usePrayerGroups()`, `useUserProfile()`, and `useAuth()` hooks!
