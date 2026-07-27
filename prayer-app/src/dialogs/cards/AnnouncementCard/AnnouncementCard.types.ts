// New file generated from AnnouncementCard.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type AnnouncementPriority =
    | "low"
    | "normal"
    | "high"
    | "urgent";

export interface AnnouncementCardProps
    extends HTMLAttributes<HTMLDivElement> {

    heading: ReactNode;

    description?: ReactNode;

    image?: string;

    author?: ReactNode;

    publishedAt?: ReactNode;

    priority?: AnnouncementPriority;

    actions?: ReactNode;
}