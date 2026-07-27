// New file generated from NoResults.tsx
import {
    SearchX,
} from "lucide-react";

import { EmptyState } from "../EmptyState";

import { cn } from "../../../utils/cn";

import {
    noResultsVariants,
} from "./NoResults.styles";

import type {
    NoResultsProps,
} from "./NoResults.types";

const NoResults = ({
    query,
    heading,
    description,
    actions,
    icon,
    className,
    ...props
}: NoResultsProps) => {

    const defaultHeading =
        heading ??
        (query
            ? `No results found for "${query}"`
            : "No results found");

    const defaultDescription =
        description ??
        "Try adjusting your search or filters and try again.";

    return (
        <EmptyState
            icon={
                icon ?? (
                    <SearchX
                        size={36}
                    />
                )
            }
            heading={defaultHeading}
            description={defaultDescription}
            actions={actions}
            className={cn(
                noResultsVariants(),
                className
            )}
            {...props}
        />
    );
};

export default NoResults;