// New file generated from PrayerStatus.tsx
import { forwardRef } from "react";

import Badge from "../../../primitives/Badge";

import { cn } from "../../../utils/cn";

import {
    prayerStatusVariants,
} from "./PrayerStatus.styles";

import type {
    PrayerStatusProps,
    PrayerStatusVariant,
} from "./PrayerStatus.types";

const badgeColorMap: Record<
    PrayerStatusVariant,
    "secondary" | "warning" | "success" | "danger"
> = {

    draft: "secondary",

    scheduled: "warning",

    live: "danger",

    active: "success",

    paused: "warning",

    completed: "success",

    archived: "secondary",

    cancelled: "danger",

};

const defaultLabelMap: Record<
    PrayerStatusVariant,
    string
> = {

    draft: "Draft",

    scheduled: "Scheduled",

    live: "Live",

    active: "Active",

    paused: "Paused",

    completed: "Completed",

    archived: "Archived",

    cancelled: "Cancelled",

};

const PrayerStatus = forwardRef<
    HTMLDivElement,
    PrayerStatusProps
>(({
    status = "draft",
    icon,
    label,
    className,
    ...props
}, ref) => {

    return (

        <div
            ref={ref}
            className={cn(
                prayerStatusVariants(),
                className
            )}
            {...props}
        >

            <Badge
                color={badgeColorMap[status]}
            >

                {icon}

                {label ??
                    defaultLabelMap[status]}

            </Badge>

        </div>

    );

});

PrayerStatus.displayName =
    "PrayerStatus";

export default PrayerStatus;