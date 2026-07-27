import type {
    ElementType,
    HTMLAttributes,
    ReactNode,
} from "react";

export interface TypographyProps
    extends HTMLAttributes<HTMLElement> {
    as?: ElementType;
    variant?:
    | "display"
    | "display-sm"
    | "h1"
    | "h2"
    | "h3"
    | "h4"
    | "title"
    | "subtitle"
    | "body-lg"
    | "body"
    | "body-sm"
    | "caption"
    | "overline";

    weight?:
    | "light"
    | "normal"
    | "medium"
    | "semibold"
    | "bold";

    align?: "left" | "center" | "right" | "justify";

    truncate?: boolean;

    children: ReactNode;
}