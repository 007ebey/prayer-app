import { useEffect, useState } from "react";
import {
    useAuth as useClerkAuth,
    useUser,
} from "@clerk/clerk-react";

import {
    setAuthToken,
    prayerApi,
} from "../services/api";

import axios from "axios";

/**
 * Custom hook to handle authentication with both
 * Clerk and Prayer API.
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

    const [isInitializing, setIsInitializing] =
        useState(true);

    useEffect(() => {
        const setupAuth = async () => {

            /*
             * Clerk has not finished loading yet.
             */
            if (!isAuthLoaded || !isUserLoaded) {
                return;
            }

            /*
             * User is signed out.
             */
            if (!isSignedIn) {
                setAuthToken(null);
                setIsInitializing(false);
                return;
            }

            try {
                /*
                 * Get the current Clerk token.
                 */
                const token = await getToken();

                if (!token) {
                    throw new Error(
                        "Clerk token is not available"
                    );
                }

                /*
                 * IMPORTANT:
                 *
                 * Set the token before making any
                 * Prayer API requests.
                 */
                setAuthToken(token);

                /*
                 * Get the authenticated user's email.
                 */
                const email =
                    user?.primaryEmailAddress
                        ?.emailAddress;

                if (!email) {
                    throw new Error(
                        "User email is not available"
                    );
                }

                /*
                 * Synchronize the Clerk user
                 * with the Prayer API.
                 */
                const response =
                    await prayerApi.login(email);

                console.log(
                    "Prayer API login:",
                    response.data
                );

                console.log(
                    "Roles:",
                    response.data.user.roles
                );

                console.log(
                    "Successfully authenticated with Prayer API"
                );

            } catch (error) {

                if (axios.isAxiosError(error)) {
                    console.error(
                        "Prayer API login failed"
                    );

                    console.error(
                        "Status:",
                        error.response?.status
                    );

                    console.error(
                        "Response:",
                        error.response?.data
                    );
                } else {
                    console.error(
                        "Failed to authenticate with Prayer API:",
                        error
                    );
                }

                setAuthToken(null);

            } finally {
                /*
                 * The complete authentication flow
                 * has finished.
                 */
                setIsInitializing(false);
            }
        };

        void setupAuth();

    }, [
        isAuthLoaded,
        isUserLoaded,
        isSignedIn,
        getToken,
        user,
    ]);

    return {
        isLoaded:
            isAuthLoaded &&
            isUserLoaded,

        isSignedIn:
            Boolean(isSignedIn),

        isInitializing,

        user,
        getToken,
    };
}