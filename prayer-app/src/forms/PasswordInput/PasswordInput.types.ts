import {
    InputHTMLAttributes,
    ReactNode,
} from "react";

export interface PasswordInputProps
    extends Omit<
        InputHTMLAttributes<HTMLInputElement>,
        "size"
    > {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    leftIcon?: ReactNode;

    showToggle?: boolean;

    variant?:
        | "outline"
        | "filled"
        | "ghost";

    inputSize?:
        | "sm"
        | "md"
        | "lg";

    fullWidth?: boolean;

}