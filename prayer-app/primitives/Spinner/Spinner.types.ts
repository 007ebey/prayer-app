import type { HTMLAttributes } from "react";

export interface SpinnerProps
    extends HTMLAttributes<HTMLDivElement> {
    size?: "xs" | "sm" | "md" | "lg" | "xl";
    thickness?: "thin" | "normal" | "thick";
}