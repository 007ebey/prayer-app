// New file generated from PrayerItem.tsx
import { cva } from "class-variance-authority";

export const prayerItemVariants = cva(
    [
        "rounded-xl",
        "border",
        "bg-card",
        "shadow-sm",
        "overflow-hidden",
    ].join(" ")
);

export const prayerItemContentVariants = cva(
    [
        "flex",
        "gap-4",
        "p-5",
    ].join(" ")
);

export const prayerItemBodyVariants = cva(
    [
        "flex",
        "flex-1",
        "flex-col",
        "gap-3",
    ].join(" ")
);

export const prayerItemFooterVariants = cva(
    [
        "border-t",
        "px-5",
        "py-3",
    ].join(" ")
);