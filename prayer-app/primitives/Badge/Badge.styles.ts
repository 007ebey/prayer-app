import { cva } from "class-variance-authority";

export const badgeVariants = cva(
    "inline-flex items-center font-medium transition-colors",
    {
        variants: {

            variant: {

                solid: "",

                soft: "",

                outline: "border"
            },

            color: {

                primary: "",

                secondary: "",

                success: "",

                warning: "",

                danger: "",

                info: "",

                neutral: ""
            },

            size: {

                sm: "px-2 py-0.5 text-xs",

                md: "px-3 py-1 text-sm",

                lg: "px-4 py-1.5 text-base"
            },

            rounded: {

                true: "rounded-full",

                false: "rounded-lg"
            }
        },

        compoundVariants: [

            {
                variant: "solid",
                color: "primary",
                className: "bg-primary-500 text-white"
            },

            {
                variant: "solid",
                color: "secondary",
                className: "bg-secondary-500 text-white"
            },

            {
                variant: "solid",
                color: "success",
                className: "bg-green-600 text-white"
            },

            {
                variant: "solid",
                color: "warning",
                className: "bg-yellow-500 text-black"
            },

            {
                variant: "solid",
                color: "danger",
                className: "bg-red-600 text-white"
            },

            {
                variant: "solid",
                color: "info",
                className: "bg-blue-600 text-white"
            },

            {
                variant: "solid",
                color: "neutral",
                className: "bg-slate-600 text-white"
            },

            {
                variant: "soft",
                color: "primary",
                className: "bg-primary-100 text-primary-700"
            },

            {
                variant: "soft",
                color: "secondary",
                className: "bg-secondary-100 text-secondary-700"
            },

            {
                variant: "soft",
                color: "success",
                className: "bg-green-100 text-green-700"
            },

            {
                variant: "soft",
                color: "warning",
                className: "bg-yellow-100 text-yellow-700"
            },

            {
                variant: "soft",
                color: "danger",
                className: "bg-red-100 text-red-700"
            },

            {
                variant: "soft",
                color: "info",
                className: "bg-blue-100 text-blue-700"
            },

            {
                variant: "soft",
                color: "neutral",
                className: "bg-slate-100 text-slate-700"
            },

            {
                variant: "outline",
                color: "primary",
                className: "border-primary-500 text-primary-600"
            },

            {
                variant: "outline",
                color: "secondary",
                className: "border-secondary-500 text-secondary-600"
            },

            {
                variant: "outline",
                color: "success",
                className: "border-green-600 text-green-700"
            },

            {
                variant: "outline",
                color: "warning",
                className: "border-yellow-500 text-yellow-700"
            },

            {
                variant: "outline",
                color: "danger",
                className: "border-red-600 text-red-700"
            },

            {
                variant: "outline",
                color: "info",
                className: "border-blue-600 text-blue-700"
            },

            {
                variant: "outline",
                color: "neutral",
                className: "border-slate-400 text-slate-700"
            }

        ],

        defaultVariants: {

            variant: "soft",

            color: "primary",

            size: "md",

            rounded: true
        }
    }
);