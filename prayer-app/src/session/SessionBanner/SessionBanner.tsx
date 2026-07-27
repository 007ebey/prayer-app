// New file generated from SessionBanner.tsx
import { forwardRef } from "react";

import Typography  from "../../../primitives/Typography";
import { cn } from "../../../utils/cn";

import {
    sessionBannerActionsVariants,
    sessionBannerBackgroundVariants,
    sessionBannerContentVariants,
    sessionBannerMetadataVariants,
    sessionBannerOverlayVariants,
    sessionBannerVariants,
} from "./SessionBanner.styles";

import type {
    SessionBannerProps,
} from "./SessionBanner.types";

const SessionBanner = forwardRef<
    HTMLElement,
    SessionBannerProps
>(({
    background,
    overlay,
    badge,
    heading,
    description,
    metadata,
    actions,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                sessionBannerVariants(),
                className
            )}
            {...props}
        >

            {background && (

                <div
                    className={sessionBannerBackgroundVariants()}
                >
                    {background}
                </div>

            )}

            <div
                className={sessionBannerOverlayVariants()}
            >
                {overlay}
            </div>

            <div
                className={sessionBannerContentVariants()}
            >

                {badge}

                {heading && (

                    <Typography variant="display">

                        {heading}

                    </Typography>

                )}

                {description && (

                    <Typography variant="body-lg">

                        {description}

                    </Typography>

                )}

                {metadata && (

                    <div
                        className={sessionBannerMetadataVariants()}
                    >

                        {metadata}

                    </div>

                )}

                {actions && (

                    <div
                        className={sessionBannerActionsVariants()}
                    >

                        {actions}

                    </div>

                )}

            </div>

        </section>

    );

});

SessionBanner.displayName =
    "SessionBanner";

export default SessionBanner;