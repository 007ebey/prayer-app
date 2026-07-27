// New file generated from PrayerItem.tsx
import { forwardRef } from "react";

import Typography from "../../../primitives/Typography";
import { cn } from "../../../utils/cn";

import {
    prayerItemBodyVariants,
    prayerItemContentVariants,
    prayerItemFooterVariants,
    prayerItemVariants,
} from "./PrayerItem.styles";

import type {
    PrayerItemProps,
} from "./PrayerItem.types";

const PrayerItem = forwardRef<
    HTMLElement,
    PrayerItemProps
>(({
    icon,
    title,
    description,
    scripture,
    badges,
    actions,
    footer,
    className,
    children,
    ...props
}, ref) => {

    return (

        <article
            ref={ref}
            className={cn(
                prayerItemVariants(),
                className
            )}
            {...props}
        >

            <div
                className={prayerItemContentVariants()}
            >

                {icon}

                <div
                    className={prayerItemBodyVariants()}
                >

                    {title && (

                        <Typography variant="h4">

                            {title}

                        </Typography>

                    )}

                    {description && (

                        <Typography variant="body">

                            {description}

                        </Typography>

                    )}

                    {scripture}

                    {badges}

                    {children}

                    {actions}

                </div>

            </div>

            {footer && (

                <footer
                    className={prayerItemFooterVariants()}
                >

                    {footer}

                </footer>

            )}

        </article>

    );

});

PrayerItem.displayName = "PrayerItem";

export default PrayerItem;