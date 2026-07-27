// New file generated from UserRoleSelector.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface UserRole {
    id: string;

    name: string;

    description?: string;

    disabled?: boolean;
}

export interface UserRoleSelectorProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "onChange"
    > {

    heading?: ReactNode;

    description?: ReactNode;

    roles: UserRole[];

    selectedRoleIds: string[];

    multiple?: boolean;

    loading?: boolean;

    onSelectionChange?: (
        roleIds: string[]
    ) => void;
}