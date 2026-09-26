import { useState } from "react";

import {
    Church,
} from "lucide-react";

import {
    Button,
    Logo,
    Stack,
} from "../../../primitives";

import ThemeSwitcher from "../../theme/ThemeSwitcher";

import ActivePrayer from "../../ActivePrayer";
import NoActivePrayer from "../../NoActivePrayer";
import UpcomingPrayers from "../../UpcomingPrayers";

import PrayerWall from "../PrayerWall/PrayerWall";

import { useClerk } from "@clerk/clerk-react";

import { useDashboard } from "../../hooks/useDashboard";

type DashboardPage =
    | "home"
    | "prayer-wall";

const Dashboard = () => {

    const [page, setPage] =
        useState<DashboardPage>("home");

    const { signOut } = useClerk();

    const {
        activePrayer,
        upcomingPrayers,
        isLoading,
        error,
    } = useDashboard();

    const handleLogout = async () => {
        await signOut();
    };

    if (page === "prayer-wall") {

        return (

            <div className="min-h-screen bg-background">

                <div className="mx-auto max-w-7xl p-8">

                    <PrayerWall
                        onLeave={() =>
                            setPage("home")
                        }
                    />

                </div>

            </div>

        );

    }

    return (

        <div className="min-h-screen bg-background">

            <div className="mx-auto max-w-7xl p-8">

                <Stack gap="2xl">

                    {/* Header */}

                    <div className="flex items-center justify-between">

                        <Logo
                            icon={<Church />}
                            name="Prayer App"
                        />

                        <ThemeSwitcher />

                        <Button
                            variant="outline"
                            onClick={handleLogout}
                        >
                            Logout
                        </Button>

                    </div>

                    {/* Loading */}

                    {isLoading && (

                        <div>
                            Loading prayer sessions...
                        </div>

                    )}

                    {/* Error */}

                    {!isLoading && error && (

                        <div>
                            {error}
                        </div>

                    )}

                    {/* Dashboard Content */}

                    {!isLoading && !error && (

                        <>

                            {/* Active Prayer */}

                            {activePrayer ? (

                                <ActivePrayer
                                    prayer={activePrayer}
                                    onJoin={() =>
                                        setPage(
                                            "prayer-wall"
                                        )
                                    }
                                />

                            ) : (

                                <NoActivePrayer />

                            )}

                            {/* Upcoming Prayers */}

                            <UpcomingPrayers
                                prayers={
                                    upcomingPrayers
                                }
                            />

                        </>

                    )}

                </Stack>

            </div>

        </div>

    );

};

export default Dashboard;

