import {
    InputHTMLAttributes,
    ReactNode
} from "react";

export type InputVariant =
    | "filled"
    | "outlined"
    | "ghost";

export type InputSize =
    | "sm"
    | "md"
    | "lg";

export interface InputProps
  extends Omit<InputHTMLAttributes<HTMLInputElement>, "size"> {

    label?: string;

    helperText?: string;

    error?: string;

    variant?: InputVariant;

    inputSize?: InputSize;

    leftIcon?: ReactNode;

    rightIcon?: ReactNode;

    fullWidth?: boolean;
}