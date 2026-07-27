import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import { gridVariants } from "./Grid.styles";

import type { GridProps } from "./Grid.types";

const Grid = forwardRef<
    HTMLDivElement,
    GridProps
>(
    (
        {
            children,
            columns,
            gap,
            className,
            ...props
        },
        ref
    ) => (
        <div
            ref={ref}
            className={cn(
                gridVariants({
                    columns,
                    gap,
                }),
                className
            )}
            {...props}
        >
            {children}
        </div>
    )
);

Grid.displayName = "Grid";

export default Grid;