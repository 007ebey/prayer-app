import { ReactNode } from "react";

export interface BottomNavigationItem {
    id: string;
    label: string;
    icon: ReactNode;
    badge?: number;
}

export interface BottomNavigationProps {
    items: BottomNavigationItem[];

    activeItem?: string;

    onItemClick?: (
        item: BottomNavigationItem
    ) => void;
}