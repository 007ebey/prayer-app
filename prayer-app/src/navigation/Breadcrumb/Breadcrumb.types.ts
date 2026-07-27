import { ReactNode } from "react";

export interface BreadcrumbItem {
    id: string;
    label: string;
    href?: string;
    icon?: ReactNode;
}

export interface BreadcrumbProps {
    items: BreadcrumbItem[];

    separator?: ReactNode;

    onItemClick?: (
        item: BreadcrumbItem
    ) => void;
}