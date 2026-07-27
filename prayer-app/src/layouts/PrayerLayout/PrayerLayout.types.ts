import { ReactNode } from "react";

export interface PrayerLayoutProps {
    header?: ReactNode;

    prayerList?: ReactNode;

    prayerWall?: ReactNode;

    participants?: ReactNode;

    footer?: ReactNode;

    children?: ReactNode;
}