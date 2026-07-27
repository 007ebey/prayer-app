// New file generated from SessionControls.tsx
import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    sessionControlsLeadingVariants,
    sessionControlsPrimaryVariants,
    sessionControlsSecondaryVariants,
    sessionControlsTrailingVariants,
    sessionControlsVariants,
} from "./SessionControls.styles";

import type {
    SessionControlsProps,
} from "./SessionControls.types";

const SessionControls = forwardRef<
    HTMLElement,
    SessionControlsProps
>(({
    leading,
    primaryActions,
    secondaryActions,
    trailing,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                sessionControlsVariants(),
                className
            )}
            {...props}
        >

            {leading && (

                <div
                    className={sessionControlsLeadingVariants()}
                >

                    {leading}

                </div>

            )}

            {primaryActions && (

                <div
                    className={sessionControlsPrimaryVariants()}
                >

                    {primaryActions}

                </div>

            )}

            {secondaryActions && (

                <div
                    className={sessionControlsSecondaryVariants()}
                >

                    {secondaryActions}

                </div>

            )}

            {trailing && (

                <div
                    className={sessionControlsTrailingVariants()}
                >

                    {trailing}

                </div>

            )}

        </section>

    );

});

SessionControls.displayName =
    "SessionControls";

export default SessionControls;