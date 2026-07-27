import { ReactNode } from "react";

export interface HeaderProps {
    title?: ReactNode;

    subtitle?: ReactNode;

    logo?: ReactNode;

    startContent?: ReactNode;

    endContent?: ReactNode;

    sticky?: boolean;

    bordered?: boolean;
}