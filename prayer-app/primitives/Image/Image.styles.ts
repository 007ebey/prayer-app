import { cva } from "class-variance-authority";

export const imageVariants = cva(
    "block w-full",
    {
        variants: {
            fit: {
                cover: "object-cover",
                contain: "object-contain",
                fill: "object-fill",
                none: "object-none",
                "scale-down": "object-scale-down",
            },

            radius: {
                none: "",
                sm: "rounded-sm",
                md: "rounded-md",
                lg: "rounded-lg",
                xl: "rounded-xl",
                full: "rounded-full",
            },

            aspectRatio: {
                auto: "",
                square: "aspect-square",
                video: "aspect-video",
                portrait: "aspect-[3/4]",
            },
        },

        defaultVariants: {
            fit: "cover",
            radius: "md",
            aspectRatio: "auto",
        },
    }
);