import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { FlexProps } from "./Flex.types";

import { flexVariants } from "./Flex.styles";

const Flex = forwardRef<
    HTMLDivElement,
    FlexProps
>(
(
    {
        children,
        direction,
        align,
        justify,
        wrap,
        gap,
        inline,
        className,
        ...props
    },
    ref
) => {

    return (

        <div

            ref={ref}

            className={cn(

                flexVariants({

                    direction,

                    align,

                    justify,

                    wrap,

                    gap,

                    inline

                }),

                className

            )}

            {...props}

        >

            {children}

        </div>

    );

});

Flex.displayName = "Flex";

export default Flex;