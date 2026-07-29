import type { HTMLAttributes, ReactNode } from "react";

export type BadgeVariant =
    | "solid"
    | "soft"
    | "outline";

export type BadgeColor =
    | "primary"
    | "secondary"
    | "success"
    | "warning"
    | "danger"
    | "info"
    | "neutral";

export type BadgeSize =
    | "sm"
    | "md"
    | "lg";

export interface BadgeProps
    extends HTMLAttributes<HTMLSpanElement> {

    children: ReactNode;

    variant?: BadgeVariant;

    color?: BadgeColor;

    size?: BadgeSize;

    rounded?: boolean;

    leftIcon?: ReactNode;

    rightIcon?: ReactNode;
}