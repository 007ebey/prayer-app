// New file generated from PrayerStatus.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type PrayerStatusVariant =
    | "draft"
    | "scheduled"
    | "live"
    | "active"
    | "paused"
    | "completed"
    | "archived"
    | "cancelled";

export interface PrayerStatusProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "children"
    > {

    status?: PrayerStatusVariant;

    icon?: ReactNode;

    label?: ReactNode;

}