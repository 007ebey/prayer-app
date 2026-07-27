// New file generated from PrayerWall.tsx
import { cn } from "../../utils/cn";

import {
    prayerWallContentVariants,
    prayerWallVariants,
} from "./PrayerWall.styles";

import type {
    PrayerWallProps,
} from "./PrayerWall.types";

const PrayerWall = ({
    background,
    overlay,
    header,
    content,
    media,
    scripture,
    footer,
}: PrayerWallProps) => {

    return (

        <main
            className={cn(
                prayerWallVariants()
            )}
        >

            {background}

            {overlay}

            <section
                className={cn(
                    prayerWallContentVariants()
                )}
            >

                {header}

                {media}

                {content}

                {scripture}

                {footer}

            </section>

        </main>

    );

};

export default PrayerWall;