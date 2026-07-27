import { cva } from "class-variance-authority";

export const inputVariants = cva(

    [
        "w-full",
        "transition-all",
        "outline-none",
        "disabled:opacity-50",
        "disabled:cursor-not-allowed",
        "placeholder:text-slate-400"
    ],

    {

        variants: {

            variant: {

                filled:
                    "bg-slate-100 border border-transparent focus:border-primary-500 focus:bg-white",

                outlined:
                    "border border-slate-300 bg-white focus:border-primary-500",

                ghost:
                    "border-b border-slate-300 rounded-none focus:border-primary-500"

            },

            size: {

                sm:
                    "h-9 px-3 text-sm rounded-lg",

                md:
                    "h-11 px-4 text-base rounded-xl",

                lg:
                    "h-14 px-5 text-lg rounded-2xl"

            },

            hasError: {

                true:
                    "border-red-500 focus:border-red-500",

                false: ""

            }

        },

        defaultVariants: {

            variant: "outlined",

            size: "md",

            hasError: false

        }

    }

);