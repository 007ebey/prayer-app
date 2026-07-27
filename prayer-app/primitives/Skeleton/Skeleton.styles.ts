import { cva } from "class-variance-authority";

export const skeletonVariants = cva(
    "bg-muted overflow-hidden relative",
    {
        variants: {
            variant: {
                text: "h-4 rounded",
                rect: "rounded-none",
                rounded: "rounded-lg",
                circle: "rounded-full",
            },
            animation: {
                none: "",
                pulse: "animate-pulse",
                wave:
                    "before:absolute before:inset-0 before:-translate-x-full before:animate-[shimmer_1.6s_infinite] before:bg-gradient-to-r before:from-transparent before:via-white/30 before:to-transparent",
            },
        },
        defaultVariants: {
            variant: "rect",
            animation: "pulse",
        },
    }
);