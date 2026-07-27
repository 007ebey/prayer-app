// New file generated from EmptyState.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type EmptyStateSize =
    | "sm"
    | "md"
    | "lg";

export interface EmptyStateProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    illustration?: ReactNode;

    icon?: ReactNode;

    heading: ReactNode;

    description?: ReactNode;

    actions?: ReactNode;

    footer?: ReactNode;

    size?: EmptyStateSize;
}