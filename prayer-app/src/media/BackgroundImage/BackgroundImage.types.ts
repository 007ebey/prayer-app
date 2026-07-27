import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type BackgroundPosition =
    | "center"
    | "top"
    | "bottom"
    | "left"
    | "right";

export interface BackgroundImageProps
    extends HTMLAttributes<HTMLDivElement> {

    src: string;

    children?: ReactNode;

    overlay?: boolean;

    overlayOpacity?: number;

    position?: BackgroundPosition;

    cover?: boolean;

    rounded?: boolean;
}