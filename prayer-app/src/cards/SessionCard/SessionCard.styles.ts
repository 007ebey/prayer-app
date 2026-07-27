// New file generated from SessionCard.tsx
import { cva } from "class-variance-authority";

export const sessionCardVariants = cva(
    "overflow-hidden rounded-xl border bg-card transition-all hover:shadow-lg",
    {
        variants: {
            status: {
                live: "border-success",
                upcoming: "",
                completed: "opacity-75",
                cancelled: "border-destructive opacity-60",
            },
        },

        defaultVariants: {
            status: "upcoming",
        },
    }
);