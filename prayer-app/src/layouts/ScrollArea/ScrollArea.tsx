import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import { scrollAreaVariants } from "./ScrollArea.styles";

import type { ScrollAreaProps } from "./ScrollArea.types";

const ScrollArea = forwardRef<
    HTMLDivElement,
    ScrollAreaProps
>(
    (
        {
            children,
            direction,
            maxHeight,
            maxWidth,
            className,
            style,
            ...props
        },
        ref
    ) => (
        <div
            ref={ref}
            style={{
                ...style,
                maxHeight,
                maxWidth,
            }}
            className={cn(
                scrollAreaVariants({
                    direction,
                }),
                className
            )}
            {...props}
        >
            {children}
        </div>
    )
);

ScrollArea.displayName = "ScrollArea";

export default ScrollArea;