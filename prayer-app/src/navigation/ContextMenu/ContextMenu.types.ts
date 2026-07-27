import { ReactNode } from "react";

export interface ContextMenuItem {
    id: string;

    label: string;

    icon?: ReactNode;

    disabled?: boolean;

    danger?: boolean;

    onClick?: () => void;
}

export interface ContextMenuProps {
    open: boolean;

    x: number;

    y: number;

    items: ContextMenuItem[];

    onClose: () => void;
}