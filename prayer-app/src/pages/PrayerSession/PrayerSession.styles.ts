// New file generated from PrayerSession.tsx
import { cva } from "class-variance-authority";

export const prayerSessionVariants = cva(
    [
        "relative",
        "min-h-screen",
        "overflow-hidden",
        "bg-background",
    ].join(" ")
);

export const prayerContentVariants = cva(
    [
        "relative",
        "z-10",
        "mx-auto",
        "max-w-7xl",
        "px-6",
        "py-8",
        "grid",
        "gap-6",
        "lg:grid-cols-[1fr_360px]",
    ].join(" ")
);