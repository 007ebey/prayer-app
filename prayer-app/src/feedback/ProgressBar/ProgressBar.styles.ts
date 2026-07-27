// New file generated from ProgressBar.tsx
import { cva } from "class-variance-authority";

export const progressVariants = cva(
    [
        "relative",
        "overflow-hidden",
        "rounded-full",
        "bg-muted",
    ].join(" "),
    {
        variants: {
            size: {
                sm: "h-2",
                md: "h-3",
                lg: "h-4",
            },
        },

        defaultVariants: {
            size: "md",
        },
    }
);

export const progressIndicatorVariants = cva(
    [
        "h-full",
        "rounded-full",
        "transition-all",
        "duration-300",
    ].join(" "),
    {
        variants: {
            tone: {
                primary: "bg-primary",
                success: "bg-green-600",
                warning: "bg-amber-500",
                danger: "bg-red-600",
            },

            animated: {
                true: "transition-all duration-500",
                false: "",
            },

            striped: {
                true: [
                    "bg-[linear-gradient(45deg,rgba(255,255,255,.15)_25%,transparent_25%,transparent_50%,rgba(255,255,255,.15)_50%,rgba(255,255,255,.15)_75%,transparent_75%,transparent)]",
                    "bg-[length:1rem_1rem]",
                    "animate-[progress-stripes_1s_linear_infinite]",
                ].join(" "),
                false: "",
            },
        },

        defaultVariants: {
            tone: "primary",
            animated: true,
            striped: false,
        },
    }
);