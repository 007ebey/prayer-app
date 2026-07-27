// New file generated from OfflineBanner.tsx
import { cva } from "class-variance-authority";

export const offlineBannerVariants = cva(
    [
        "fixed",
        "left-0",
        "right-0",
        "z-50",
        "border-b",
        "border-amber-300",
        "bg-amber-50",
        "text-amber-900",
        "shadow-md",
    ].join(" "),
    {
        variants: {
            position: {
                top: "top-0",
                bottom: [
                    "bottom-0",
                    "top-auto",
                    "border-b-0",
                    "border-t",
                ].join(" "),
            },
        },

        defaultVariants: {
            position: "top",
        },
    }
);