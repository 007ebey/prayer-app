import { cva } from "class-variance-authority";

export const gridVariants = cva(
    "grid",
    {
        variants: {

            columns: {

                1: "grid-cols-1",

                2: "grid-cols-2",

                3: "grid-cols-3",

                4: "grid-cols-4",

                5: "grid-cols-5",

                6: "grid-cols-6",

                12: "grid-cols-12"

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

            columns: 1,

            gap: "md"

        }

    }
);