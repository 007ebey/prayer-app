// New file generated from JoinButton.tsx
import { forwardRef } from "react";

import { Loader2 } from "lucide-react";

import Button  from "../../../primitives/Button";
import { cn } from "../../../utils/cn";

import {
    joinButtonVariants,
} from "./JoinButton.styles";

import type {
    JoinButtonProps,
    JoinButtonState,
} from "./JoinButton.types";

const labelMap: Record<
    JoinButtonState,
    string
> = {

    join: "Join Session",

    joining: "Joining...",

    joined: "Joined",

    disabled: "Unavailable",

};

const variantMap: Record<
    JoinButtonState,
    "primary" | "secondary" | "success"
> = {

    join: "primary",

    joining: "secondary",

    joined: "success",

    disabled: "secondary",

};

const JoinButton = forwardRef<
    HTMLButtonElement,
    JoinButtonProps
>(({
    state = "join",
    icon,
    label,
    className,
    disabled,
    ...props
}, ref) => {

    const isDisabled =
        disabled ||
        state === "joining" ||
        state === "joined" ||
        state === "disabled";

    return (

        <Button
            ref={ref}
            variant={variantMap[state]}
            disabled={isDisabled}
            className={cn(
                joinButtonVariants({
                    state,
                }),
                className
            )}
            {...props}
        >

            {state === "joining"
                ? (
                    <Loader2
                        className="size-4 animate-spin"
                    />
                )
                : icon}

            {label ??
                labelMap[state]}

        </Button>

    );

});

JoinButton.displayName =
    "JoinButton";

export default JoinButton;