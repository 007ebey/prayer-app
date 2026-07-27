// New file generated from LiveIndicator.tsx
import { cva } from "class-variance-authority";

export const liveIndicatorVariants = cva(
    [
        "inline-flex",
        "items-center",
        "gap-2",
        "rounded-full",
        "px-3",
        "py-1",
        "text-sm",
        "font-medium",
    ].join(" "),
    {
        variants: {
            state: {
                live: "",

                connecting: "",

                offline: "",
            },
        },
        defaultVariants: {
            state: "live",
        },
    }
);

export const liveIndicatorDotVariants = cva(
    [
        "h-2.5",
        "w-2.5",
        "rounded-full",
    ].join(" "),
    {
        variants: {
            state: {
                live: [
                    "bg-emerald-500",
                    "animate-pulse",
                ].join(" "),

                connecting: [
                    "bg-amber-500",
                    "animate-pulse",
                ].join(" "),

                offline: [
                    "bg-gray-400",
                ].join(" "),
            },
        },
        defaultVariants: {
            state: "live",
        },
    }
);