// New file generated from Login.tsx
import { cva } from "class-variance-authority";

export const loginVariants = cva(
    "min-h-screen flex items-center justify-center bg-background px-6 py-12"
);

export const loginCardVariants = cva(
    [
        "w-full",
        "max-w-md",
        "rounded-2xl",
        "border",
        "bg-card",
        "shadow-lg",
        "p-8",
        "space-y-6",
    ].join(" ")
);