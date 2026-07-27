import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { CardProps } from "./Card.types";

import { cardVariants } from "./Card.styles";

const Card = forwardRef<
    HTMLDivElement,
    CardProps
>(

(
    {
        children,

        variant,

        padding,

        radius,

        hoverable,

        clickable,

        className,

        ...props
    },

    ref

) => {

    return (

        <div

            ref={ref}

            className={cn(

                cardVariants({

                    variant,

                    padding,

                    radius,

                    hoverable,

                    clickable

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

Card.displayName = "Card";

export default Card;