// New file generated from WelcomeCard.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface WelcomeCardProps
    extends HTMLAttributes<HTMLDivElement> {

    greeting?: ReactNode;

    name: ReactNode;

    message?: ReactNode;

    backgroundImage?: string;

    illustration?: ReactNode;

    actions?: ReactNode;

    footer?: ReactNode;
}