import { cva } from "class-variance-authority";

export const switchVariants = cva(
    "peer sr-only"
);

export const switchTrackVariants = cva(
    [
        "relative",
        "inline-flex",
        "h-6",
        "w-11",
        "cursor-pointer",
        "items-center",
        "rounded-full",
        "transition-colors",
        "duration-200",
    ].join(" "),
    {
        variants: {
            checked: {
                true: "bg-primary",
                false: "bg-muted",
            },

            disabled: {
                true: "cursor-not-allowed opacity-50",
                false: "",
            },
        },

        defaultVariants: {
            checked: false,
            disabled: false,
        },
    }
);

export const switchThumbVariants = cva(
    [
        "absolute",
        "left-0.5",
        "top-0.5",
        "h-5",
        "w-5",
        "rounded-full",
        "bg-white",
        "shadow",
        "transition-transform",
        "duration-200",
    ].join(" "),
    {
        variants: {
            checked: {
                true: "translate-x-5",
                false: "translate-x-0",
            },
        },

        defaultVariants: {
            checked: false,
        },
    }
);