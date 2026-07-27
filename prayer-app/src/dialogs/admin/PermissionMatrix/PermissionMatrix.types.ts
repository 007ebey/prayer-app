// New file generated from PermissionMatrix.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface Permission {
    id: string;

    label: string;
}

export interface PermissionRole {
    id: string;

    label: string;
}

export interface PermissionMatrixProps
    extends HTMLAttributes<HTMLDivElement> {

    heading?: ReactNode;

    description?: ReactNode;

    roles: PermissionRole[];

    permissions: Permission[];

    values: Record<
        string,
        Record<string, boolean>
    >;

    loading?: boolean;

    onPermissionChange?: (
        roleId: string,
        permissionId: string,
        value: boolean
    ) => void;
}