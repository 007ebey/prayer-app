import { cva } from "class-variance-authority";

export const typographyVariants = cva("", {
    variants: {
        variant: {
            display: "text-6xl font-bold tracking-tight",

            "display-sm": "text-5xl font-bold tracking-tight",

            h1: "text-5xl font-bold tracking-tight",
            h2: "text-4xl font-bold tracking-tight",
            h3: "text-3xl font-semibold",
            h4: "text-2xl font-semibold",

            title: "text-xl font-semibold",
            subtitle: "text-lg text-muted-foreground",

            "body-lg": "text-lg",
            body: "text-base",
            "body-sm": "text-sm",

            caption: "text-xs text-muted-foreground",

            overline:
                "text-xs uppercase tracking-widest text-muted-foreground",
        },

        weight: {
            light: "font-light",
            normal: "font-normal",
            medium: "font-medium",
            semibold: "font-semibold",
            bold: "font-bold",
        },

        align: {
            left: "text-left",
            center: "text-center",
            right: "text-right",
            justify: "text-justify",
        },

        truncate: {
            true: "truncate",
            false: "",
        },
    },

    defaultVariants: {
        variant: "body",
        weight: "normal",
        align: "left",
        truncate: false,
    },
});