import {
    forwardRef,
    useState,
} from "react";

import { cn } from "../../../utils/cn";

import { imageVariants } from "./Image.styles";

import type { ImageProps } from "./Image.types";

const Image = forwardRef<
    HTMLImageElement,
    ImageProps
>(
    (
        {
            src,
            alt,
            fallback,
            fit,
            rounded,
            className,
            onError,
            ...props
        },
        ref
    ) => {
        const [
            hasError,
            setHasError,
        ] = useState(false);

        if (hasError && fallback) {
            return <>{fallback}</>;
        }

        return (
            <img
                ref={ref}
                src={src}
                alt={alt}
                className={cn(
                    imageVariants({
                        fit,
                        rounded,
                    }),
                    className
                )}
                onError={(event) => {
                    setHasError(true);
                    onError?.(event);
                }}
                {...props}
            />
        );
    }
);

Image.displayName = "Image";

export default Image;