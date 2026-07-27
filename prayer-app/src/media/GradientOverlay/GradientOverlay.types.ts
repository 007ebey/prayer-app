import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type GradientDirection =
    | "top"
    | "bottom"
    | "left"
    | "right"
    | "top-left"
    | "top-right"
    | "bottom-left"
    | "bottom-right";

export interface GradientOverlayProps
    extends HTMLAttributes<HTMLDivElement> {

    children?: ReactNode;

    direction?: GradientDirection;

    from?: string;

    to?: string;

    opacity?: number;
}