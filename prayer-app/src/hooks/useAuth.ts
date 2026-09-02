import { useEffect } from 'react';
import { useAuth as useClerkAuth, useUser } from '@clerk/clerk-react';
import { setAuthToken } from '../services/api';
import { prayerApi } from '../services/api';

const hasClerkKey = Boolean(import.meta.env.VITE_CLERK_PUBLISHABLE_KEY);

/**
 * Custom hook to handle authentication with both Clerk and Prayer API
 */
export function useAuth() {
  if (!hasClerkKey) {
    return {
      isLoaded: true,
      isSignedIn: false,
      user: undefined,
      getToken: async () => null,
    };
  }

  const { getToken, isLoaded: isAuthLoaded, isSignedIn } = useClerkAuth();
  const { user, isLoaded: isUserLoaded } = useUser();

  useEffect(() => {
    const setupAuth = async () => {
      if (isAuthLoaded && isSignedIn) {
        try {
          const token = await getToken();
          setAuthToken(token);

          // Call prayer-api login endpoint to sync user
          await prayerApi.login();
        } catch (error) {
          console.error('Failed to authenticate with Prayer API:', error);
          setAuthToken(null);
        }
      } else if (isAuthLoaded && !isSignedIn) {
        setAuthToken(null);
      }
    };

    setupAuth();
  }, [isAuthLoaded, isSignedIn, getToken]);

  return {
    isLoaded: isAuthLoaded && isUserLoaded,
    isSignedIn,
    user,
    getToken,
  };
}
