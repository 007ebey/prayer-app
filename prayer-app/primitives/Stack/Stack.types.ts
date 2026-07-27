import {
    HTMLAttributes,
    ReactNode
} from "react";

export type StackGap =
    | "none"
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl";

export type StackAlign =
    | "start"
    | "center"
    | "end"
    | "stretch";

export type StackJustify =
    | "start"
    | "center"
    | "end"
    | "between"
    | "around"
    | "evenly";

export interface StackProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    gap?: StackGap;

    align?: StackAlign;

    justify?: StackJustify;

    reverse?: boolean;
}