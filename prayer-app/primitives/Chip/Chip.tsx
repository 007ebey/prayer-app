import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { ChipProps } from "./Chip.types";

import { chipVariants } from "./Chip.styles";

const Chip = forwardRef<
    HTMLDivElement,
    ChipProps
>(

(
    {
        children,

        variant,

        color,

        size,

        selected,

        removable,

        disabled,

        leftIcon,

        rightIcon,

        onRemove,

        className,

        ...props
    },

    ref

) => {

    return (

        <div

            ref={ref}

            className={cn(

                chipVariants({

                    variant,

                    color,

                    size,

                    selected,

                    disabled

                }),

                className

            )}

            {...props}

        >

            {leftIcon}

            <span>

                {children}

            </span>

            {rightIcon}

            {removable && (

                <button

                    type="button"

                    onClick={onRemove}

                    className="ml-1 rounded-full p-1 hover:bg-black/10"

                >

                    ✕

                </button>

            )}

        </div>

    );

}

);

Chip.displayName = "Chip";

export default Chip;