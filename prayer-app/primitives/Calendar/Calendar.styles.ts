import { cva } from "class-variance-authority";

export const calendarVariants = cva(
    [
        "rounded-xl",
        "border",
        "bg-background",
        "p-4",
        "shadow-sm",
    ].join(" ")
);