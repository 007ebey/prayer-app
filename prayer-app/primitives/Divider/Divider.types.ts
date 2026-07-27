import { HTMLAttributes, ReactNode } from "react";

export type DividerOrientation =
    | "horizontal"
    | "vertical";

export type DividerSpacing =
    | "none"
    | "sm"
    | "md"
    | "lg";

export type DividerThickness =
    | "thin"
    | "medium"
    | "thick";

export interface DividerProps
    extends HTMLAttributes<HTMLDivElement> {

    orientation?: DividerOrientation;

    spacing?: DividerSpacing;

    thickness?: DividerThickness;

    label?: ReactNode;
}