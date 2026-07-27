// New file generated from SessionHeader.tsx
import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    sessionHeaderActionsVariants,
    sessionHeaderIdentityVariants,
    sessionHeaderLeadingVariants,
    sessionHeaderMetadataVariants,
    sessionHeaderTrailingVariants,
    sessionHeaderVariants,
} from "./SessionHeader.styles";

import type {
    SessionHeaderProps,
} from "./SessionHeader.types";

const SessionHeader = forwardRef<
    HTMLElement,
    SessionHeaderProps
>(({
    leading,
    identity,
    metadata,
    actions,
    trailing,
    className,
    ...props
}, ref) => {

    return (

        <header
            ref={ref}
            className={cn(
                sessionHeaderVariants(),
                className
            )}
            {...props}
        >

            {leading && (

                <div
                    className={sessionHeaderLeadingVariants()}
                >

                    {leading}

                </div>

            )}

            <div
                className={sessionHeaderIdentityVariants()}
            >

                {identity}

                {metadata && (

                    <div
                        className={sessionHeaderMetadataVariants()}
                    >

                        {metadata}

                    </div>

                )}

            </div>

            {actions && (

                <div
                    className={sessionHeaderActionsVariants()}
                >

                    {actions}

                </div>

            )}

            {trailing && (

                <div
                    className={sessionHeaderTrailingVariants()}
                >

                    {trailing}

                </div>

            )}

        </header>

    );

});

SessionHeader.displayName =
    "SessionHeader";

export default SessionHeader;