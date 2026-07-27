import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import { containerVariants } from "./Container.styles";

import type { ContainerProps } from "./Container.types";

const Container = forwardRef<
    HTMLDivElement,
    ContainerProps
>(
    (
        {
            children,
            size,
            centered,
            fluid,
            className,
            ...props
        },
        ref
    ) => (
        <div
            ref={ref}
            className={cn(
                containerVariants({
                    size,
                    centered,
                    fluid,
                }),
                className
            )}
            {...props}
        >
            {children}
        </div>
    )
);

Container.displayName = "Container";

export default Container;