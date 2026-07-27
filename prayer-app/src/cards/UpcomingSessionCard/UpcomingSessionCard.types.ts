// New file generated from UpcomingSessionCard.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface UpcomingSessionCardProps
    extends HTMLAttributes<HTMLDivElement> {

    heading: ReactNode;

    description?: ReactNode;

    image?: string;

    host?: ReactNode;

    startDate: ReactNode;

    duration?: ReactNode;

    countdown?: ReactNode;

    participantCount?: number;

    maxParticipants?: number;

    reminderEnabled?: boolean;

    tags?: ReactNode;

    actions?: ReactNode;
}