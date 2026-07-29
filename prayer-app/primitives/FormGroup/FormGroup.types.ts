// New file generated from FormGroup.tsx
import type {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface FormGroupProps
    extends HTMLAttributes<HTMLDivElement> {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    required?: boolean;

    optional?: boolean;

    disabled?: boolean;

    fullWidth?: boolean;

    children: ReactNode;

}