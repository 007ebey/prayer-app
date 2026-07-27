import { ReactNode } from "react";

export interface NotificationItem {
    id: string;

    title: string;

    description?: string;

    timestamp?: string;

    icon?: ReactNode;

    unread?: boolean;

    onClick?: () => void;
}

export interface NotificationPopupProps {
    open: boolean;

    notifications: NotificationItem[];

    title?: string;

    emptyMessage?: string;

    anchorRef?: React.RefObject<HTMLElement>;

    onClose?: () => void;

    onViewAll?: () => void;

    onMarkAllRead?: () => void;
}