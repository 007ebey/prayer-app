// New file generated from AuditLog.tsx
import { ReactNode } from "react";

export type AuditAction =
    | "create"
    | "update"
    | "delete"
    | "login"
    | "logout"
    | "assign"
    | "remove"
    | "approve"
    | "reject";

export interface AuditLogItem {
    id: string;

    user: string;

    action: AuditAction;

    target: string;

    timestamp: string;

    description?: string;

    ipAddress?: string;
}

export interface AuditLogProps {
    heading?: ReactNode;

    description?: ReactNode;

    logs: AuditLogItem[];

    loading?: boolean;

    emptyState?: ReactNode;
}