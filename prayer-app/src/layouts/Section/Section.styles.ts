import { cva } from "class-variance-authority";

export const sectionVariants = cva(
    "w-full",
    {
        variants: {
            bordered: {
                true: "border-b border-border",
                false: "",
            },

            padded: {
                true: "py-6",
                false: "",
            },
        },

        defaultVariants: {
            bordered: false,
            padded: true,
        },
    }
);