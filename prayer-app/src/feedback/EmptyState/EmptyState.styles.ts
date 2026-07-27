// New file generated from EmptyState.tsx
import { cva } from "class-variance-authority";

export const emptyStateVariants = cva(
    [
        "flex",
        "flex-col",
        "items-center",
        "justify-center",
        "rounded-xl",
        "border",
        "border-dashed",
        "bg-muted/20",
        "text-center",
    ].join(" "),
    {
        variants: {
            size: {
                sm: "gap-3 p-6",
                md: "gap-5 p-10",
                lg: "gap-6 p-16",
            },
        },

        defaultVariants: {
            size: "md",
        },
    }
);

export const illustrationVariants = cva(
    "flex items-center justify-center rounded-full bg-muted",
    {
        variants: {
            size: {
                sm: "h-14 w-14",
                md: "h-20 w-20",
                lg: "h-28 w-28",
            },
        },

        defaultVariants: {
            size: "md",
        },
    }
);