// New file generated from Alert.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type AlertVariant =
    | "solid"
    | "soft"
    | "outline";

export type AlertTone =
    | "info"
    | "success"
    | "warning"
    | "danger";

export interface AlertProps
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

    variant?: AlertVariant;

    tone?: AlertTone;
}