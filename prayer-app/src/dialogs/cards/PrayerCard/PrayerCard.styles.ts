// New file generated from PrayerCard.tsx
import { cva } from "class-variance-authority";

export const prayerCardVariants = cva(
    [
        "rounded-xl",
        "border",
        "bg-card",
        "transition-all",
        "hover:shadow-md",
    ].join(" "),
    {
        variants: {
            status: {
                active: "",
                answered: "border-success",
                archived: "opacity-70",
            },
        },

        defaultVariants: {
            status: "active",
        },
    }
);