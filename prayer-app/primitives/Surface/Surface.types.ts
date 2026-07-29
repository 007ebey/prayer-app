import type { HTMLAttributes, ReactNode } from "react";

export type SurfaceVariant =
    | "default"
    | "muted"
    | "primary"
    | "secondary"
    | "transparent";

export type SurfaceElevation =
    | "none"
    | "flat"
    | "raised"
    | "floating"
    | "overlay"
    | "modal";

export type SurfaceRadius =
    | "none"
    | "sm"
    | "md"
    | "lg"
    | "xl"
    | "2xl"
    | "full";

export type SurfacePadding =
    | "none"
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl";

export interface SurfaceProps
    extends HTMLAttributes<HTMLElement> {

    children?: ReactNode;

    variant?: SurfaceVariant;

    elevation?: SurfaceElevation;

    radius?: SurfaceRadius;

    padding?: SurfacePadding;

    bordered?: boolean;

    hoverable?: boolean;

    interactive?: boolean;

    fullWidth?: boolean;
}