import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import { surfaceVariants } from "./Surface.styles";

import type { SurfaceProps } from "./Surface.types";

const Surface = forwardRef<
    HTMLElement,
    SurfaceProps
>(
    (
        {
            children,
            className,

            variant,
            elevation,
            radius,
            padding,

            bordered,
            hoverable,
            interactive,
            fullWidth,

            ...props
        },
        ref
    ) => {

        return (

            <section
                ref={ref}
                className={cn(
                    surfaceVariants({

                        variant,

                        elevation,

                        radius,

                        padding,

                        bordered,

                        hoverable,

                        interactive,

                        fullWidth,

                    }),
                    className
                )}
                {...props}
            >

                {children}

            </section>

        );

    }
);

Surface.displayName = "Surface";

export default Surface;