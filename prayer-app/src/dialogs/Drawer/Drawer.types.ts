// New file generated from Drawer.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type DrawerSide =
    | "left"
    | "right"
    | "top"
    | "bottom";

export type DrawerSize =
    | "sm"
    | "md"
    | "lg"
    | "full";

export interface DrawerProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    open: boolean;

    onClose: () => void;

    heading?: ReactNode;

    description?: ReactNode;

    footer?: ReactNode;

    children?: ReactNode;

    side?: DrawerSide;

    size?: DrawerSize;

    closeOnBackdrop?: boolean;

    closeOnEscape?: boolean;
}