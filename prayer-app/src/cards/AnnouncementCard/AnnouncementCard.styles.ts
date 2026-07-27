// New file generated from AnnouncementCard.tsx
import { cva } from "class-variance-authority";

export const announcementCardVariants = cva(
    "overflow-hidden rounded-xl border bg-card transition-shadow hover:shadow-md",
    {
        variants: {
            priority: {
                low: "",
                normal: "",
                high: "border-warning/30",
                urgent: "border-destructive/30",
            },
        },

        defaultVariants: {
            priority: "normal",
        },
    }
);