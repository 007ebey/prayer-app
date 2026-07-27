import { cva } from "class-variance-authority";

export const splitViewVariants = cva(
    "flex w-full h-full",
    {
        variants: {
            direction: {
                horizontal: "flex-row",
                vertical: "flex-col",
            },

            reverse: {
                true: "flex-row-reverse",
                false: "",
            },
        },

        defaultVariants: {
            direction: "horizontal",
            reverse: false,
        },
    }
);

export const splitRatios = {
    "25-75": ["basis-1/4", "basis-3/4"],
    "30-70": ["basis-[30%]", "basis-[70%]"],
    "40-60": ["basis-2/5", "basis-3/5"],
    "50-50": ["basis-1/2", "basis-1/2"],
    "60-40": ["basis-3/5", "basis-2/5"],
    "70-30": ["basis-[70%]", "basis-[30%]"],
    "75-25": ["basis-3/4", "basis-1/4"],
} as const;