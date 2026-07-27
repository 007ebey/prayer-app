import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ScrollDirection =
    | "vertical"
    | "horizontal"
    | "both";

export interface ScrollAreaProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    direction?: ScrollDirection;

    maxHeight?: number | string;

    maxWidth?: number | string;
}