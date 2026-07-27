import { cva } from "class-variance-authority";

export const flexVariants = cva(
    "",
    {
        variants: {

            inline: {

                true: "inline-flex",

                false: "flex"
            },

            direction: {

                row: "flex-row",

                column: "flex-col",

                "row-reverse": "flex-row-reverse",

                "column-reverse": "flex-col-reverse"
            },

            align: {

                start: "items-start",

                center: "items-center",

                end: "items-end",

                stretch: "items-stretch",

                baseline: "items-baseline"
            },

            justify: {

                start: "justify-start",

                center: "justify-center",

                end: "justify-end",

                between: "justify-between",

                around: "justify-around",

                evenly: "justify-evenly"
            },

            wrap: {

                nowrap: "flex-nowrap",

                wrap: "flex-wrap",

                "wrap-reverse": "flex-wrap-reverse"
            },

            gap: {

                none: "",

                xs: "gap-1",

                sm: "gap-2",

                md: "gap-4",

                lg: "gap-6",

                xl: "gap-8"
            }

        },

        defaultVariants: {

            inline: false,

            direction: "row",

            align: "stretch",

            justify: "start",

            wrap: "nowrap",

            gap: "none"
        }

    }
);