// New file generated from PrayerWall.tsx
import { forwardRef } from "react";

import { cn } from "../../../utils/cn";

import {
    prayerWallBackgroundVariants,
    prayerWallContainerVariants,
    prayerWallContentVariants,
    prayerWallFooterVariants,
    prayerWallHeaderVariants,
    prayerWallOverlayVariants,
    prayerWallVariants,
} from "./PrayerWall.styles";

import type {
    PrayerWallProps,
} from "./PrayerWall.types";

const PrayerWall = forwardRef<
    HTMLElement,
    PrayerWallProps
>(({
    background,
    overlay,
    header,
    timer,
    progress,
    content,
    footer,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                prayerWallVariants(),
                className
            )}
            {...props}
        >

            {background && (

                <div
                    className={prayerWallBackgroundVariants()}
                >
                    {background}
                </div>

            )}

            {overlay && (

                <div
                    className={prayerWallOverlayVariants()}
                >
                    {overlay}
                </div>

            )}

            <div
                className={prayerWallContainerVariants()}
            >

                {(header || timer || progress) && (

                    <header
                        className={prayerWallHeaderVariants()}
                    >

                        {header}

                        {timer}

                        {progress}

                    </header>

                )}

                <main
                    className={prayerWallContentVariants()}
                >

                    {content}

                </main>

                {footer && (

                    <footer
                        className={prayerWallFooterVariants()}
                    >

                        {footer}

                    </footer>

                )}

            </div>

        </section>

    );

});

PrayerWall.displayName =
    "PrayerWall";

export default PrayerWall;