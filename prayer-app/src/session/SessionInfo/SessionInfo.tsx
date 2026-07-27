// New file generated from SessionInfo.tsx
import { forwardRef } from "react";

import Typography  from "../../../primitives/Typography";

import { cn } from "../../utils/cn";

import {
    sessionInfoDescriptionVariants,
    sessionInfoHeadingVariants,
    sessionInfoMetadataVariants,
    sessionInfoVariants,
} from "./SessionInfo.styles";

import type {
    SessionInfoProps,
} from "./SessionInfo.types";

const SessionInfo = forwardRef<
    HTMLElement,
    SessionInfoProps
>(({
    badge,
    heading,
    description,
    metadata,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                sessionInfoVariants(),
                className
            )}
            {...props}
        >

            {(badge || heading) && (

                <div
                    className={sessionInfoHeadingVariants()}
                >

                    {badge}

                    {heading && (

                        <Typography
                            variant="h3"
                        >

                            {heading}

                        </Typography>

                    )}

                </div>

            )}

            {description && (

                <Typography
                    variant="body"
                    className={sessionInfoDescriptionVariants()}
                >

                    {description}

                </Typography>

            )}

            {metadata && (

                <div
                    className={sessionInfoMetadataVariants()}
                >

                    {metadata}

                </div>

            )}

        </section>

    );

});

SessionInfo.displayName =
    "SessionInfo";

export default SessionInfo;