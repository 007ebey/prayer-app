// New file generated from ParticipantCounter.tsx
import { cva } from "class-variance-authority";

export const participantCounterVariants = cva(
    [
        "inline-flex",
        "items-center",
        "gap-3",
        "rounded-lg",
        "border",
        "bg-card",
        "px-4",
        "py-2",
    ].join(" ")
);

export const participantCounterIconVariants = cva(
    [
        "flex",
        "items-center",
        "justify-center",
    ].join(" ")
);

export const participantCounterContentVariants = cva(
    [
        "flex",
        "flex-col",
        "leading-tight",
    ].join(" ")
);

export const participantCounterValueVariants = cva(
    [
        "text-lg",
        "font-semibold",
    ].join(" ")
);

export const participantCounterLabelVariants = cva(
    [
        "text-sm",
        "text-muted-foreground",
    ].join(" ")
);

export const participantCounterSecondaryVariants = cva(
    [
        "text-xs",
        "text-muted-foreground",
    ].join(" ")
);