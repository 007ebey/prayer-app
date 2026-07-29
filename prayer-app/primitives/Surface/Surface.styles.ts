import { cva } from "class-variance-authority";

export const surfaceVariants = cva(
    "transition-all duration-200",
    {
        variants: {

            variant: {

                default: "bg-surface",

                muted: "bg-surface-muted",

                primary: "bg-primary text-primary-foreground",

                secondary: "bg-secondary",

                transparent: "bg-transparent",

            },

            elevation: {

                none: "",

                flat: "shadow-none",

                raised: "shadow-sm",

                floating: "shadow-md",

                overlay: "shadow-lg",

                modal: "shadow-2xl",

            },

            radius: {

                none: "rounded-none",

                sm: "rounded-sm",

                md: "rounded-md",

                lg: "rounded-lg",

                xl: "rounded-xl",

                "2xl": "rounded-2xl",

                full: "rounded-full",

            },

            padding: {

                none: "p-0",

                xs: "p-2",

                sm: "p-3",

                md: "p-4",

                lg: "p-6",

                xl: "p-8",

            },

            bordered: {

                true: "border border-border",

                false: "",

            },

            hoverable: {

                true: "hover:shadow-lg",

                false: "",

            },

            interactive: {

                true: "cursor-pointer active:scale-[0.99]",

                false: "",

            },

            fullWidth: {

                true: "w-full",

                false: "inline-block",

            },

        },

        defaultVariants: {

            variant: "default",

            elevation: "flat",

            radius: "lg",

            padding: "md",

            bordered: false,

            hoverable: false,

            interactive: false,

            fullWidth: true,

        },

    }
);