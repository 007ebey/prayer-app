export interface PrayerPoint {
    id: string;

    title: string;

    description: string;
}

export interface PrayerSession {
    id: string;

    title: string;

    date: string;

    time: string;

    prayerPointIds: string[];
}

export interface PrayerPointsProps {
    prayerPoints: PrayerPoint[];

    sessions: PrayerSession[];

    title: string;

    description: string;

    onTitleChange: (
        value: string
    ) => void;

    onDescriptionChange: (
        value: string
    ) => void;

    onCreate: () => void;
}