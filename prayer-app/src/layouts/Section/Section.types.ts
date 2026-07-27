import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface SectionProps
    extends HTMLAttributes<HTMLElement> {

    heading?: ReactNode;

    description?: ReactNode;

    actions?: ReactNode;

    children: ReactNode;

    bordered?: boolean;

    padded?: boolean;
}