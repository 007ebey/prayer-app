// New file generated from SessionHeader.tsx
import { cva } from "class-variance-authority";

export const sessionHeaderVariants = cva(
    [
        "flex",
        "flex-wrap",
        "items-start",
        "gap-6",
        "border-b",
        "bg-card",
        "px-6",
        "py-5",
    ].join(" ")
);

export const sessionHeaderLeadingVariants = cva(
    [
        "flex",
        "items-center",
        "justify-center",
        "shrink-0",
    ].join(" ")
);

export const sessionHeaderIdentityVariants = cva(
    [
        "flex",
        "min-w-0",
        "flex-1",
        "flex-col",
        "gap-2",
    ].join(" ")
);

export const sessionHeaderMetadataVariants = cva(
    [
        "flex",
        "flex-wrap",
        "items-center",
        "gap-3",
    ].join(" ")
);

export const sessionHeaderActionsVariants = cva(
    [
        "flex",
        "items-center",
        "gap-2",
        "flex-wrap",
    ].join(" ")
);

export const sessionHeaderTrailingVariants = cva(
    [
        "ml-auto",
        "flex",
        "items-center",
        "gap-2",
    ].join(" ")
);