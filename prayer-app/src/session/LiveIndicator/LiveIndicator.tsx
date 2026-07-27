// New file generated from LiveIndicator.tsx
import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import {
    liveIndicatorDotVariants,
    liveIndicatorVariants,
} from "./LiveIndicator.styles";

import type {
    LiveIndicatorProps,
    LiveIndicatorState,
} from "./LiveIndicator.types";

const defaultLabelMap: Record<
    LiveIndicatorState,
    string
> = {

    live: "LIVE",

    connecting: "Connecting",

    offline: "Offline",

};

const LiveIndicator = forwardRef<
    HTMLSpanElement,
    LiveIndicatorProps
>(({
    state = "live",
    label,
    className,
    ...props
}, ref) => {

    return (

        <span
            ref={ref}
            className={cn(
                liveIndicatorVariants({
                    state,
                }),
                className
            )}
            {...props}
        >

            <span
                aria-hidden="true"
                className={liveIndicatorDotVariants({
                    state,
                })}
            />

            {label ??
                defaultLabelMap[state]}

        </span>

    );

});

LiveIndicator.displayName =
    "LiveIndicator";

export default LiveIndicator;