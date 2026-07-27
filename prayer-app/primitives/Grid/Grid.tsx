import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { GridProps } from "./Grid.types";

import { gridVariants } from "./Grid.styles";

const Grid = forwardRef<
    HTMLDivElement,
    GridProps
>(
(
    {
        children,

        columns = 1,

        gap,

        sm,

        md,

        lg,

        xl,

        autoFit = false,

        minWidth = "250px",

        className,

        style,

        ...props
    },

    ref

) => {

    const responsive = cn(

        sm && `sm:grid-cols-${sm}`,

        md && `md:grid-cols-${md}`,

        lg && `lg:grid-cols-${lg}`,

        xl && `xl:grid-cols-${xl}`

    );

    return (

        <div

            ref={ref}

            style={
                autoFit
                    ? {
                          ...style,
                          gridTemplateColumns:
                              `repeat(auto-fit, minmax(${minWidth}, 1fr))`
                      }
                    : style
            }

            className={cn(

                gridVariants({

                    columns,

                    gap

                }),

                responsive,

                className

            )}

            {...props}

        >

            {children}

        </div>

    );

});

Grid.displayName = "Grid";

export default Grid;