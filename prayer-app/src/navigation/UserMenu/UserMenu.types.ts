import { ReactNode } from "react";

export interface UserMenuProps {
    name?: string;

    subtitle?: string;

    avatar?: ReactNode;

    showName?: boolean;

    onClick?: () => void;
}