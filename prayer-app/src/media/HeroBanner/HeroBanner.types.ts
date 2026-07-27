import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type HeroAlignment =
    | "left"
    | "center"
    | "right";

export interface HeroBannerProps
    extends HTMLAttributes<HTMLElement> {

    backgroundImage: string;

    heading: ReactNode;

    description?: ReactNode;

    primaryAction?: ReactNode;

    secondaryAction?: ReactNode;

    badge?: ReactNode;

    overlay?: boolean;

    alignment?: HeroAlignment;

    fullHeight?: boolean;

    children?: ReactNode;
}