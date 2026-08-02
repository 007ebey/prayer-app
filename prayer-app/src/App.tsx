import { Church } from "lucide-react";

import {
    Stack,
    Logo,
} from "./../primitives";

import "./index.css";

import ThemeSwitcher from "./theme/ThemeSwitcher";
import ActivePrayer from "./ActivePrayer";
import NoActivePrayer from "./NoActivePrayer";
import UpcomingPrayers from "./UpcomingPrayers";
import PrayerWall from "./pages/PrayerWall/PrayerWall";
import { useState } from "react";

const App = () => {

    // Temporary UI state.
    // Later this comes from your backend/session state.
    const hasActivePrayer = true;
    const [page, setPage] = useState<"home" | "prayer-wall">("home");

     if (page === "prayer-wall") {
        return (
            <div className="min-h-screen bg-background">
                <div className="mx-auto max-w-7xl p-8">
                    <PrayerWall
                        onLeave={() => setPage("home")}
                    />
                </div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-background">

            <div className="mx-auto max-w-7xl p-8">

                <Stack gap="2xl">

                    <div className="flex items-center justify-between">

                        <Logo
                            icon={<Church />}
                            name="Prayer App"
                        />

                        <ThemeSwitcher />

                    </div>

                    {hasActivePrayer ? (
                        <ActivePrayer
                            onJoin={() => setPage("prayer-wall")}
                        />
                    ) : (
                        <NoActivePrayer />
                    )}

                    {/* Upcoming */}

                    <UpcomingPrayers />

                </Stack>

            </div>

        </div>
    );
};

export default App;