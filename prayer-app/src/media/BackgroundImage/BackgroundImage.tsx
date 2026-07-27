import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import { backgroundImageVariants } from "./BackgroundImage.styles";

import type { BackgroundImageProps } from "./BackgroundImage.types";

const BackgroundImage = forwardRef<
    HTMLDivElement,
    BackgroundImageProps
>(
    (
        {
            src,
            children,
            overlay = false,
            overlayOpacity = 0.45,
            position,
            cover,
            rounded,
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
                    backgroundImageVariants({
                        position,
                        cover,
                        rounded,
                    }),
                    className
                )}
                style={{
                    backgroundImage: `url(${src})`,
                    backgroundSize: cover
                        ? "cover"
                        : "contain",
                    backgroundRepeat: "no-repeat",
                    ...style,
                }}
                {...props}
            >
                {overlay && (
                    <div
                        className="absolute inset-0 bg-black"
                        style={{
                            opacity:
                                overlayOpacity,
                        }}
                    />
                )}

                {children && (
                    <div className="relative z-10">
                        {children}
                    </div>
                )}
            </div>
        );
    }
);

BackgroundImage.displayName =
    "BackgroundImage";

export default BackgroundImage;