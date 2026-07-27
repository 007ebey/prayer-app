// New file generated from PrayerCategory.tsx
import { cva } from "class-variance-authority";

export const prayerCategoryVariants = cva(
    [
        "rounded-xl",
        "border",
        "bg-card",
        "shadow-sm",
        "overflow-hidden",
    ].join(" ")
);

export const prayerCategoryHeaderVariants = cva(
    [
        "flex",
        "items-start",
        "justify-between",
        "gap-4",
        "border-b",
        "px-6",
        "py-4",
    ].join(" ")
);

export const prayerCategoryTitleVariants = cva(
    [
        "flex",
        "items-center",
        "gap-3",
    ].join(" ")
);

export const prayerCategoryContentVariants = cva(
    [
        "flex",
        "flex-col",
        "gap-4",
        "p-6",
    ].join(" ")
);