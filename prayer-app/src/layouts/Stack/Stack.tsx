import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import { stackVariants } from "./Stack.styles";

import type { StackProps } from "./Stack.types";

const Stack = forwardRef<
    HTMLDivElement,
    StackProps
>(
    (
        {
            children,
            gap,
            align,
            justify,
            reverse,
            className,
            ...props
        },
        ref
    ) => {
        return (
            <div
                ref={ref}
                className={cn(
                    stackVariants({
                        gap,
                        align,
                        justify,
                        reverse,
                    }),
                    className
                )}
                {...props}
            >
                {children}
            </div>
        );
    }
);

Stack.displayName = "Stack";

export default Stack;