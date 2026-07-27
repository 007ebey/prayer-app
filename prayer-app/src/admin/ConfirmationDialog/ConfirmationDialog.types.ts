// New file generated from ConfirmationDialog.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ConfirmationIntent =
    | "default"
    | "success"
    | "warning"
    | "danger";

export interface ConfirmationDialogProps
    extends HTMLAttributes<HTMLDivElement> {

    open: boolean;

    heading: ReactNode;

    description?: ReactNode;

    confirmLabel?: ReactNode;

    cancelLabel?: ReactNode;

    intent?: ConfirmationIntent;

    loading?: boolean;

    icon?: ReactNode;

    onConfirm?: () => void;

    onCancel?: () => void;
}