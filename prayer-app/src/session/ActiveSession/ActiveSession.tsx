// New file generated from ActiveSession.tsx
import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import {
    activeSessionContentVariants,
    activeSessionFooterVariants,
    activeSessionHeaderVariants,
    activeSessionVariants,
} from "./ActiveSession.styles";

import type {
    ActiveSessionProps,
} from "./ActiveSession.types";

const ActiveSession = forwardRef<
    HTMLElement,
    ActiveSessionProps
>(({
    header,
    status,
    timer,
    participants,
    agenda,
    actions,
    footer,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                activeSessionVariants(),
                className
            )}
            {...props}
        >

            {(header || status || timer) && (

                <header
                    className={activeSessionHeaderVariants()}
                >

                    {header}

                    {status}

                    {timer}

                </header>

            )}

            <div
                className={activeSessionContentVariants()}
            >

                {participants}

                {agenda}

                {actions}

            </div>

            {footer && (

                <footer
                    className={activeSessionFooterVariants()}
                >

                    {footer}

                </footer>

            )}

        </section>

    );

});

ActiveSession.displayName =
    "ActiveSession";

export default ActiveSession;