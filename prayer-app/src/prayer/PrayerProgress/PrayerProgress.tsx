// New file generated from PrayerProgress.tsx
import { forwardRef } from "react";

import Typography  from "../../../primitives/Typography";
import { cn } from "../../utils/cn";

import {
    prayerProgressContentVariants,
    prayerProgressFooterVariants,
    prayerProgressHeaderVariants,
    prayerProgressVariants,
} from "./PrayerProgress.styles";

import type {
    PrayerProgressProps,
} from "./PrayerProgress.types";

const PrayerProgress = forwardRef<
    HTMLElement,
    PrayerProgressProps
>(({
    heading,
    description,
    progress,
    details,
    actions,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                prayerProgressVariants(),
                className
            )}
            {...props}
        >

            {(heading || description) && (

                <header
                    className={prayerProgressHeaderVariants()}
                >

                    {heading && (

                        <Typography variant="h4">

                            {heading}

                        </Typography>

                    )}

                    {description && (

                        <Typography
                            variant="body-sm"
                            color="muted"
                        >

                            {description}

                        </Typography>

                    )}

                </header>

            )}

            <div
                className={prayerProgressContentVariants()}
            >

                {progress}

                {details}

            </div>

            {actions && (

                <footer
                    className={prayerProgressFooterVariants()}
                >

                    {actions}

                </footer>

            )}

        </section>

    );

});

PrayerProgress.displayName =
    "PrayerProgress";

export default PrayerProgress;