// New file generated from PrayerTimer.tsx
import { forwardRef } from "react";

import Typography  from "../../../primitives/Typography";
import { cn } from "../../../utils/cn";

import {
    prayerTimerActionsVariants,
    prayerTimerContentVariants,
    prayerTimerVariants,
} from "./PrayerTimer.styles";

import type {
    PrayerTimerProps,
} from "./PrayerTimer.types";

const PrayerTimer = forwardRef<
    HTMLElement,
    PrayerTimerProps
>(({
    heading,
    time,
    status,
    actions,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                prayerTimerVariants(),
                className
            )}
            {...props}
        >

            <div
                className={prayerTimerContentVariants()}
            >

                {heading && (

                    <Typography variant="caption">

                        {heading}

                    </Typography>

                )}

                <Typography variant="h3">

                    {time}

                </Typography>

                {status}

            </div>

            {actions && (

                <div
                    className={prayerTimerActionsVariants()}
                >

                    {actions}

                </div>

            )}

        </section>

    );

});

PrayerTimer.displayName =
    "PrayerTimer";

export default PrayerTimer;