import { ReactNode } from "react";

import type { SidebarItemProps } from "./SideBarItem/SidebarItem.types";

export interface SidebarProps {
    items: Omit<
        SidebarItemProps,
        "active" | "collapsed" | "onClick"
    >[];

    selectedItem?: string;

    footer?: ReactNode;

    onItemClick?: (
        item: Omit<
            SidebarItemProps,
            "active" | "collapsed" | "onClick"
        >
    ) => void;
}