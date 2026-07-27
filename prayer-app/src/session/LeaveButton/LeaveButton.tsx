// New file generated from LeaveButton.tsx
import { forwardRef } from "react";

import {
    Loader2,
    LogOut,
} from "lucide-react";

import Button  from "../../../primitives/Button";
import { cn } from "../../utils/cn";

import {
    leaveButtonVariants,
} from "./LeaveButton.styles";

import type {
    LeaveButtonProps,
    LeaveButtonState,
} from "./LeaveButton.types";

const labelMap: Record<
    LeaveButtonState,
    string
> = {

    leave: "Leave Session",

    leaving: "Leaving...",

    left: "Left Session",

    disabled: "Unavailable",

};

const variantMap: Record<
    LeaveButtonState,
    "danger" | "secondary"
> = {

    leave: "danger",

    leaving: "secondary",

    left: "secondary",

    disabled: "secondary",

};

const LeaveButton = forwardRef<
    HTMLButtonElement,
    LeaveButtonProps
>(({
    state = "leave",
    icon,
    label,
    className,
    disabled,
    ...props
}, ref) => {

    const isDisabled =
        disabled ||
        state === "leaving" ||
        state === "left" ||
        state === "disabled";

    return (

        <Button
            ref={ref}
            variant={variantMap[state]}
            disabled={isDisabled}
            className={cn(
                leaveButtonVariants({
                    state,
                }),
                className
            )}
            {...props}
        >

            {state === "leaving"
                ? (
                    <Loader2
                        className="size-4 animate-spin"
                    />
                )
                : (
                    icon ?? (
                        <LogOut className="size-4" />
                    )
                )}

            {label ??
                labelMap[state]}

        </Button>

    );

});

LeaveButton.displayName =
    "LeaveButton";

export default LeaveButton;