// New file generated from ActivityCard.tsx
import { cva } from "class-variance-authority";

export const activityCardVariants = cva(
    "rounded-lg border bg-card p-4 transition-colors hover:bg-muted/30",
    {
        variants: {
            status: {
                success: "border-success/20",
                warning: "border-warning/20",
                danger: "border-destructive/20",
                info: "border-info/20",
                neutral: "",
            },
        },

        defaultVariants: {
            status: "neutral",
        },
    }
);