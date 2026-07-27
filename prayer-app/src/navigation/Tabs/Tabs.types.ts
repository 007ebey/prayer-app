import { ReactNode } from "react";

export interface TabItem {
    id: string;

    label: string;

    icon?: ReactNode;

    badge?: ReactNode;

    disabled?: boolean;
}

export interface TabsProps {
    items: TabItem[];

    activeTab: string;

    onTabChange?: (
        tab: TabItem
    ) => void;

    fullWidth?: boolean;
}