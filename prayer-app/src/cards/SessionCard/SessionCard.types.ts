// New file generated from SessionCard.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type SessionStatus =
    | "live"
    | "upcoming"
    | "completed"
    | "cancelled";

export interface SessionCardProps
    extends HTMLAttributes<HTMLDivElement> {

    heading: ReactNode;

    description?: ReactNode;

    image?: string;

    host?: ReactNode;

    startTime?: ReactNode;

    endTime?: ReactNode;

    duration?: ReactNode;

    participantCount?: number;

    maxParticipants?: number;

    tags?: ReactNode;

    status?: SessionStatus;

    actions?: ReactNode;
}