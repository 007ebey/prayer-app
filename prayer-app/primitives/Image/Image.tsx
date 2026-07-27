import {
    forwardRef,
    useState,
} from "react";

import { cn } from "../../utils/cn";

import { imageVariants } from "./Image.styles";

import type { ImageProps } from "./Image.types";

const Image = forwardRef<
    HTMLImageElement,
    ImageProps
>(
    (
        {
            src,
            alt = "",
            fit,
            radius,
            aspectRatio,
            fallback,
            loadingIndicator,
            className,
            onLoad,
            onError,
            ...props
        },
        ref
    ) => {

        const [loaded, setLoaded] =
            useState(false);

        const [failed, setFailed] =
            useState(false);

        if (failed && fallback) {
            return <>{fallback}</>;
        }

        return (
            <div className="relative overflow-hidden">

                {!loaded &&
                    loadingIndicator}

                <img
                    ref={ref}
                    src={src}
                    alt={alt}
                    loading="lazy"
                    className={cn(
                        imageVariants({
                            fit,
                            radius,
                            aspectRatio,
                        }),
                        !loaded &&
                            "opacity-0",
                        loaded &&
                            "opacity-100 transition-opacity duration-300",
                        className
                    )}
                    onLoad={(e) => {
                        setLoaded(true);
                        onLoad?.(e);
                    }}
                    onError={(e) => {
                        setFailed(true);
                        onError?.(e);
                    }}
                    {...props}
                />

            </div>
        );
    }
);

Image.displayName = "Image";

export default Image;