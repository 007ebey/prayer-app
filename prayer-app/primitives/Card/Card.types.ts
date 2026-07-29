import type{
    HTMLAttributes,
    ReactNode
} from "react";

export type CardVariant =
    | "filled"
    | "outlined"
    | "elevated"
    | "glass";

export type CardPadding =
    | "none"
    | "sm"
    | "md"
    | "lg";

export type CardRadius =
    | "none"
    | "sm"
    | "md"
    | "lg"
    | "xl";

export interface CardProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    variant?: CardVariant;

    padding?: CardPadding;

    radius?: CardRadius;

    hoverable?: boolean;

    clickable?: boolean;
}