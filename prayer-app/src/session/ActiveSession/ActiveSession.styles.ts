// New file generated from ActiveSession.tsx
import { cva } from "class-variance-authority";

export const activeSessionVariants = cva(
    [
        "rounded-xl",
        "border",
        "bg-card",
        "shadow-sm",
        "overflow-hidden",
    ].join(" ")
);

export const activeSessionHeaderVariants = cva(
    [
        "flex",
        "items-center",
        "justify-between",
        "gap-4",
        "border-b",
        "px-6",
        "py-4",
    ].join(" ")
);

export const activeSessionContentVariants = cva(
    [
        "flex",
        "flex-col",
        "gap-6",
        "p-6",
    ].join(" ")
);

export const activeSessionFooterVariants = cva(
    [
        "border-t",
        "px-6",
        "py-4",
    ].join(" ")
);