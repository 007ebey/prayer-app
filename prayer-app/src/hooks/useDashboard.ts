import { useEffect, useMemo, useState } from "react";

import {
    prayerApi,
    type PrayerSession,
} from "../services/api";

import { useAuth } from "./useAuth";

interface DashboardState {
    activePrayer: PrayerSession | null;
    upcomingPrayers: PrayerSession[];
    isLoading: boolean;
    error: string | null;
}

const getSessionStart = (
    session: PrayerSession
): Date => {
    // Go's time.Time is serialized as an ISO timestamp.
    // The session's clock time is stored separately.
    const date = session.date.slice(0, 10);

    return new Date(
        `${date}T${session.time}`
    );
};

const isSessionActive = (
    session: PrayerSession,
    now: Date
): boolean => {
    const start = getSessionStart(session);

    const end = new Date(
        start.getTime() +
        session.duration * 60 * 1000
    );

    return now >= start && now < end;
};

export const useDashboard = (): DashboardState => {

    const {
        isLoaded,
        isSignedIn,
        isInitializing,
    } = useAuth();

    const [sessions, setSessions] =
        useState<PrayerSession[]>([]);

    const [now, setNow] =
        useState(new Date());

    const [isLoading, setIsLoading] =
        useState(true);

    const [error, setError] =
        useState<string | null>(null);

    /*
     * Keep the dashboard clock updated.
     */
    useEffect(() => {

        const timer = setInterval(() => {
            setNow(new Date());
        }, 30_000);

        return () => {
            clearInterval(timer);
        };

    }, []);

    /*
     * Load prayer sessions only after the complete
     * authentication flow has finished.
     *
     * isLoaded:
     *   Clerk + user state is loaded.
     *
     * isInitializing:
     *   Our application is still synchronizing the
     *   authenticated user with the Prayer API.
     */
    useEffect(() => {

        if (!isLoaded || isInitializing) {
            return;
        }

        /*
         * User is not authenticated.
         */
        if (!isSignedIn) {
            setSessions([]);
            setIsLoading(false);
            return;
        }

        const loadSessions = async () => {

            try {

                setIsLoading(true);
                setError(null);

                const response =
                    await prayerApi.listPrayerSessions();

                console.log(
                    "Loaded prayer sessions:",
                    response.data
                );

                setSessions(
                    response.data.prayerSessions
                );

            } catch (error) {

                console.error(
                    "Failed to load prayer sessions",
                    error
                );

                setError(
                    "Unable to load prayer sessions."
                );

            } finally {

                setIsLoading(false);

            }

        };

        void loadSessions();

    }, [
        isLoaded,
        isSignedIn,
        isInitializing,
    ]);

    /*
     * Find the currently active prayer.
     */
    const activePrayer =
        useMemo(() => {

            return (
                sessions.find(
                    (session) =>
                        isSessionActive(
                            session,
                            now
                        )
                ) ?? null
            );

        }, [sessions, now]);

    /*
     * Find future prayer sessions.
     */
    const upcomingPrayers =
        useMemo(() => {

            return sessions
                .filter((session) => {

                    const start =
                        getSessionStart(session);

                    return start > now;

                })
                .sort((a, b) => {

                    return (
                        getSessionStart(a).getTime() -
                        getSessionStart(b).getTime()
                    );

                });

        }, [sessions, now]);

    return {
        activePrayer,
        upcomingPrayers,

        /*
         * Dashboard remains loading while:
         *
         * 1. Clerk is loading
         * 2. Prayer API authentication is initializing
         * 3. Prayer sessions are being fetched
         */
        isLoading:
            !isLoaded ||
            isInitializing ||
            isLoading,

        error,
    };
};

export default useDashboard;