// New file generated from PrayerTag.tsx
import { forwardRef } from "react";

import Badge  from "../../../primitives/Badge";
import { cn } from "../../../utils/cn";

import {
    prayerTagVariants,
} from "./PrayerTag.styles";

import type {
    PrayerTagProps,
} from "./PrayerTag.types";

const PrayerTag = forwardRef<
    HTMLSpanElement,
    PrayerTagProps
>(({
    icon,
    label,
    className,
    ...props
}, ref) => {

    return (

        <span
            ref={ref}
            className={cn(
                prayerTagVariants(),
                className
            )}
            {...props}
        >

            <Badge>

                {icon}

                {label}

            </Badge>

        </span>

    );

});

PrayerTag.displayName =
    "PrayerTag";

export default PrayerTag;