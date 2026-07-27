// New file generated from Alert.tsx
import { cva } from "class-variance-authority";

export const alertVariants = cva(
    [
        "flex",
        "gap-4",
        "rounded-xl",
        "border",
        "p-4",
        "transition-colors",
    ].join(" "),
    {
        variants: {
            variant: {
                solid: "",
                soft: "",
                outline: "bg-transparent",
            },

            tone: {
                info: "",
                success: "",
                warning: "",
                danger: "",
            },
        },

        compoundVariants: [
            {
                variant: "solid",
                tone: "info",
                class:
                    "border-blue-600 bg-blue-600 text-white",
            },
            {
                variant: "solid",
                tone: "success",
                class:
                    "border-green-600 bg-green-600 text-white",
            },
            {
                variant: "solid",
                tone: "warning",
                class:
                    "border-amber-500 bg-amber-500 text-black",
            },
            {
                variant: "solid",
                tone: "danger",
                class:
                    "border-red-600 bg-red-600 text-white",
            },

            {
                variant: "soft",
                tone: "info",
                class:
                    "border-blue-200 bg-blue-50 text-blue-900",
            },
            {
                variant: "soft",
                tone: "success",
                class:
                    "border-green-200 bg-green-50 text-green-900",
            },
            {
                variant: "soft",
                tone: "warning",
                class:
                    "border-amber-200 bg-amber-50 text-amber-900",
            },
            {
                variant: "soft",
                tone: "danger",
                class:
                    "border-red-200 bg-red-50 text-red-900",
            },

            {
                variant: "outline",
                tone: "info",
                class:
                    "border-blue-500 text-blue-700",
            },
            {
                variant: "outline",
                tone: "success",
                class:
                    "border-green-500 text-green-700",
            },
            {
                variant: "outline",
                tone: "warning",
                class:
                    "border-amber-500 text-amber-700",
            },
            {
                variant: "outline",
                tone: "danger",
                class:
                    "border-red-500 text-red-700",
            },
        ],

        defaultVariants: {
            variant: "soft",
            tone: "info",
        },
    }
);