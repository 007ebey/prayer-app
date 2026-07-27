import { forwardRef } from "react";

import Flex from "../Flex";

import type { StackProps } from "./Stack.types";

import { stackVariants } from "./Stack.styles";

import { cn } from "../../utils/cn";

const Stack = forwardRef<
    HTMLDivElement,
    StackProps
>(
(
    {
        children,

        gap = "none",

        align = "stretch",

        justify = "start",

        reverse = false,

        className,

        ...props
    },

    ref

) => {

    return (

        <Flex

            ref={ref}

            direction={
                reverse
                    ? "column-reverse"
                    : "column"
            }

            gap={gap}

            align={align}

            justify={justify}

            className={cn(

                stackVariants({

                    reverse

                }),

                className

            )}

            {...props}

        >

            {children}

        </Flex>

    );

});

Stack.displayName = "Stack";

export default Stack;