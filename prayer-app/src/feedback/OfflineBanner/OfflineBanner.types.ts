// New file generated from OfflineBanner.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type OfflineBannerPosition =
    | "top"
    | "bottom";

export interface OfflineBannerProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    offline: boolean;

    heading?: ReactNode;

    description?: ReactNode;

    icon?: ReactNode;

    actions?: ReactNode;

    position?: OfflineBannerPosition;
}