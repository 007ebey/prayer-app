// New file generated from ErrorMessage.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ErrorMessageVariant =
    | "inline"
    | "card";

export interface ErrorMessageProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    heading?: ReactNode;

    message: ReactNode;

    icon?: ReactNode;

    actions?: ReactNode;

    variant?: ErrorMessageVariant;
}