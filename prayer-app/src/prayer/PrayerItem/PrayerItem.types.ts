import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface PrayerItemProps
    extends HTMLAttributes<HTMLElement> {

    icon?: ReactNode;

    heading?: ReactNode;

    description?: ReactNode;

    scripture?: ReactNode;

    badges?: ReactNode;

    actions?: ReactNode;

    footer?: ReactNode;
}