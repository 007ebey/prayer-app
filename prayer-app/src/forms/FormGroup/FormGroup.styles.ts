// New file generated from FormGroup.tsx
import { cva } from "class-variance-authority";

export const formGroupVariants = cva(
    "flex flex-col gap-2",
    {
        variants: {
            fullWidth: {
                true: "w-full",
                false: "",
            },
        },
        defaultVariants: {
            fullWidth: true,
        },
    }
);