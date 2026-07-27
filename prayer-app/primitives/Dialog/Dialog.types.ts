import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type DialogSize =
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl"
    | "full";

export interface DialogProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    open: boolean;

    onClose: () => void;

    heading?: ReactNode;

    description?: ReactNode;

    children?: ReactNode;

    footer?: ReactNode;

    size?: DialogSize;

    showCloseButton?: boolean;

    closeOnBackdrop?: boolean;

    closeOnEscape?: boolean;
}