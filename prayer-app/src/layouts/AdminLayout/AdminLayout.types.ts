import { ReactNode } from "react";

export interface AdminLayoutProps {
    header?: ReactNode;

    mobileHeader?: ReactNode;

    sidebar?: ReactNode;

    footer?: ReactNode;

    children: ReactNode;
}