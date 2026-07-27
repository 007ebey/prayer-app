// New file generated from PrayerHostPanel.tsx
import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import {
    prayerHostPanelContentVariants,
    prayerHostPanelFooterVariants,
    prayerHostPanelHeaderVariants,
    prayerHostPanelVariants,
} from "./PrayerHostPanel.styles";

import type {
    PrayerHostPanelProps,
} from "./PrayerHostPanel.types";

const PrayerHostPanel = forwardRef<
    HTMLElement,
    PrayerHostPanelProps
>(({
    header,
    sessionControls,
    agendaControls,
    participantControls,
    mediaControls,
    quickActions,
    footer,
    className,
    children,
    ...props
}, ref) => {

    return (

        <aside
            ref={ref}
            className={cn(
                prayerHostPanelVariants(),
                className
            )}
            {...props}
        >

            {header && (

                <header
                    className={prayerHostPanelHeaderVariants()}
                >

                    {header}

                </header>

            )}

            <div
                className={prayerHostPanelContentVariants()}
            >

                {sessionControls}

                {agendaControls}

                {participantControls}

                {mediaControls}

                {quickActions}

                {children}

            </div>

            {footer && (

                <footer
                    className={prayerHostPanelFooterVariants()}
                >

                    {footer}

                </footer>

            )}

        </aside>

    );

});

PrayerHostPanel.displayName =
    "PrayerHostPanel";

export default PrayerHostPanel;