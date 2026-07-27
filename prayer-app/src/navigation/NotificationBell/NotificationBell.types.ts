import { ReactNode } from "react";

export interface NotificationBellProps {
    icon?: ReactNode;

    count?: number;

    maxCount?: number;

    showZero?: boolean;

    onClick?: () => void;
}