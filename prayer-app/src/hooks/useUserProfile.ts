import { useEffect, useState } from 'react';
import { prayerApi } from '../services/api';

export interface UserRole {
  id: string;
  name: string;
  permissions: string[];
}

export interface UserProfile {
  id: string;
  name: string;
  status: string;
  roles: UserRole[];
  prayerGroups: Array<{
    id: string;
    name: string;
    blocked: boolean;
  }>;
}

export interface UseUserProfileReturn {
  profile: UserProfile | null;
  loading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
  assignRole: (roleId: string) => Promise<void>;
  removeRole: (roleId: string) => Promise<void>;
}

/**
 * Hook to manage user profile and interact with the prayer-api
 */
export function useUserProfile(userId: string | undefined): UseUserProfileReturn {
  const [profile, setProfile] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchProfile = async () => {
    if (!userId) {
      setLoading(false);
      return;
    }

    try {
      setLoading(true);
      setError(null);
      const response = await prayerApi.getProfile(userId);
      setProfile(response.data.user);
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to load user profile';
      setError(message);
      console.error('Error fetching user profile:', err);
    } finally {
      setLoading(false);
    }
  };

  const refetch = async () => {
    await fetchProfile();
  };

  useEffect(() => {
    fetchProfile();
  }, [userId]);

  const assignRole = async (roleId: string): Promise<void> => {
    if (!userId) return;

    try {
      await prayerApi.assignRole(userId, roleId);
      await fetchProfile();
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to assign role';
      setError(message);
      throw err;
    }
  };

  const removeRole = async (roleId: string): Promise<void> => {
    if (!userId) return;

    try {
      await prayerApi.removeRole(userId, roleId);
      await fetchProfile();
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to remove role';
      setError(message);
      throw err;
    }
  };

  return {
    profile,
    loading,
    error,
    refetch,
    assignRole,
    removeRole,
  };
}
