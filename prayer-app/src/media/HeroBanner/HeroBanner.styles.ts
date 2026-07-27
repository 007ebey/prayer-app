import { cva } from "class-variance-authority";

export const heroBannerVariants = cva(
    "relative overflow-hidden",
    {
        variants: {

            alignment: {
                left: "text-left items-start",
                center: "text-center items-center",
                right: "text-right items-end",
            },

            fullHeight: {
                true: "min-h-screen",
                false: "min-h-[32rem]",
            },

        },

        defaultVariants: {
            alignment: "left",
            fullHeight: false,
        },
    }
);