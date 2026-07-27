import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { SkeletonProps } from "./Skeleton.types";

import { skeletonVariants } from "./Skeleton.styles";

export const Skeleton = forwardRef<
    HTMLDivElement,
    SkeletonProps
>(
    (
        {
            className,
            variant = "rect",
            animation = "pulse",
            width,
            height,
            size,
            style,
            ...props
        },
        ref
    ) => {
        const styles = {
            ...style,
            width: size ?? width,
            height: size ?? height,
        };

        return (
            <div
                ref={ref}
                style={styles}
                className={cn(
                    skeletonVariants({
                        variant,
                        animation,
                    }),
                    className
                )}
                {...props}
            />
        );
    }
);

Skeleton.displayName = "Skeleton";

export default Skeleton;