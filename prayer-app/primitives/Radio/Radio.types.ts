import type {
    InputHTMLAttributes,
    ReactNode,
} from "react";

export interface RadioProps
    extends Omit<
        InputHTMLAttributes<HTMLInputElement>,
        "type" | "size"
    > {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;
}