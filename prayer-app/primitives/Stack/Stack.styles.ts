import { cva } from "class-variance-authority";

export const stackVariants = cva(
    "",
    {
        variants: {

            reverse: {

                true: "flex-col-reverse",

                false: "flex-col"

            }

        },

        defaultVariants: {

            reverse: false

        }

    }
);