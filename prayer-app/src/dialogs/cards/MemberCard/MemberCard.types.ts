import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type MemberStatus =
    | "online"
    | "offline"
    | "away"
    | "busy";

export type MemberRole =
    | "member"
    | "host"
    | "admin"
    | "super-admin";

export interface MemberCardProps
    extends HTMLAttributes<HTMLDivElement> {

    name: ReactNode;

    avatar?: string;

    status?: MemberStatus;

    memberRole?: MemberRole;

    roleLabel?: ReactNode;

    location?: ReactNode;

    bio?: ReactNode;

    joinedDate?: ReactNode;

    badges?: ReactNode;

    actions?: ReactNode;
}