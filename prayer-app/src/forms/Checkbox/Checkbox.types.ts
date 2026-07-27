// New file generated from Checkbox.tsx
import {
    InputHTMLAttributes,
    ReactNode,
} from "react";

export interface CheckboxProps
    extends Omit<
        InputHTMLAttributes<HTMLInputElement>,
        "size"
    > {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    indeterminate?: boolean;

    fullWidth?: boolean;
}