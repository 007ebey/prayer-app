import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    splitRatios,
    splitViewVariants,
} from "./SplitView.styles";

import type { SplitViewProps } from "./SplitView.types";

const SplitView = forwardRef<
    HTMLDivElement,
    SplitViewProps
>(
    (
        {
            primary,
            secondary,
            direction,
            ratio = "50-50",
            reverse = false,
            className,
            ...props
        },
        ref
    ) => {
        const [primarySize, secondarySize] =
            splitRatios[ratio];

        return (
            <div
                ref={ref}
                className={cn(
                    splitViewVariants({
                        direction,
                        reverse,
                    }),
                    className
                )}
                {...props}
            >
                <div
                    className={cn(
                        primarySize,
                        "overflow-auto"
                    )}
                >
                    {primary}
                </div>

                <div
                    className={cn(
                        secondarySize,
                        "overflow-auto"
                    )}
                >
                    {secondary}
                </div>
            </div>
        );
    }
);

SplitView.displayName = "SplitView";

export default SplitView;