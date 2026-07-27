// New file generated from Toast.tsx
import { cva } from "class-variance-authority";

export const toastVariants = cva(
    [
        "flex",
        "items-start",
        "gap-4",
        "rounded-xl",
        "border",
        "bg-background",
        "p-4",
        "shadow-lg",
        "min-w-[320px]",
        "max-w-md",
        "transition-all",
    ].join(" "),
    {
        variants: {
            tone: {
                info: "border-blue-200",
                success: "border-green-200",
                warning: "border-amber-200",
                danger: "border-red-200",
            },
        },

        defaultVariants: {
            tone: "info",
        },
    }
);

export const iconVariants = cva(
    "mt-0.5 shrink-0",
    {
        variants: {
            tone: {
                info: "text-blue-600",
                success: "text-green-600",
                warning: "text-amber-600",
                danger: "text-red-600",
            },
        },

        defaultVariants: {
            tone: "info",
        },
    }
);