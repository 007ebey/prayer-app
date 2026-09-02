import axios, { type AxiosError, type AxiosResponse } from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

// Create axios instance with base configuration
const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add Clerk token
apiClient.interceptors.request.use(
  async (config) => {
    // Get token from Clerk
    // This will be set by the context/hook
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor for error handling
apiClient.interceptors.response.use(
  (response: AxiosResponse) => response,
  (error: AxiosError) => {
    if (error.response?.status === 401) {
      // Handle unauthorized - could redirect to login
      console.error('Unauthorized - please log in');
    } else if (error.response?.status === 403) {
      console.error('Forbidden - insufficient permissions');
    }
    return Promise.reject(error);
  }
);

// API request functions
export const prayerApi = {
  // Health check
  health: () => apiClient.get('/api/health'),

  // Authentication
  login: () => apiClient.post('/api/auth/login'),

  // User endpoints
  getProfile: (userId: string) => apiClient.get(`/api/users/${userId}`),

  // Role endpoints
  assignRole: (userId: string, roleId: string) =>
    apiClient.post(`/api/users/${userId}/roles/${roleId}`),
  
  removeRole: (userId: string, roleId: string) =>
    apiClient.delete(`/api/users/${userId}/roles/${roleId}`),

  // Prayer group endpoints
  createPrayerGroup: (data: {
    name: string;
    description?: string;
    type?: string;
  }) => apiClient.post('/api/prayer-groups', data),

  listPrayerGroups: () => apiClient.get('/api/prayer-groups'),

  getPrayerGroup: (groupId: string) =>
    apiClient.get(`/api/prayer-groups/${groupId}`),

  updatePrayerGroup: (groupId: string, data: {
    name?: string;
    description?: string;
  }) => apiClient.patch(`/api/prayer-groups/${groupId}`, data),

  deletePrayerGroup: (groupId: string) =>
    apiClient.delete(`/api/prayer-groups/${groupId}`),

  // Prayer group access endpoints
  assignUserToPrayerGroup: (groupId: string, userId: string) =>
    apiClient.put(`/api/prayer-groups/${groupId}/users/${userId}`),

  removeUserFromPrayerGroup: (userId: string, groupId: string) =>
    apiClient.delete(`/api/users/${userId}/prayer-groups/${groupId}`),

  blockUserFromPrayerGroup: (userId: string, groupId: string) =>
    apiClient.post(`/api/users/${userId}/prayer-groups/${groupId}/block`),
};

// Utility function to set authorization token
export const setAuthToken = (token: string | null) => {
  if (token) {
    apiClient.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  } else {
    delete apiClient.defaults.headers.common['Authorization'];
  }
};

export default apiClient;
