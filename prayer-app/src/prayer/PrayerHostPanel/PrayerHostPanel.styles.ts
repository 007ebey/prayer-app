// New file generated from PrayerHostPanel.tsx
import { cva } from "class-variance-authority";

export const prayerHostPanelVariants = cva(
    [
        "flex",
        "flex-col",
        "gap-6",
        "rounded-xl",
        "border",
        "bg-card",
        "shadow-sm",
        "overflow-hidden",
    ].join(" ")
);

export const prayerHostPanelHeaderVariants = cva(
    [
        "border-b",
        "px-6",
        "py-4",
    ].join(" ")
);

export const prayerHostPanelContentVariants = cva(
    [
        "flex",
        "flex-col",
        "gap-6",
        "p-6",
    ].join(" ")
);

export const prayerHostPanelFooterVariants = cva(
    [
        "border-t",
        "px-6",
        "py-4",
    ].join(" ")
);