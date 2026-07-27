import { cva } from "class-variance-authority";

export const imageVariants = cva(
    "block max-w-full",
    {
        variants: {
            fit: {
                cover: "object-cover",
                contain: "object-contain",
                fill: "object-fill",
                none: "object-none",
                "scale-down": "object-scale-down",
            },

            rounded: {
                true: "rounded-lg",
                false: "",
            },
        },

        defaultVariants: {
            fit: "cover",
            rounded: false,
        },
    }
);