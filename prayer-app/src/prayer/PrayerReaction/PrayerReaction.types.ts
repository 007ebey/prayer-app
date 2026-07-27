// New file generated from PrayerReaction.tsx
import {
    ButtonHTMLAttributes,
    ReactNode,
} from "react";

export interface PrayerReactionProps
    extends Omit<
        ButtonHTMLAttributes<HTMLButtonElement>,
        "children"
    > {

    icon?: ReactNode;

    label?: ReactNode;

    count?: ReactNode;

    active?: boolean;

}