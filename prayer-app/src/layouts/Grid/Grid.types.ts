import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type GridColumns =
    | 1
    | 2
    | 3
    | 4
    | 5
    | 6
    | 12;

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

    columns?: GridColumns;

    gap?: GridGap;
}