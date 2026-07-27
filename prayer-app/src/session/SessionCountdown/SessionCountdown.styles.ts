// New file generated from SessionCountdown.tsx
import { cva } from "class-variance-authority";

export const sessionCountdownVariants = cva(
    [
        "inline-flex",
        "items-center",
        "gap-4",
        "rounded-xl",
        "border",
        "bg-card",
        "px-5",
        "py-4",
        "shadow-sm",
    ].join(" ")
);

export const sessionCountdownContentVariants = cva(
    [
        "flex",
        "flex-col",
        "gap-1",
        "flex-1",
    ].join(" ")
);

export const sessionCountdownActionsVariants = cva(
    [
        "flex",
        "items-center",
        "gap-2",
        "ml-auto",
    ].join(" ")
);