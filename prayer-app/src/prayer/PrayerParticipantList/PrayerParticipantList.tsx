// New file generated from PrayerParticipantList.tsx
import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    prayerParticipantListContentVariants,
    prayerParticipantListFooterVariants,
    prayerParticipantListHeaderVariants,
    prayerParticipantListVariants,
} from "./PrayerParticipantList.styles";

import type {
    PrayerParticipantListProps,
} from "./PrayerParticipantList.types";

const PrayerParticipantList = forwardRef<
    HTMLElement,
    PrayerParticipantListProps
>(({
    header,
    toolbar,
    children,
    emptyState,
    footer,
    className,
    ...props
}, ref) => {

    const hasParticipants =
        children !== undefined &&
        children !== null;

    return (

        <section
            ref={ref}
            className={cn(
                prayerParticipantListVariants(),
                className
            )}
            {...props}
        >

            {(header || toolbar) && (

                <header
                    className={prayerParticipantListHeaderVariants()}
                >

                    {header}

                    {toolbar}

                </header>

            )}

            <div
                className={prayerParticipantListContentVariants()}
            >

                {hasParticipants
                    ? children
                    : emptyState}

            </div>

            {footer && (

                <footer
                    className={prayerParticipantListFooterVariants()}
                >

                    {footer}

                </footer>

            )}

        </section>

    );

});

PrayerParticipantList.displayName =
    "PrayerParticipantList";

export default PrayerParticipantList;