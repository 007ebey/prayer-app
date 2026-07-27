import { cva } from "class-variance-authority";

export const backgroundImageVariants = cva(
    "relative overflow-hidden",
    {
        variants: {

            cover: {
                true: "",
                false: "",
            },

            rounded: {
                true: "rounded-lg",
                false: "",
            },

            position: {
                center: "bg-center",
                top: "bg-top",
                bottom: "bg-bottom",
                left: "bg-left",
                right: "bg-right",
            },

        },

        defaultVariants: {
            cover: true,
            rounded: false,
            position: "center",
        },
    }
);