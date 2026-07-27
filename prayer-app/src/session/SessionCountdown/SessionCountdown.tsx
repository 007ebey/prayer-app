// New file generated from SessionCountdown.tsx
import { forwardRef } from "react";

import  Typography  from "../../../primitives/Typography";
import { cn } from "../../../utils/cn";

import {
    sessionCountdownActionsVariants,
    sessionCountdownContentVariants,
    sessionCountdownVariants,
} from "./SessionCountdown.styles";

import type {
    SessionCountdownProps,
} from "./SessionCountdown.types";

const SessionCountdown = forwardRef<
    HTMLElement,
    SessionCountdownProps
>(({
    heading,
    countdown,
    description,
    status,
    actions,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                sessionCountdownVariants(),
                className
            )}
            {...props}
        >

            <div
                className={sessionCountdownContentVariants()}
            >

                {heading && (

                    <Typography variant="caption">

                        {heading}

                    </Typography>

                )}

                <Typography variant="display-sm">

                    {countdown}

                </Typography>

                {description && (

                    <Typography variant="body-sm">

                        {description}

                    </Typography>

                )}

                {status}

            </div>

            {actions && (

                <div
                    className={sessionCountdownActionsVariants()}
                >

                    {actions}

                </div>

            )}

        </section>

    );

});

SessionCountdown.displayName =
    "SessionCountdown";

export default SessionCountdown;