import { useState } from "react";

import {
    Church,
} from "lucide-react";

import {
    Logo,
    Stack,
} from "../../../primitives";

import ThemeSwitcher from "../../theme/ThemeSwitcher";

import ActivePrayer from "../../ActivePrayer";
import NoActivePrayer from "../../NoActivePrayer";
import UpcomingPrayers from "../../UpcomingPrayers";

import PrayerWall from "../PrayerWall/PrayerWall";

type DashboardPage =
    | "home"
    | "prayer-wall";

const Dashboard = () => {

    // Temporary UI state.
    // Later this will come from session/backend state.
    const hasActivePrayer = true;

    const [page, setPage] =
        useState<DashboardPage>("home");


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

                    </div>


                    {/* Active Prayer */}

                    {hasActivePrayer ? (

                        <ActivePrayer
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

                    <UpcomingPrayers />

                </Stack>

            </div>

        </div>

    );

};

export default Dashboard;