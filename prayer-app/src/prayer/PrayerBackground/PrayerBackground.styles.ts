// New file generated from PrayerBackground.tsx
import { cva } from "class-variance-authority";

export const prayerBackgroundVariants = cva(
    [
        "relative",
        "overflow-hidden",
        "w-full",
        "h-full",
        "bg-background",
    ].join(" ")
);

export const prayerBackgroundMediaVariants = cva(
    [
        "absolute",
        "inset-0",
        "z-0",
    ].join(" ")
);

export const prayerBackgroundOverlayVariants = cva(
    [
        "absolute",
        "inset-0",
        "z-10",
    ].join(" ")
);

export const prayerBackgroundContentVariants = cva(
    [
        "relative",
        "z-20",
        "h-full",
        "w-full",
    ].join(" ")
);