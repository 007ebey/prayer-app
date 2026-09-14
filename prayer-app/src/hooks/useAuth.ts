import { useEffect } from 'react';
import { useAuth as useClerkAuth, useUser } from '@clerk/clerk-react';
import { setAuthToken } from '../services/api';
import { prayerApi } from '../services/api';
import axios from 'axios';

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
        try {
          const email = user?.primaryEmailAddress?.emailAddress;
          const response = await prayerApi.login(email);

          console.log("Prayer API login:", response.data);
          console.log("Roles:", response.data.user.roles);
        } catch (error) {
          if (axios.isAxiosError(error)) {
            console.error("Prayer API login failed");
            console.error("Status:", error.response?.status);
            console.error("Response:", error.response?.data);
          } else {
            console.error(error);
          }
          setAuthToken(null);
        }
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
