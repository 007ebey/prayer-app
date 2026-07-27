// New file generated from Toast.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ToastTone =
    | "info"
    | "success"
    | "warning"
    | "danger";

export interface ToastProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    heading?: ReactNode;

    description?: ReactNode;

    icon?: ReactNode;

    actions?: ReactNode;

    dismissible?: boolean;

    onDismiss?: () => void;

    tone?: ToastTone;
}