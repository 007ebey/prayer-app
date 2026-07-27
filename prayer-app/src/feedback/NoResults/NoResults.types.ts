// New file generated from NoResults.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface NoResultsProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    query?: string;

    heading?: ReactNode;

    description?: ReactNode;

    actions?: ReactNode;

    icon?: ReactNode;
}