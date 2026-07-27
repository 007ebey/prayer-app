// New file generated from PrayerReaction.tsx
import { cva } from "class-variance-authority";

export const prayerReactionVariants = cva(
    [
        "inline-flex",
        "items-center",
        "gap-2",
        "rounded-full",
        "border",
        "transition-all",
        "duration-200",
        "select-none",
    ].join(" "),
    {
        variants: {

            size: {

                sm: "h-8 px-3 text-sm",

                md: "h-10 px-4",

                lg: "h-12 px-5 text-lg",

            },

            active: {

                true:
                    "border-primary bg-primary text-primary-foreground",

                false:
                    "border-border bg-background hover:bg-muted",

            },

        },

        defaultVariants: {

            size: "md",

            active: false,

        },

    }
);