import { cva } from "class-variance-authority";

export const scrollAreaVariants = cva(
    "relative",
    {
        variants: {
            direction: {
                vertical: "overflow-y-auto overflow-x-hidden",
                horizontal: "overflow-x-auto overflow-y-hidden",
                both: "overflow-auto",
            },
        },

        defaultVariants: {
            direction: "vertical",
        },
    }
);