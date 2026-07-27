// New file generated from PrayerEditor.tsx
import { cva } from "class-variance-authority";

export const prayerEditorVariants = cva(
    [
        "rounded-xl",
        "border",
        "bg-card",
        "shadow-sm",
        "overflow-hidden",
    ].join(" ")
);

export const prayerEditorHeaderVariants = cva(
    [
        "border-b",
        "px-6",
        "py-4",
    ].join(" ")
);

export const prayerEditorContentVariants = cva(
    [
        "flex",
        "flex-col",
        "gap-6",
        "p-6",
    ].join(" ")
);

export const prayerEditorFooterVariants = cva(
    [
        "border-t",
        "px-6",
        "py-4",
    ].join(" ")
);