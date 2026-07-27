import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import { gradientOverlayVariants } from "./GradientOverlay.styles";

import type { GradientOverlayProps, GradientDirection } from "./GradientOverlay.types";

const GradientOverlay = forwardRef<
    HTMLDivElement,
    GradientOverlayProps
>(
    (
        {
            children,
            direction,
            from = "#000000",
            to = "transparent",
            opacity = 0.6,
            className,
            style,
            ...props
        },
        ref
    ) => {
        return (
            <div
                ref={ref}
                className={cn(
                    gradientOverlayVariants({
                        direction,
                    }),
                    className
                )}
                style={{
                    ...style,
                    opacity,
                    backgroundImage: `linear-gradient(${getDirection(
                        direction
                    )}, ${from}, ${to})`,
                }}
                {...props}
            >
                {children && (
                    <div className="relative z-10 h-full w-full">
                        {children}
                    </div>
                )}
            </div>
        );
    }
);

GradientOverlay.displayName =
    "GradientOverlay";

function getDirection(
    direction?: GradientDirection
) {
    switch (direction) {
        case "top":
            return "to top";

        case "left":
            return "to left";

        case "right":
            return "to right";

        case "top-left":
            return "to top left";

        case "top-right":
            return "to top right";

        case "bottom-left":
            return "to bottom left";

        case "bottom-right":
            return "to bottom right";

        case "bottom":
        default:
            return "to bottom";
    }
}

export default GradientOverlay;