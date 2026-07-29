import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    imageVariants,
    logoVariants,
    textVariants,
} from "./Logo.styles";

import type { LogoProps } from "./Logo.types";

const Logo = forwardRef<
    HTMLImageElement,
    LogoProps
>(
    (
        {
            src,
            alt = "Logo",
            name,
            icon,
            variant = "full",
            size = "md",
            href,
            className,
            ...props
        },
        ref
    ) => {

        const content = (

            <span
                className={cn(
                    logoVariants({
                        variant,
                        size,
                    }),
                    className
                )}
            >

                {(variant === "full" || variant === "icon") && (

                    src ? (

                        <img
                            ref={ref}
                            src={src}
                            alt={alt}
                            className={imageVariants({
                                size,
                            })}
                            {...props}
                        />

                    ) : (

                        icon

                    )

                )}

                {(variant === "full" || variant === "text") && name && (

                    <span
                        className={textVariants({
                            size,
                        })}
                    >
                        {name}
                    </span>

                )}

            </span>

        );

        if (href) {

            return (

                <a
                    href={href}
                    aria-label={typeof name === "string" ? name : "Home"}
                >
                    {content}
                </a>

            );

        }

        return content;

    }
);

Logo.displayName = "Logo";

export default Logo;