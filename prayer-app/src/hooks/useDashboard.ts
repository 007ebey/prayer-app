import { useEffect, useMemo, useState } from "react";
import { prayerApi } from "../services/api";

export interface PrayerSession {
    id: string;
    prayerGroupID: string;
    title: string;
    date: string;
    time: string;
    duration: number;
    prayerPointIDs?: string[];
}

interface DashboardState {
    activePrayer: PrayerSession | null;
    upcomingPrayers: PrayerSession[];
    isLoading: boolean;
    error: string | null;
}

const isSessionActive = (
    session: PrayerSession,
    now: Date
): boolean => {

    const start = new Date(
        `${session.date}T${session.time}`
    );

    const end = new Date(
        start.getTime() +
        session.duration * 60 * 1000
    );

    return now >= start && now < end;
};

const useDashboard = (): DashboardState => {

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
     *
     * This allows an upcoming prayer to automatically
     * become the active prayer without requiring a
     * page refresh.
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
     * Load prayer sessions.
     */
    useEffect(() => {

        const loadSessions = async () => {

            try {

                setIsLoading(true);
                setError(null);

                const response =
                    await prayerApi.listPrayerSessions();

                console.log("Loaded prayer sessions:", response.data);

                setSessions(response.data);

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

        loadSessions();

    }, []);

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
                        new Date(
                            `${session.date}T${session.time}`
                        );

                    return start > now;

                })
                .sort((a, b) => {

                    const aStart =
                        new Date(
                            `${a.date}T${a.time}`
                        ).getTime();

                    const bStart =
                        new Date(
                            `${b.date}T${b.time}`
                        ).getTime();

                    return aStart - bStart;

                });

        }, [sessions, now]);

    return {
        activePrayer,
        upcomingPrayers,
        isLoading,
        error,
    };
};

export default useDashboard;