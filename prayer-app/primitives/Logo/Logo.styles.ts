import { cva } from "class-variance-authority";

export const logoVariants = cva(
    "inline-flex items-center select-none",
    {
        variants: {

            variant: {
                full: "",
                icon: "",
                text: "",
            },

            size: {

                xs: "gap-1",

                sm: "gap-2",

                md: "gap-3",

                lg: "gap-4",

                xl: "gap-5",

            },

        },

        defaultVariants: {

            variant: "full",

            size: "md",

        },

    }
);

export const imageVariants = cva(
    "object-contain",
    {
        variants: {

            size: {

                xs: "h-5 w-5",

                sm: "h-6 w-6",

                md: "h-8 w-8",

                lg: "h-10 w-10",

                xl: "h-14 w-14",

            },

        },

        defaultVariants: {

            size: "md",

        },

    }
);

export const textVariants = cva(
    "font-bold tracking-tight",
    {
        variants: {

            size: {

                xs: "text-sm",

                sm: "text-base",

                md: "text-xl",

                lg: "text-2xl",

                xl: "text-4xl",

            },

        },

        defaultVariants: {

            size: "md",

        },

    }
);