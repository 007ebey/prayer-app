// New file generated from LoadingOverlay.tsx
import { cva } from "class-variance-authority";

export const overlayVariants = cva(
    [
        "absolute",
        "inset-0",
        "z-50",
        "flex",
        "items-center",
        "justify-center",
        "transition-opacity",
    ].join(" "),
    {
        variants: {
            backdrop: {
                true: "bg-background/70",
                false: "",
            },

            blur: {
                true: "backdrop-blur-sm",
                false: "",
            },

            fullScreen: {
                true: "fixed",
                false: "absolute",
            },
        },

        defaultVariants: {
            backdrop: true,
            blur: true,
            fullScreen: false,
        },
    }
);

export const containerVariants = cva(
    [
        "flex",
        "flex-col",
        "items-center",
        "gap-4",
        "rounded-xl",
        "bg-background",
        "shadow-lg",
    ].join(" "),
    {
        variants: {
            size: {
                sm: "p-4",
                md: "p-6",
                lg: "p-8",
            },
        },

        defaultVariants: {
            size: "md",
        },
    }
);

export const spinnerVariants = cva(
    "animate-spin text-primary",
    {
        variants: {
            size: {
                sm: "h-5 w-5",
                md: "h-8 w-8",
                lg: "h-12 w-12",
            },
        },

        defaultVariants: {
            size: "md",
        },
    }
);