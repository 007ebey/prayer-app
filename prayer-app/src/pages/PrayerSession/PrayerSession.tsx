// New file generated from PrayerSession.tsx
import { cn } from "../../utils/cn";

import {
    prayerContentVariants,
    prayerSessionVariants,
} from "./PrayerSession.styles";

import type {
    PrayerSessionProps,
} from "./PrayerSession.types";

const PrayerSession = ({
    background,
    sessionHeader,
    prayerPoints,
    participants,
    audioPlayer,
    controls,
    footer,
}: PrayerSessionProps) => {

    return (

        <main
            className={cn(
                prayerSessionVariants()
            )}
        >

            {background}

            <div
                className={cn(
                    prayerContentVariants()
                )}
            >

                <section
                    className="space-y-6"
                >

                    {sessionHeader}

                    {prayerPoints}

                    {audioPlayer}

                    {controls}

                </section>

                <aside
                    className="space-y-6"
                >

                    {participants}

                </aside>

            </div>

            {footer}

        </main>

    );

};

export default PrayerSession;