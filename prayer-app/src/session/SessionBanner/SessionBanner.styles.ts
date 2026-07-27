// New file generated from SessionBanner.tsx
import { cva } from "class-variance-authority";

export const sessionBannerVariants = cva(
    [
        "relative",
        "overflow-hidden",
        "rounded-2xl",
        "min-h-[320px]",
        "bg-card",
        "shadow-lg",
        "isolate",
    ].join(" ")
);

export const sessionBannerBackgroundVariants = cva(
    [
        "absolute",
        "inset-0",
        "-z-20",
    ].join(" ")
);

export const sessionBannerOverlayVariants = cva(
    [
        "absolute",
        "inset-0",
        "-z-10",
        "bg-gradient-to-r",
        "from-black/70",
        "via-black/40",
        "to-transparent",
    ].join(" ")
);

export const sessionBannerContentVariants = cva(
    [
        "relative",
        "flex",
        "min-h-[320px]",
        "flex-col",
        "justify-center",
        "gap-6",
        "p-8",
        "md:max-w-2xl",
    ].join(" ")
);

export const sessionBannerMetadataVariants = cva(
    [
        "flex",
        "flex-wrap",
        "items-center",
        "gap-4",
    ].join(" ")
);

export const sessionBannerActionsVariants = cva(
    [
        "flex",
        "flex-wrap",
        "items-center",
        "gap-3",
    ].join(" ")
);