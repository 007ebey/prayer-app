// New file generated from ConfirmDialog.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ConfirmDialogTone =
    | "default"
    | "danger"
    | "warning"
    | "success";

export interface ConfirmDialogProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    open: boolean;

    onClose: () => void;

    onConfirm: () => void;

    heading: ReactNode;

    description?: ReactNode;

    icon?: ReactNode;

    confirmLabel?: ReactNode;

    cancelLabel?: ReactNode;

    tone?: ConfirmDialogTone;

    loading?: boolean;

    closeOnBackdrop?: boolean;

    closeOnEscape?: boolean;
}