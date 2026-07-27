// New file generated from PrayerReaction.tsx
import { forwardRef } from "react";

import Typography  from "../../../primitives/Typography";
import { cn } from "../../../utils/cn";

import {
    prayerReactionVariants,
} from "./PrayerReaction.styles";

import type {
    PrayerReactionProps,
} from "./PrayerReaction.types";

const PrayerReaction = forwardRef<
    HTMLButtonElement,
    PrayerReactionProps
>(({
    icon,
    label,
    count,
    active = false,
    className,
    ...props
}, ref) => {

    return (

        <button
            ref={ref}
            type="button"
            className={cn(
                prayerReactionVariants({
                    active,
                }),
                className
            )}
            {...props}
        >

            {icon}

            {label && (

                <Typography
                    variant="body-sm"
                >

                    {label}

                </Typography>

            )}

            {count && (

                <Typography
                    variant="caption"
                    color="muted"
                >

                    {count}

                </Typography>

            )}

        </button>

    );

});

PrayerReaction.displayName =
    "PrayerReaction";

export default PrayerReaction;