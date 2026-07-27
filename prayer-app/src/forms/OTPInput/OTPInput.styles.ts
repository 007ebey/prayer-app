// New file generated from OTPInput.tsx
import { cva } from "class-variance-authority";

export const otpInputVariants = cva(
    [
        "h-12",
        "w-12",
        "rounded-lg",
        "border",
        "text-center",
        "text-lg",
        "font-semibold",
        "transition-colors",
        "focus:outline-none",
        "focus:ring-2",
        "focus:ring-primary",
    ].join(" "),
    {
        variants: {

            state: {

                default:
                    "border-border",

                error:
                    "border-destructive",

                disabled:
                    "opacity-50 cursor-not-allowed",

            },

        },

        defaultVariants: {

            state:
                "default",

        },

    }
);