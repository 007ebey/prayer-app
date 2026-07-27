import { forwardRef } from "react";

import type { ContainerProps } from "./Container.types";

import {
    containerSizes,
    containerPadding
} from "./Container.styles";

import { cn } from "../../utils/cn";

const Container = forwardRef<HTMLDivElement, ContainerProps>(
    (
        {
            children,
            size = "xl",
            centered = true,
            fluid = false,
            padding = true,
            className,
            ...props
        },
        ref
    ) => {

        return (

            <div
                ref={ref}
                className={cn(

                    fluid
                        ? "w-full"
                        : containerSizes[size],

                    centered && "mx-auto",

                    padding
                        ? containerPadding.enabled
                        : containerPadding.disabled,

                    className
                )}

                {...props}
            >

                {children}

            </div>

        );

    }
);

Container.displayName = "Container";

export default Container;