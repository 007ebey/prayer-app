import { HTMLAttributes, ReactNode } from "react";

export type ContainerSize =
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl"
    | "2xl"
    | "full";

export interface ContainerProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    size?: ContainerSize;

    centered?: boolean;

    fluid?: boolean;

    padding?: boolean;
}