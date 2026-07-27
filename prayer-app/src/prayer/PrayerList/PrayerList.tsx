// New file generated from PrayerList.tsx
import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import {
    prayerListContentVariants,
    prayerListFooterVariants,
    prayerListHeaderVariants,
    prayerListVariants,
} from "./PrayerList.styles";

import type {
    PrayerListProps,
} from "./PrayerList.types";

const PrayerList = forwardRef<
    HTMLElement,
    PrayerListProps
>(({
    header,
    toolbar,
    children,
    footer,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                prayerListVariants(),
                className
            )}
            {...props}
        >

            {(header || toolbar) && (

                <header
                    className={prayerListHeaderVariants()}
                >

                    {header}

                    {toolbar}

                </header>

            )}

            <div
                className={prayerListContentVariants()}
            >

                {children}

            </div>

            {footer && (

                <footer
                    className={prayerListFooterVariants()}
                >

                    {footer}

                </footer>

            )}

        </section>

    );

});

PrayerList.displayName =
    "PrayerList";

export default PrayerList;