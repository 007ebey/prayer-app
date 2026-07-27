// New file generated from PrayerBackground.tsx
import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    prayerBackgroundContentVariants,
    prayerBackgroundMediaVariants,
    prayerBackgroundOverlayVariants,
    prayerBackgroundVariants,
} from "./PrayerBackground.styles";

import type {
    PrayerBackgroundProps,
} from "./PrayerBackground.types";

const PrayerBackground = forwardRef<
    HTMLDivElement,
    PrayerBackgroundProps
>(({
    image,
    video,
    overlay,
    children,
    className,
    ...props
}, ref) => {

    return (

        <div
            ref={ref}
            className={cn(
                prayerBackgroundVariants(),
                className
            )}
            {...props}
        >

            {(image || video) && (

                <div
                    className={prayerBackgroundMediaVariants()}
                >

                    {video ?? image}

                </div>

            )}

            {overlay && (

                <div
                    className={prayerBackgroundOverlayVariants()}
                >

                    {overlay}

                </div>

            )}

            <div
                className={prayerBackgroundContentVariants()}
            >

                {children}

            </div>

        </div>

    );

});

PrayerBackground.displayName =
    "PrayerBackground";

export default PrayerBackground;