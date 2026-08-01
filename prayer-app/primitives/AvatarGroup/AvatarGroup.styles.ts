import { cva } from "class-variance-authority";

export const avatarGroupVariants = cva(
    "flex items-center",
    {
        variants: {

            overlap: {

                none: "space-x-2",

                sm: "-space-x-2",

                md: "-space-x-3",

                lg: "-space-x-4",

            },

            reverse: {

                true: "flex-row-reverse",

                false: "",

            },

        },

        defaultVariants: {

            overlap: "md",

            reverse: false,

        },

    }
);