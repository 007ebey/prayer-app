import { ReactNode } from "react";

export interface MobileHeaderProps {
    title?: ReactNode;

    logo?: ReactNode;

    menuButton?: ReactNode;

    endContent?: ReactNode;

    sticky?: boolean;

    bordered?: boolean;
}