// New file generated from EventCard.tsx
import { cva } from "class-variance-authority";

export const eventCardVariants = cva(
    "overflow-hidden rounded-xl border bg-card transition-all hover:shadow-lg",
    {
        variants: {
            status: {
                upcoming: "",
                live: "border-success",
                completed: "opacity-80",
                cancelled: "border-destructive opacity-70",
            },
        },

        defaultVariants: {
            status: "upcoming",
        },
    }
);