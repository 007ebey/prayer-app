import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ActivityStatus =
    | "success"
    | "warning"
    | "danger"
    | "info"
    | "neutral";

export interface ActivityCardProps
    extends HTMLAttributes<HTMLDivElement> {

    icon?: ReactNode;

    heading: ReactNode;

    description?: ReactNode;

    timestamp?: ReactNode;

    status?: ActivityStatus;

    actions?: ReactNode;
}