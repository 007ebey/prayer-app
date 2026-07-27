// New file generated from FileUpload.tsx
import { cva } from "class-variance-authority";

export const fileUploadVariants =
    cva(
        [
            "rounded-xl",
            "border-2",
            "border-dashed",
            "transition-colors",
            "cursor-pointer",
            "p-6",
            "text-center",
            "hover:border-primary",
            "hover:bg-accent/40",
        ].join(" "),
        {
            variants: {
                state: {
                    default: "",

                    dragging:
                        "border-primary bg-primary/5",

                    disabled:
                        "opacity-50 cursor-not-allowed",

                    error:
                        "border-destructive",
                },
            },

            defaultVariants: {
                state: "default",
            },
        }
    );