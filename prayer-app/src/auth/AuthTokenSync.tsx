import { useAuth } from "@clerk/clerk-react";
import { useEffect } from "react";

import { setAuthToken } from "../services/api";

const AuthTokenSync = () => {
    const { getToken, isSignedIn } = useAuth();

    useEffect(() => {
        const syncToken = async () => {
            if (!isSignedIn) {
                setAuthToken(null);
                return;
            }

            try {
                const token = await getToken();

                console.log(
                    "Clerk token available:",
                    !!token
                );

                setAuthToken(token);
            } catch (error) {
                console.error(
                    "Failed to get Clerk token:",
                    error
                );

                setAuthToken(null);
            }
        };

        syncToken();
    }, [getToken, isSignedIn]);

    return null;
};

export default AuthTokenSync;