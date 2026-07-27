import { cva } from "class-variance-authority";

export const spinnerVariants = cva(
    "inline-block animate-spin rounded-full border-current border-solid border-t-transparent",
    {
        variants: {
            size: {
                xs: "h-3 w-3",
                sm: "h-4 w-4",
                md: "h-6 w-6",
                lg: "h-8 w-8",
                xl: "h-12 w-12",
            },
            thickness: {
                thin: "border",
                normal: "border-2",
                thick: "border-4",
            },
        },
        defaultVariants: {
            size: "md",
            thickness: "normal",
        },
    }
);