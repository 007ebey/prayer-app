// New file generated from JoinButton.tsx
import {
    ButtonHTMLAttributes,
    ReactNode,
} from "react";

export type JoinButtonState =
    | "join"
    | "joining"
    | "joined"
    | "disabled";

export interface JoinButtonProps
    extends Omit<
        ButtonHTMLAttributes<HTMLButtonElement>,
        "children"
    > {

    state?: JoinButtonState;

    icon?: ReactNode;

    label?: ReactNode;

}