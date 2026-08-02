import type {
    ReactNode,
} from "react";

export type LoginRole =
    | "participant"
    | "admin";

export interface LoginProps {

    logo?: ReactNode;

    title?: ReactNode;

    subtitle?: ReactNode;

    hero?: ReactNode;

    footer?: ReactNode;

    loading?: boolean;

    role?: LoginRole;

    onRoleChange?: (
        role: LoginRole
    ) => void;

    onGoogleLogin?: (
        role: LoginRole
    ) => void;

    onInstagramLogin?: (
        role: LoginRole
    ) => void;
}