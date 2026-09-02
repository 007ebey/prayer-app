# Quick Start Guide - Prayer App + Prayer API

## 5-Minute Setup

### 1. Install Dependencies
```bash
cd prayer-app
npm install
```

### 2. Configure Environment
Create `.env.local`:
```env
VITE_API_BASE_URL=http://localhost:8080
VITE_CLERK_PUBLISHABLE_KEY=your_key_from_clerk_dashboard
```

### 3. Start Both Servers

**Terminal 1 - Prayer API:**
```bash
cd prayer-api
go run ./cmd/api
```

**Terminal 2 - Prayer App:**
```bash
cd prayer-app
npm run dev
```

### 4. Open Browser
Visit `http://localhost:5173` and sign in with Clerk

---

## Using the Hooks

### Get Prayer Groups
```typescript
import { usePrayerGroups } from '@/hooks';

function MyComponent() {
  const { groups, loading } = usePrayerGroups();
  
  return (
    <ul>
      {groups.map(g => <li key={g.id}>{g.name}</li>)}
    </ul>
  );
}
```

### Get User Profile
```typescript
import { useAuth } from '@/hooks';
import { useUserProfile } from '@/hooks';

function MyComponent() {
  const { user } = useAuth();
  const { profile } = useUserProfile(user?.id);
  
  return <div>Roles: {profile?.roles.map(r => r.name).join(', ')}</div>;
}
```

### Create Prayer Group
```typescript
import { usePrayerGroups } from '@/hooks';

function MyComponent() {
  const { createGroup } = usePrayerGroups();
  
  const handleCreate = async () => {
    await createGroup({
      name: 'New Group',
      description: 'Group description'
    });
  };
  
  return <button onClick={handleCreate}>Create</button>;
}
```

---

## API Methods

All available via `prayerApi`:

```typescript
import { prayerApi } from '@/services/api';

// Prayer Groups
prayerApi.listPrayerGroups()
prayerApi.getPrayerGroup(groupId)
prayerApi.createPrayerGroup({ name, description })
prayerApi.updatePrayerGroup(groupId, { name, description })
prayerApi.deletePrayerGroup(groupId)

// User Management
prayerApi.getProfile(userId)
prayerApi.assignRole(userId, roleId)
prayerApi.removeRole(userId, roleId)

// Group Access
prayerApi.assignUserToPrayerGroup(groupId, userId)
prayerApi.removeUserFromPrayerGroup(userId, groupId)
prayerApi.blockUserFromPrayerGroup(userId, groupId)

// Auth
prayerApi.login()
prayerApi.health()
```

---

## Common Patterns

### Loading + Error Handling
```typescript
const { groups, loading, error } = usePrayerGroups();

if (loading) return <p>Loading...</p>;
if (error) return <p>Error: {error}</p>;
return <GroupList groups={groups} />;
```

### Create with Error Handling
```typescript
const { createGroup } = usePrayerGroups();

const handleCreate = async () => {
  try {
    await createGroup({ name: 'New Group' });
    alert('Created!');
  } catch (err) {
    alert('Failed: ' + err.message);
  }
};
```

### Use Auth for Protected Sections
```typescript
const { isSignedIn, user } = useAuth();

if (!isSignedIn) {
  return <div>Please sign in</div>;
}

return <div>Hello {user?.firstName}</div>;
```

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| `VITE_CLERK_PUBLISHABLE_KEY is undefined` | Create `.env.local` with your Clerk key |
| Cannot connect to prayer-api | Ensure `go run ./cmd/api` is running on port 8080 |
| Authentication fails | Verify Clerk publishable key is correct |
| "Cannot find module @clerk/clerk-react" | Run `npm install` again |
| Changes not reflecting | Restart `npm run dev` after `.env.local` changes |

---

## Next: Explore Full Guides

- **Setup Details**: Read [SETUP_GUIDE.md](./SETUP_GUIDE.md)
- **All Changes**: Read [INTEGRATION_CHANGES.md](./INTEGRATION_CHANGES.md)
- **API Docs**: Read [../prayer-api/INTEGRATION_GUIDE.md](../prayer-api/INTEGRATION_GUIDE.md)

---

## File Locations

New files created:
```
prayer-app/
├── .env.example
├── SETUP_GUIDE.md
├── INTEGRATION_CHANGES.md
├── QUICK_START.md (this file)
├── src/
│   ├── services/api.ts
│   └── hooks/
│       ├── index.ts
│       ├── useAuth.ts
│       ├── usePrayerGroups.ts
│       └── useUserProfile.ts
```

Happy coding! 🙏
