import type {
    ReactNode,
    TextareaHTMLAttributes,
} from "react";

export type TextareaVariant =
    | "default"
    | "filled"
    | "outline";

export type TextareaResize =
    | "none"
    | "vertical"
    | "horizontal"
    | "both";

export interface TextareaProps
    extends Omit<
        TextareaHTMLAttributes<HTMLTextAreaElement>,
        "size"
    > {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    variant?: TextareaVariant;

    resize?: TextareaResize;

    fullWidth?: boolean;
}