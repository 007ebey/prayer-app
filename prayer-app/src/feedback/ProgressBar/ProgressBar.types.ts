// New file generated from ProgressBar.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ProgressBarSize =
    | "sm"
    | "md"
    | "lg";

export type ProgressBarTone =
    | "primary"
    | "success"
    | "warning"
    | "danger";

export interface ProgressBarProps
    extends HTMLAttributes<HTMLDivElement> {

    value: number;

    max?: number;

    showLabel?: boolean;

    label?: ReactNode;

    animated?: boolean;

    striped?: boolean;

    size?: ProgressBarSize;

    tone?: ProgressBarTone;
}