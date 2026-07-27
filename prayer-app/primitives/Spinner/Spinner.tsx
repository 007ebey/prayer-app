import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { SpinnerProps } from "./Spinner.types";

import { spinnerVariants } from "./Spinner.styles";

const Spinner = forwardRef<
    HTMLDivElement,
    SpinnerProps
>(
    (
        {
            className,
            size,
            thickness,
            ...props
        },
        ref
    ) => (
        <div
            ref={ref}
            role="status"
            aria-label="Loading"
            className={cn(
                spinnerVariants({
                    size,
                    thickness,
                }),
                className
            )}
            {...props}
        >
            <span className="sr-only">
                Loading...
            </span>
        </div>
    )
);

Spinner.displayName = "Spinner";

export default Spinner;