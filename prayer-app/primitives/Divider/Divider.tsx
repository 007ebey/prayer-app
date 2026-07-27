import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { DividerProps } from "./Divider.types";

import { dividerVariants } from "./Divider.styles";

const Divider = forwardRef<
    HTMLDivElement,
    DividerProps
>(
(
    {
        orientation,
        spacing,
        thickness,
        label,
        className,
        ...props
    },
    ref
) => {

    if (label && orientation === "horizontal") {

        return (

            <div
                ref={ref}
                className={cn(
                    "flex items-center w-full",
                    spacing !== "none" && dividerVariants({
                        orientation,
                        spacing,
                        thickness
                    }).replace("border-t", "")
                )}
                {...props}
            >

                <div
                    className={cn(
                        "flex-1 border-t border-slate-200"
                    )}
                />

                <span
                    className="mx-4 text-sm text-slate-500 whitespace-nowrap"
                >
                    {label}
                </span>

                <div
                    className="flex-1 border-t border-slate-200"
                />

            </div>

        );

    }

    return (

        <div

            ref={ref}

            className={cn(

                dividerVariants({

                    orientation,

                    spacing,

                    thickness

                }),

                className

            )}

            {...props}

        />

    );

});

Divider.displayName = "Divider";

export default Divider;