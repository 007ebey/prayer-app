import {
    HTMLAttributes,
    ReactNode
} from "react";

export type FlexDirection =
    | "row"
    | "column"
    | "row-reverse"
    | "column-reverse";

export type FlexAlign =
    | "start"
    | "center"
    | "end"
    | "stretch"
    | "baseline";

export type FlexJustify =
    | "start"
    | "center"
    | "end"
    | "between"
    | "around"
    | "evenly";

export type FlexWrap =
    | "nowrap"
    | "wrap"
    | "wrap-reverse";

export type FlexGap =
    | "none"
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl";

export interface FlexProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    direction?: FlexDirection;

    align?: FlexAlign;

    justify?: FlexJustify;

    wrap?: FlexWrap;

    gap?: FlexGap;

    inline?: boolean;
}