// New file generated from PrayerWall.tsx
import { cva } from "class-variance-authority";

export const prayerWallVariants = cva(
    [
        "relative",
        "min-h-screen",
        "overflow-hidden",
        "bg-background",
    ].join(" ")
);

export const prayerWallContentVariants = cva(
    [
        "relative",
        "z-10",
        "mx-auto",
        "flex",
        "min-h-screen",
        "max-w-5xl",
        "flex-col",
        "justify-center",
        "px-6",
        "py-12",
        "text-center",
        "space-y-10",
    ].join(" ")
);