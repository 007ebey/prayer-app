// New file generated from PrayerCategory.tsx
import { forwardRef } from "react";

import Typography from "../../../primitives/Typography";
import { cn } from "../../../utils/cn";

import {
    prayerCategoryContentVariants,
    prayerCategoryHeaderVariants,
    prayerCategoryTitleVariants,
    prayerCategoryVariants,
} from "./PrayerCategory.styles";

import type {
    PrayerCategoryProps,
} from "./PrayerCategory.types";

const PrayerCategory = forwardRef<
    HTMLElement,
    PrayerCategoryProps
>(({
    heading,
    description,
    icon,
    actions,
    children,
    className,
    ...props
}, ref) => {

    return (

        <section
            ref={ref}
            className={cn(
                prayerCategoryVariants(),
                className
            )}
            {...props}
        >

            {(heading || actions) && (

                <header
                    className={prayerCategoryHeaderVariants()}
                >

                    <div
                        className={prayerCategoryTitleVariants()}
                    >

                        {icon}

                        <div>

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

                        </div>

                    </div>

                    {actions}

                </header>

            )}

            <div
                className={prayerCategoryContentVariants()}
            >

                {children}

            </div>

        </section>

    );

});

PrayerCategory.displayName = "PrayerCategory";

export default PrayerCategory;