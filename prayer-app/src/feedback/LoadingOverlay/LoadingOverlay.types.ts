// New file generated from LoadingOverlay.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type LoadingOverlaySize =
    | "sm"
    | "md"
    | "lg";

export interface LoadingOverlayProps
    extends HTMLAttributes<HTMLDivElement> {

    loading: boolean;

    message?: ReactNode;

    spinner?: ReactNode;

    backdrop?: boolean;

    blur?: boolean;

    fullScreen?: boolean;

    size?: LoadingOverlaySize;
}