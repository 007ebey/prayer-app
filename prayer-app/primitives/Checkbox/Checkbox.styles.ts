import { cva } from "class-variance-authority";

export const checkboxVariants = cva(
    [
        "h-4",
        "w-4",
        "rounded",
        "border",
        "border-input",
        "text-primary",
        "focus:outline-none",
        "focus:ring-2",
        "focus:ring-ring",
        "disabled:cursor-not-allowed",
        "disabled:opacity-50",
    ].join(" "),
    {
        variants: {
            error: {
                true: "border-destructive",
                false: "",
            },
        },

        defaultVariants: {
            error: false,
        },
    }
);