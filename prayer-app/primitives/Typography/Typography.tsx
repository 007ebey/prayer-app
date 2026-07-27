import { createElement, forwardRef } from "react";

import { cn } from "../../utils/cn";

import { typographyVariants } from "./Typography.styles";

import type { TypographyProps } from "./Typography.types";

const defaultElements = {
    display: "h1",
    "display-sm": "h1",

    h1: "h1",
    h2: "h2",
    h3: "h3",
    h4: "h4",

    title: "h5",
    subtitle: "h6",

    "body-lg": "p",
    body: "p",
    "body-sm": "p",

    caption: "span",
    overline: "span",
} as const;

const Typography = forwardRef<
    HTMLElement,
    TypographyProps
>(
    (
        {
            as,
            variant = "body",
            weight,
            align,
            truncate,
            className,
            children,
            ...props
        },
        ref
    ) => {
        const Component =
            as ??
            defaultElements[variant] ??
            "p";

        return createElement(
            Component,
            {
                ...props,
                ref: ref as React.Ref<HTMLElement>,
                className: cn(
                    typographyVariants({
                        variant,
                        weight,
                        align,
                        truncate,
                    }),
                    className
                ),
            },
            children
        );
    }
);

Typography.displayName = "Typography";

export default Typography;