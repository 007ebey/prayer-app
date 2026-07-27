import { cva } from "class-variance-authority";

export const cardVariants = cva(

    [
        "transition-all",
        "duration-200",
        "overflow-hidden"
    ],

    {

        variants: {

            variant: {

                filled:
                    "bg-white",

                outlined:
                    "bg-white border border-slate-200",

                elevated:
                    "bg-white shadow-lg",

                glass:
                    "backdrop-blur-md bg-white/70 border border-white/40 shadow-xl"

            },

            padding: {

                none: "p-0",

                sm: "p-4",

                md: "p-6",

                lg: "p-8"

            },

            radius: {

                none: "",

                sm: "rounded",

                md: "rounded-xl",

                lg: "rounded-2xl",

                xl: "rounded-3xl"

            },

            hoverable: {

                true:
                    "hover:-translate-y-1 hover:shadow-xl",

                false: ""

            },

            clickable: {

                true:
                    "cursor-pointer",

                false: ""

            }

        },

        defaultVariants: {

            variant: "filled",

            padding: "md",

            radius: "lg",

            hoverable: false,

            clickable: false

        }

    }

);