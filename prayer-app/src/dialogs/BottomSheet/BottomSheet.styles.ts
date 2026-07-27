// New file generated from BottomSheet.tsx
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

export const bottomSheetVariants = cva(
    [
        "fixed",
        "bottom-0",
        "left-0",
        "right-0",
        "z-50",
        "rounded-t-3xl",
        "border",
        "bg-background",
        "shadow-2xl",
        "animate-in",
        "slide-in-from-bottom",
        "duration-300",
        "max-h-[95vh]",
        "overflow-hidden",
    ].join(" "),
    {
        variants: {
            size: {
                sm: "h-[30vh]",
                md: "h-[50vh]",
                lg: "h-[75vh]",
                full: "h-[95vh]",
            },
        },

        defaultVariants: {
            size: "md",
        },
    }
);