// New file generated from Dashboard.tsx
import { cva } from "class-variance-authority";

export const dashboardVariants = cva(
    "grid gap-6",
    {
        variants: {
            columns: {
                1: "grid-cols-1",
                2: "grid-cols-1 lg:grid-cols-2",
                3: "grid-cols-1 lg:grid-cols-3",
            },
        },

        defaultVariants: {
            columns: 2,
        },
    }
);