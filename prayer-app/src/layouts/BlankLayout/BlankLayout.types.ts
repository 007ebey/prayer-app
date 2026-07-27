import { ReactNode } from "react";

export interface BlankLayoutProps {
    children: ReactNode;

    centered?: boolean;

    padded?: boolean;

    background?: boolean;
}