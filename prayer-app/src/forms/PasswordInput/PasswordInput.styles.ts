// New file generated from PasswordInput.tsx
import { cva } from "class-variance-authority";

export const passwordInputVariants = cva(
    [
        "flex",
        "items-center",
        "rounded-lg",
        "border",
        "transition-colors",
        "focus-within:ring-2",
        "focus-within:ring-primary",
        "overflow-hidden",
    ].join(" "),
    {
        variants: {

            variant: {

                outline:
                    "border-border bg-background",

                filled:
                    "border-transparent bg-muted",

                ghost:
                    "border-transparent",

            },

            inputSize: {

                sm:
                    "h-9 px-3 text-sm",

                md:
                    "h-11 px-4",

                lg:
                    "h-12 px-5 text-lg",

            },

            error: {

                true:
                    "border-destructive",

                false:
                    "",

            },

            disabled: {

                true:
                    "opacity-50",

                false:
                    "",

            },

            fullWidth: {

                true:
                    "w-full",

                false:
                    "",

            },

        },

        defaultVariants: {

            variant:
                "outline",

            inputSize:
                "md",

            fullWidth:
                true,

        },

    }
);