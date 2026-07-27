import { cva } from "class-variance-authority";

export const backdropVariants = cva(
    [
        "fixed",
        "inset-0",
        "z-40",
        "bg-black/50",
        "backdrop-blur-sm",
        "animate-in",
        "fade-in",
        "duration-200",
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
        "rounded-2xl",
        "border",
        "bg-background",
        "shadow-2xl",
        "animate-in",
        "fade-in",
        "zoom-in-95",
        "duration-200",
        "flex",
        "flex-col",
        "max-h-[90vh]",
    ].join(" "),
    {
        variants: {
            size: {
                xs: "max-w-sm",
                sm: "max-w-md",
                md: "max-w-lg",
                lg: "max-w-2xl",
                xl: "max-w-4xl",
                full: "max-w-[95vw] h-[90vh]",
            },
        },

        defaultVariants: {
            size: "md",
        },
    }
);