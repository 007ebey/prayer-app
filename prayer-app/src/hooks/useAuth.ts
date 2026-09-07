import { useEffect } from 'react';
import { useAuth as useClerkAuth, useUser } from '@clerk/clerk-react';
import { setAuthToken } from '../services/api';
import { prayerApi } from '../services/api';

/**
 * Custom hook to handle authentication with both Clerk and Prayer API
 */

export function useAuth() {
  const {
    getToken,
    isLoaded: isAuthLoaded,
    isSignedIn,
  } = useClerkAuth();

  const {
    user,
    isLoaded: isUserLoaded,
  } = useUser();

  useEffect(() => {
    const setupAuth = async () => {
      // Clerk has not finished determining authentication state.
      if (!isAuthLoaded) {
        return;
      }

      // User is signed out.
      if (!isSignedIn) {
        setAuthToken(null);
        return;
      }

      // User is signed in.
      try {
        const token = await getToken();

        setAuthToken(token);

        // Synchronize the authenticated Clerk user
        // with the Prayer API.
        await prayerApi.login();
      } catch (error) {
        console.error(
          "Failed to authenticate with Prayer API:",
          error
        );

        setAuthToken(null);
      }
    };

    void setupAuth();
  }, [isAuthLoaded, isSignedIn, getToken]);

  return {
    isLoaded: isAuthLoaded && isUserLoaded,
    isSignedIn: Boolean(isSignedIn),
    user,
    getToken,
  };
}
