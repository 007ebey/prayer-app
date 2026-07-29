import type{
    HTMLAttributes,
    MouseEventHandler,
    ReactNode
} from "react";

export type ChipVariant =
    | "filled"
    | "outlined"
    | "soft";

export type ChipColor =
    | "primary"
    | "secondary"
    | "success"
    | "warning"
    | "danger"
    | "neutral";

export type ChipSize =
    | "sm"
    | "md"
    | "lg";

export interface ChipProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    variant?: ChipVariant;

    color?: ChipColor;

    size?: ChipSize;

    selected?: boolean;

    removable?: boolean;

    disabled?: boolean;

    leftIcon?: ReactNode;

    rightIcon?: ReactNode;

    onRemove?: MouseEventHandler<HTMLButtonElement>;
}