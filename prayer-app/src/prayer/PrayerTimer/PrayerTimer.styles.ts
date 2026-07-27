// New file generated from PrayerTimer.tsx
import { cva } from "class-variance-authority";

export const prayerTimerVariants = cva(
    [
        "inline-flex",
        "items-center",
        "gap-4",
        "rounded-xl",
        "border",
        "bg-card",
        "px-5",
        "py-4",
        "shadow-sm",
    ].join(" ")
);

export const prayerTimerContentVariants = cva(
    [
        "flex",
        "flex-col",
        "gap-1",
    ].join(" ")
);

export const prayerTimerActionsVariants = cva(
    [
        "ml-auto",
        "flex",
        "items-center",
        "gap-2",
    ].join(" ")
);