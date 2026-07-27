import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type SplitDirection =
    | "horizontal"
    | "vertical";

export type SplitRatio =
    | "25-75"
    | "30-70"
    | "40-60"
    | "50-50"
    | "60-40"
    | "70-30"
    | "75-25";

export interface SplitViewProps
    extends HTMLAttributes<HTMLDivElement> {

    primary: ReactNode;

    secondary: ReactNode;

    direction?: SplitDirection;

    ratio?: SplitRatio;

    reverse?: boolean;
}