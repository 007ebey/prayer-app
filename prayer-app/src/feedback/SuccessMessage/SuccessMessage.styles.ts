// New file generated from SuccessMessage.tsx
import { cva } from "class-variance-authority";

export const successMessageVariants = cva(
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
                    "border-green-200",
                    "bg-green-50",
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
        "text-green-600",
    ].join(" ")
);

export const headingVariants = cva(
    "text-green-700"
);

export const messageVariants = cva(
    "text-green-600"
);