// New file generated from PrayerTimeline.tsx
import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import {
    prayerTimelineContentVariants,
    prayerTimelineFooterVariants,
    prayerTimelineHeaderVariants,
    prayerTimelineVariants,
} from "./PrayerTimeline.styles";

import type {
    PrayerTimelineProps,
} from "./PrayerTimeline.types";

const PrayerTimeline = forwardRef<
    HTMLElement,
    PrayerTimelineProps
>(({
    header,
    children,
    footer,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                prayerTimelineVariants(),
                className
            )}
            {...props}
        >

            {header && (

                <header
                    className={prayerTimelineHeaderVariants()}
                >

                    {header}

                </header>

            )}

            <div
                className={prayerTimelineContentVariants()}
            >

                {children}

            </div>

            {footer && (

                <footer
                    className={prayerTimelineFooterVariants()}
                >

                    {footer}

                </footer>

            )}

        </section>

    );

});

PrayerTimeline.displayName =
    "PrayerTimeline";

export default PrayerTimeline;