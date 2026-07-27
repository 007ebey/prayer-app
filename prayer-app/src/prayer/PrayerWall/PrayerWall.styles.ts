// New file generated from PrayerWall.tsx
import { cva } from "class-variance-authority";

export const prayerWallVariants = cva(
    [
        "relative",
        "isolate",
        "overflow-hidden",
        "rounded-2xl",
        "bg-background",
        "min-h-screen",
    ].join(" ")
);

export const prayerWallBackgroundVariants = cva(
    [
        "absolute",
        "inset-0",
        "-z-20",
    ].join(" ")
);

export const prayerWallOverlayVariants = cva(
    [
        "absolute",
        "inset-0",
        "-z-10",
    ].join(" ")
);

export const prayerWallContainerVariants = cva(
    [
        "relative",
        "flex",
        "min-h-screen",
        "flex-col",
        "justify-between",
        "gap-8",
        "px-6",
        "py-8",
    ].join(" ")
);

export const prayerWallHeaderVariants = cva(
    [
        "flex",
        "items-center",
        "justify-between",
    ].join(" ")
);

export const prayerWallContentVariants = cva(
    [
        "flex",
        "flex-1",
        "items-center",
        "justify-center",
    ].join(" ")
);

export const prayerWallFooterVariants = cva(
    [
        "flex",
        "items-center",
        "justify-center",
    ].join(" ")
);