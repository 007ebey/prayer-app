import { ReactNode } from "react";

export interface AuthLayoutProps {
    children: ReactNode;

    logo?: ReactNode;

    title?: ReactNode;

    subtitle?: ReactNode;

    illustration?: ReactNode;

    footer?: ReactNode;

    reverse?: boolean;
}