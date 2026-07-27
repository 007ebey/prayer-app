// New file generated from Drawer.tsx
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

export const drawerVariants = cva(
    [
        "fixed",
        "z-50",
        "bg-background",
        "border",
        "shadow-2xl",
        "flex",
        "flex-col",
        "overflow-hidden",
        "animate-in",
        "duration-300",
    ].join(" "),
    {
        variants: {
            side: {
                left:
                    "left-0 top-0 h-full slide-in-from-left",

                right:
                    "right-0 top-0 h-full slide-in-from-right",

                top:
                    "top-0 left-0 w-full slide-in-from-top",

                bottom:
                    "bottom-0 left-0 w-full slide-in-from-bottom rounded-t-2xl",
            },

            size: {
                sm: "",
                md: "",
                lg: "",
                full: "",
            },
        },

        compoundVariants: [
            {
                side: ["left", "right"],
                size: "sm",
                class: "w-72",
            },
            {
                side: ["left", "right"],
                size: "md",
                class: "w-96",
            },
            {
                side: ["left", "right"],
                size: "lg",
                class: "w-[32rem]",
            },
            {
                side: ["left", "right"],
                size: "full",
                class: "w-screen",
            },
            {
                side: ["top", "bottom"],
                size: "sm",
                class: "h-56",
            },
            {
                side: ["top", "bottom"],
                size: "md",
                class: "h-80",
            },
            {
                side: ["top", "bottom"],
                size: "lg",
                class: "h-[32rem]",
            },
            {
                side: ["top", "bottom"],
                size: "full",
                class: "h-screen",
            },
        ],

        defaultVariants: {
            side: "right",
            size: "md",
        },
    }
);