import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { BadgeProps } from "./Badge.types";

import { badgeVariants } from "./Badge.styles";

const Badge = forwardRef<HTMLSpanElement, BadgeProps>(
(
    {
        children,
        variant,
        color,
        size,
        rounded,
        leftIcon,
        rightIcon,
        className,
        ...props
    },
    ref
) => {

    return (

        <span

            ref={ref}

            className={cn(

                badgeVariants({

                    variant,

                    color,

                    size,

                    rounded

                }),

                className

            )}

            {...props}

        >

            {leftIcon && (
                <span className="mr-1 flex">
                    {leftIcon}
                </span>
            )}

            {children}

            {rightIcon && (
                <span className="ml-1 flex">
                    {rightIcon}
                </span>
            )}

        </span>

    );

});

Badge.displayName = "Badge";

export default Badge;