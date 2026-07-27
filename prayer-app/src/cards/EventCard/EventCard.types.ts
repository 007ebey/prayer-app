// New file generated from EventCard.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type EventStatus =
    | "upcoming"
    | "live"
    | "completed"
    | "cancelled";

export interface EventCardProps
    extends HTMLAttributes<HTMLDivElement> {

    image?: string;

    heading: ReactNode;

    description?: ReactNode;

    startDate: ReactNode;

    endDate?: ReactNode;

    location?: ReactNode;

    organizer?: ReactNode;

    attendees?: number;

    status?: EventStatus;

    actions?: ReactNode;
}