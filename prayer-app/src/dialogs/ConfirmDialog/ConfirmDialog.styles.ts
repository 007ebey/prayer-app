// New file generated from ConfirmDialog.tsx
import { cva } from "class-variance-authority";

export const backdropVariants = cva(
    [
        "fixed",
        "inset-0",
        "z-40",
        "bg-black/50",
        "backdrop-blur-sm",
    ].join(" ")
);

export const dialogVariants = cva(
    [
        "fixed",
        "left-1/2",
        "top-1/2",
        "-translate-x-1/2",
        "-translate-y-1/2",
        "z-50",
        "w-[calc(100%-2rem)]",
        "max-w-md",
        "rounded-2xl",
        "border",
        "bg-background",
        "shadow-xl",
        "animate-in",
        "fade-in",
        "zoom-in-95",
        "duration-200",
    ].join(" ")
);

export const iconVariants = cva(
    [
        "flex",
        "h-14",
        "w-14",
        "items-center",
        "justify-center",
        "rounded-full",
    ].join(" "),
    {
        variants: {
            tone: {
                default:
                    "bg-primary/10 text-primary",

                success:
                    "bg-success/10 text-success",

                warning:
                    "bg-warning/10 text-warning",

                danger:
                    "bg-destructive/10 text-destructive",
            },
        },

        defaultVariants: {
            tone: "default",
        },
    }
);