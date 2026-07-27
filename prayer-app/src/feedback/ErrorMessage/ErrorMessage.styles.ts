// New file generated from ErrorMessage.tsx
import { cva } from "class-variance-authority";

export const errorMessageVariants = cva(
    [
        "flex",
        "gap-3",
    ].join(" "),
    {
        variants: {
            variant: {
                inline: "",

                card: [
                    "rounded-xl",
                    "border",
                    "border-destructive/20",
                    "bg-destructive/5",
                    "p-4",
                ].join(" "),
            },
        },

        defaultVariants: {
            variant: "inline",
        },
    }
);

export const iconVariants = cva(
    [
        "mt-0.5",
        "shrink-0",
        "text-destructive",
    ].join(" ")
);