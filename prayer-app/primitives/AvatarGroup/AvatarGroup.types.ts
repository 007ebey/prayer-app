import type {
    HTMLAttributes,
    ReactElement,
} from "react";

export type AvatarGroupOverlap =
    | "none"
    | "sm"
    | "md"
    | "lg";

export interface AvatarGroupProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactElement | ReactElement[];

    max?: number;

    overlap?: AvatarGroupOverlap;

    reverse?: boolean;

    bordered?: boolean;
}