import { cva } from "class-variance-authority";

export const chipVariants = cva(

    [
        "inline-flex",
        "items-center",
        "gap-2",
        "font-medium",
        "transition-all",
        "duration-200",
        "select-none"
    ],

    {

        variants: {

            variant: {

                filled: "",

                outlined: "border",

                soft: ""

            },

            color: {

                primary: "",

                secondary: "",

                success: "",

                warning: "",

                danger: "",

                neutral: ""

            },

            size: {

                sm: "px-2 py-1 text-xs rounded-full",

                md: "px-3 py-1.5 text-sm rounded-full",

                lg: "px-4 py-2 text-base rounded-full"

            },

            selected: {

                true: "ring-2 ring-primary-300",

                false: ""

            },

            disabled: {

                true: "opacity-50 pointer-events-none",

                false: ""

            }

        },

        compoundVariants: [

            {
                variant: "filled",
                color: "primary",
                className:
                    "bg-primary-500 text-white"
            },

            {
                variant: "filled",
                color: "secondary",
                className:
                    "bg-secondary-500 text-white"
            },

            {
                variant: "filled",
                color: "success",
                className:
                    "bg-green-600 text-white"
            },

            {
                variant: "filled",
                color: "warning",
                className:
                    "bg-yellow-500 text-black"
            },

            {
                variant: "filled",
                color: "danger",
                className:
                    "bg-red-600 text-white"
            },

            {
                variant: "filled",
                color: "neutral",
                className:
                    "bg-slate-500 text-white"
            },

            {
                variant: "soft",
                color: "primary",
                className:
                    "bg-primary-100 text-primary-700"
            },

            {
                variant: "soft",
                color: "secondary",
                className:
                    "bg-secondary-100 text-secondary-700"
            },

            {
                variant: "outlined",
                color: "primary",
                className:
                    "border-primary-500 text-primary-600"
            },

            {
                variant: "outlined",
                color: "secondary",
                className:
                    "border-secondary-500 text-secondary-600"
            }

        ],

        defaultVariants: {

            variant: "soft",

            color: "primary",

            size: "md",

            selected: false,

            disabled: false

        }

    }

);