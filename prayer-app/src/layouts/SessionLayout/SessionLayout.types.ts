import { ReactNode } from "react";

export interface SessionLayoutProps {
    children: ReactNode;

    header?: ReactNode;

    prayerPanel?: ReactNode;

    participantsPanel?: ReactNode;

    chatPanel?: ReactNode;

    footer?: ReactNode;
}