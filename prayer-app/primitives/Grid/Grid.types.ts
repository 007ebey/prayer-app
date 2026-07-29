import type {
    HTMLAttributes,
    ReactNode
} from "react";

export type GridGap =
    | "none"
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl";

export interface GridProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    columns?: 1 | 2 | 3 | 4 | 5 | 6 | 12;

    gap?: GridGap;

    sm?: number;

    md?: number;

    lg?: number;

    xl?: number;

    autoFit?: boolean;

    minWidth?: string;
}