// New file generated from Checkbox.tsx
import { cva } from "class-variance-authority";

export const checkboxContainerVariants = cva(
    "flex items-start gap-3",
    {
        variants: {
            fullWidth: {
                true: "w-full",
                false: "",
            },
        },

        defaultVariants: {
            fullWidth: true,
        },
    }
);

export const checkboxVariants = cva(
    [
        "mt-0.5",
        "flex",
        "h-5",
        "w-5",
        "shrink-0",
        "items-center",
        "justify-center",
        "rounded",
        "border",
        "transition-colors",
        "focus-visible:outline-none",
        "focus-visible:ring-2",
        "focus-visible:ring-ring",
        "disabled:cursor-not-allowed",
        "disabled:opacity-50",
    ].join(" "),
    {
        variants: {
            checked: {
                true: "border-primary bg-primary text-primary-foreground",
                false: "border-input bg-background",
            },

            error: {
                true: "border-destructive",
                false: "",
            },
        },

        defaultVariants: {
            checked: false,
            error: false,
        },
    }
);