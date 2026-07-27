// New file generated from ConfirmationDialog.tsx
import { cva } from "class-variance-authority";

export const confirmationDialogVariants = cva(
    "w-full max-w-md rounded-xl bg-card shadow-xl",
    {
        variants: {
            intent: {
                default: "",
                success: "",
                warning: "",
                danger: "",
            },
        },

        defaultVariants: {
            intent: "default",
        },
    }
);