// New file generated from OTPInput.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface OTPInputProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "onChange"
    > {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    value?: string;

    defaultValue?: string;

    onChange?(
        value: string,
    ): void;

    length?: number;

    disabled?: boolean;

    autoFocus?: boolean;

    numericOnly?: boolean;

    fullWidth?: boolean;

}