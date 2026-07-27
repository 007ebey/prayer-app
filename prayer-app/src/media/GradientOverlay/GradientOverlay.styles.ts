import { cva } from "class-variance-authority";

export const gradientOverlayVariants = cva(
    "absolute inset-0",
    {
        variants: {
            direction: {
                top: "bg-gradient-to-t",
                bottom: "bg-gradient-to-b",
                left: "bg-gradient-to-l",
                right: "bg-gradient-to-r",
                "top-left": "bg-gradient-to-tl",
                "top-right": "bg-gradient-to-tr",
                "bottom-left": "bg-gradient-to-bl",
                "bottom-right": "bg-gradient-to-br",
            },
        },

        defaultVariants: {
            direction: "bottom",
        },
    }
);