import { useEffect, useState } from 'react';
import { prayerApi } from '../services/api';

export interface PrayerGroup {
  id: string;
  name: string;
  description: string;
  type: string;
  status: string;
  createdBy?: string;
  createdAt?: string;
}

export interface UsePrayerGroupsReturn {
  groups: PrayerGroup[];
  loading: boolean;
  error: string | null;
  createGroup: (data: {
    name: string;
    description?: string;
    type?: string;
  }) => Promise<PrayerGroup>;
  updateGroup: (
    groupId: string,
    data: { name?: string; description?: string }
  ) => Promise<PrayerGroup>;
  deleteGroup: (groupId: string) => Promise<void>;
  assignUserToGroup: (groupId: string, userId: string) => Promise<void>;
  removeUserFromGroup: (userId: string, groupId: string) => Promise<void>;
  blockUserFromGroup: (userId: string, groupId: string) => Promise<void>;
  refetch: () => Promise<void>;
}

/**
 * Hook to manage prayer groups and interact with the prayer-api
 */
export function usePrayerGroups(): UsePrayerGroupsReturn {
  const [groups, setGroups] = useState<PrayerGroup[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchGroups = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await prayerApi.listPrayerGroups();
      setGroups(response.data.groups || []);
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to load prayer groups';
      setError(message);
      console.error('Error fetching prayer groups:', err);
    } finally {
      setLoading(false);
    }
  };

  const refetch = async () => {
    await fetchGroups();
  };

  useEffect(() => {
    fetchGroups();
  }, []);

  const createGroup = async (data: {
    name: string;
    description?: string;
    type?: string;
  }): Promise<PrayerGroup> => {
    try {
      const response = await prayerApi.createPrayerGroup(data);
      const newGroup = response.data;
      setGroups([...groups, newGroup]);
      return newGroup;
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to create prayer group';
      setError(message);
      throw err;
    }
  };

  const updateGroup = async (
    groupId: string,
    data: { name?: string; description?: string }
  ): Promise<PrayerGroup> => {
    try {
      const response = await prayerApi.updatePrayerGroup(groupId, data);
      const updatedGroup = response.data;
      setGroups(
        groups.map((g) => (g.id === groupId ? updatedGroup : g))
      );
      return updatedGroup;
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to update prayer group';
      setError(message);
      throw err;
    }
  };

  const deleteGroup = async (groupId: string): Promise<void> => {
    try {
      await prayerApi.deletePrayerGroup(groupId);
      setGroups(groups.filter((g) => g.id !== groupId));
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to delete prayer group';
      setError(message);
      throw err;
    }
  };

  const assignUserToGroup = async (groupId: string, userId: string): Promise<void> => {
    try {
      await prayerApi.assignUserToPrayerGroup(groupId, userId);
      // Optionally refresh the groups data
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to assign user to prayer group';
      setError(message);
      throw err;
    }
  };

  const removeUserFromGroup = async (userId: string, groupId: string): Promise<void> => {
    try {
      await prayerApi.removeUserFromPrayerGroup(userId, groupId);
      // Optionally refresh the groups data
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to remove user from prayer group';
      setError(message);
      throw err;
    }
  };

  const blockUserFromGroup = async (userId: string, groupId: string): Promise<void> => {
    try {
      await prayerApi.blockUserFromPrayerGroup(userId, groupId);
      // Optionally refresh the groups data
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to block user from prayer group';
      setError(message);
      throw err;
    }
  };

  return {
    groups,
    loading,
    error,
    createGroup,
    updateGroup,
    deleteGroup,
    assignUserToGroup,
    removeUserFromGroup,
    blockUserFromGroup,
    refetch,
  };
}
