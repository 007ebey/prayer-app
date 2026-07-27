import {
    InputHTMLAttributes,
    ReactNode,
} from "react";

export interface SwitchProps
    extends Omit<
        InputHTMLAttributes<HTMLInputElement>,
        "type" | "size" | "onChange"
    > {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    checked?: boolean;

    onCheckedChange?: (
        checked: boolean
    ) => void;
}