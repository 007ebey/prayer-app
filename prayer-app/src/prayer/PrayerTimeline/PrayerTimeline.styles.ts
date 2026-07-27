// New file generated from PrayerTimeline.tsx
import { cva } from "class-variance-authority";

export const prayerTimelineVariants = cva(
    [
        "rounded-xl",
        "border",
        "bg-card",
        "shadow-sm",
        "overflow-hidden",
    ].join(" ")
);

export const prayerTimelineHeaderVariants = cva(
    [
        "border-b",
        "px-6",
        "py-4",
    ].join(" ")
);

export const prayerTimelineContentVariants = cva(
    [
        "flex",
        "flex-col",
        "divide-y",
    ].join(" ")
);

export const prayerTimelineFooterVariants = cva(
    [
        "border-t",
        "px-6",
        "py-4",
    ].join(" ")
);