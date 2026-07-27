// New file generated from SuccessMessage.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type SuccessMessageVariant =
    | "inline"
    | "card";

export interface SuccessMessageProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    heading?: ReactNode;

    message: ReactNode;

    icon?: ReactNode;

    actions?: ReactNode;

    variant?: SuccessMessageVariant;
}