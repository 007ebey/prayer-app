import { ReactNode } from "react";

export interface MenuItem {
    id: string;

    label: string;

    icon?: ReactNode;

    badge?: ReactNode;

    disabled?: boolean;

    danger?: boolean;

    selected?: boolean;

    onClick?: () => void;
}

export interface MenuProps {
    items: MenuItem[];

    direction?: "vertical" | "horizontal";

    fullWidth?: boolean;
}