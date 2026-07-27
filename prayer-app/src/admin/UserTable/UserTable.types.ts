// New file generated from UserTable.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type UserStatus =
    | "online"
    | "offline"
    | "away"
    | "busy";

export interface UserRole {
    id: string;

    name: string;
}

export interface User {
    id: string;

    name: string;

    email: string;

    avatar?: string;

    status?: UserStatus;

    roles: UserRole[];

    lastActive?: string;
}

export interface UserTableProps
    extends HTMLAttributes<HTMLDivElement> {

    users: User[];

    loading?: boolean;

    emptyState?: ReactNode;

    onView?: (
        user: User
    ) => void;

    onEdit?: (
        user: User
    ) => void;

    onDelete?: (
        user: User
    ) => void;

    onManageRoles?: (
        user: User
    ) => void;
}