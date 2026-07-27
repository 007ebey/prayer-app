import { HTMLAttributes } from "react";

export type AvatarSize =
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl"
    | "2xl";

export type AvatarShape =
    | "circle"
    | "rounded"
    | "square";

export type AvatarStatus =
    | "online"
    | "offline"
    | "busy"
    | "away";

export interface AvatarProps
    extends HTMLAttributes<HTMLDivElement> {

    src?: string;

    alt?: string;

    name?: string;

    size?: AvatarSize;

    shape?: AvatarShape;

    status?: AvatarStatus;

    bordered?: boolean;
}