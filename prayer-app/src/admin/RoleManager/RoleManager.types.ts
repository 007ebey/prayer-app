// New file generated from RoleManager.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface Role {
    id: string;

    name: string;

    description?: string;

    userCount?: number;

    system?: boolean;
}

export interface RoleManagerProps
    extends HTMLAttributes<HTMLDivElement> {

    heading?: ReactNode;

    description?: ReactNode;

    roles: Role[];

    loading?: boolean;

    emptyState?: ReactNode;

    onCreateRole?: () => void;

    onEditRole?: (
        role: Role
    ) => void;

    onDeleteRole?: (
        role: Role
    ) => void;

    onManagePermissions?: (
        role: Role
    ) => void;
}