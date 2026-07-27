import { ReactNode } from "react";

export interface DashboardLayoutProps {
    children: ReactNode;

    header?: ReactNode;

    sidebar?: ReactNode;

    mobileHeader?: ReactNode;

    bottomNavigation?: ReactNode;

    footer?: ReactNode;
}