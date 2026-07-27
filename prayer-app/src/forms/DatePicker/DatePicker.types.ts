// New file generated from DatePicker.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface DatePickerProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "onChange"
    > {

    value?: Date;

    onChange?: (
        date?: Date
    ) => void;

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    placeholder?: string;

    disabled?: boolean;

    required?: boolean;

    fullWidth?: boolean;

    minDate?: Date;

    maxDate?: Date;
}