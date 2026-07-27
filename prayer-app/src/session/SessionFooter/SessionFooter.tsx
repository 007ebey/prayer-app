// New file generated from SessionFooter.tsx
import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import {
    sessionFooterActionsVariants,
    sessionFooterLeadingVariants,
    sessionFooterMetadataVariants,
    sessionFooterTrailingVariants,
    sessionFooterVariants,
} from "./SessionFooter.styles";

import type {
    SessionFooterProps,
} from "./SessionFooter.types";

const SessionFooter = forwardRef<
    HTMLElement,
    SessionFooterProps
>(({
    leading,
    metadata,
    actions,
    trailing,
    className,
    ...props
}, ref) => {

    return (

        <footer
            ref={ref}
            className={cn(
                sessionFooterVariants(),
                className
            )}
            {...props}
        >

            {leading && (

                <div
                    className={sessionFooterLeadingVariants()}
                >

                    {leading}

                </div>

            )}

            {metadata && (

                <div
                    className={sessionFooterMetadataVariants()}
                >

                    {metadata}

                </div>

            )}

            {actions && (

                <div
                    className={sessionFooterActionsVariants()}
                >

                    {actions}

                </div>

            )}

            {trailing && (

                <div
                    className={sessionFooterTrailingVariants()}
                >

                    {trailing}

                </div>

            )}

        </footer>

    );

});

SessionFooter.displayName =
    "SessionFooter";

export default SessionFooter;