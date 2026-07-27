// New file generated from ImageViewer.tsx
import { cva } from "class-variance-authority";

export const backdropVariants = cva(
    [
        "fixed",
        "inset-0",
        "z-50",
        "bg-black/90",
        "backdrop-blur-sm",
    ].join(" ")
);

export const viewerVariants = cva(
    [
        "fixed",
        "inset-0",
        "z-[60]",
        "flex",
        "items-center",
        "justify-center",
        "p-6",
    ].join(" ")
);

export const imageVariants = cva(
    [
        "max-h-full",
        "max-w-full",
        "rounded-xl",
        "object-contain",
        "shadow-2xl",
        "select-none",
    ].join(" ")
);