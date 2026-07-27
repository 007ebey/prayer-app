import { ReactNode } from "react";

export interface PageTitleProps {
    title: ReactNode;

    subtitle?: ReactNode;

    actions?: ReactNode;

    centered?: boolean;
}