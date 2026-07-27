import { ReactNode } from "react";

export interface SidebarItemProps {
    id: string;

    label: string;

    icon?: ReactNode;

    badge?: ReactNode;

    active?: boolean;

    disabled?: boolean;

    collapsed?: boolean;

    onClick?: () => void;
}