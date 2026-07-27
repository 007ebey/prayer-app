import { cva } from "class-variance-authority";

export const dividerVariants = cva(
    "flex items-center border-slate-200",

    {
        variants: {

            orientation: {

                horizontal:
                    "w-full border-t",

                vertical:
                    "h-full border-l self-stretch"

            },

            spacing: {

                none: "",

                sm: "my-2",

                md: "my-4",

                lg: "my-8"

            },

            thickness: {

                thin: "border",

                medium: "border-2",

                thick: "border-4"

            }

        },

        defaultVariants: {

            orientation: "horizontal",

            spacing: "md",

            thickness: "thin"

        }

    }

);