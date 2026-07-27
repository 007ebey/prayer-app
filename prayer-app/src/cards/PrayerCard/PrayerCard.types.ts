// New file generated from PrayerCard.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type PrayerPriority =
    | "low"
    | "normal"
    | "high"
    | "urgent";

export type PrayerStatus =
    | "active"
    | "answered"
    | "archived";

export interface PrayerCardProps
    extends HTMLAttributes<HTMLDivElement> {

    heading: ReactNode;

    description?: ReactNode;

    author?: ReactNode;

    createdAt?: ReactNode;

    category?: ReactNode;

    priority?: PrayerPriority;

    status?: PrayerStatus;

    prayerCount?: number;

    commentCount?: number;

    isPraying?: boolean;

    actions?: ReactNode;
}