import type {
    AnchorHTMLAttributes,
    ImgHTMLAttributes,
    ReactNode,
} from "react";

export type LogoVariant =
    | "full"
    | "icon"
    | "text";

export type LogoSize =
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl";

export interface LogoProps
    extends Omit<
        ImgHTMLAttributes<HTMLImageElement>,
        "children"
    > {

    /**
     * Image source.
     */
    src?: string;

    /**
     * Brand name.
     */
    name?: ReactNode;

    /**
     * Optional icon or SVG.
     */
    icon?: ReactNode;

    /**
     * Display variant.
     */
    variant?: LogoVariant;

    /**
     * Size.
     */
    size?: LogoSize;

    /**
     * Optional link.
     */
    href?: AnchorHTMLAttributes<HTMLAnchorElement>["href"];

    /**
     * Custom class.
     */
    className?: string;
}