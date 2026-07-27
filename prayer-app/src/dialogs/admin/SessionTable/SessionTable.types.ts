// New file generated from SessionTable.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type SessionStatus =
    | "scheduled"
    | "live"
    | "completed"
    | "cancelled";

export interface SessionHost {
    id: string;

    name: string;

    avatar?: string;
}

export interface Session {
    id: string;

    title: string;

    host: SessionHost;

    participants: number;

    startTime: string;

    endTime: string;

    status: SessionStatus;
}

export interface SessionTableProps
    extends HTMLAttributes<HTMLDivElement> {

    sessions: Session[];

    loading?: boolean;

    emptyState?: ReactNode;

    onEdit?: (
        session: Session
    ) => void;

    onDelete?: (
        session: Session
    ) => void;

    onView?: (
        session: Session
    ) => void;

    onStart?: (
        session: Session
    ) => void;

    onEnd?: (
        session: Session
    ) => void;
}