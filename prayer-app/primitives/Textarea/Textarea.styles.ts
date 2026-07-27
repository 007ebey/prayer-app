import { cva } from "class-variance-authority";

export const textareaVariants = cva(
    [
        "rounded-lg",
        "border",
        "px-3",
        "py-2",
        "text-sm",
        "transition-colors",
        "focus:outline-none",
        "focus:ring-2",
        "disabled:cursor-not-allowed",
        "disabled:opacity-50",
        "placeholder:text-muted-foreground",
    ].join(" "),
    {
        variants: {
            variant: {
                default:
                    "border-input bg-background focus:ring-primary",

                filled:
                    "border-transparent bg-muted focus:ring-primary",

                outline:
                    "border-input bg-transparent focus:ring-primary",
            },

            resize: {
                none: "resize-none",
                vertical: "resize-y",
                horizontal: "resize-x",
                both: "resize",
            },

            error: {
                true: "border-destructive focus:ring-destructive",
                false: "",
            },

            fullWidth: {
                true: "w-full",
                false: "",
            },
        },

        defaultVariants: {
            variant: "default",
            resize: "vertical",
            error: false,
            fullWidth: true,
        },
    }
);