// New file generated from LeaveButton.tsx
import {
    ButtonHTMLAttributes,
    ReactNode,
} from "react";

export type LeaveButtonState =
    | "leave"
    | "leaving"
    | "left"
    | "disabled";

export interface LeaveButtonProps
    extends Omit<
        ButtonHTMLAttributes<HTMLButtonElement>,
        "children"
    > {

    state?: LeaveButtonState;

    icon?: ReactNode;

    label?: ReactNode;

}