// New file generated from PrayerWall.tsx
import { ReactNode } from "react";

export interface PrayerPoint {
    ID: string;
    GroupID: string;
    Title: string;
    Content: string;
    Status: string;
    CreatedAt: string;
    UpdatedAt: string;
}

export interface PrayerSession {
    id: string;
    prayerGroupID: string;
    title: string;
    description: string;
    date: string;
    time: string;
    duration: number;
    prayerPointIDs: string[];
    prayerPoints: PrayerPoint[];
}

export interface PrayerWallProps {
    session: PrayerSession;
    onLeave: () => void;
}