export interface PrayerSession {
    id: string;
    title: string;
    date: string;
    time: string;
}

export interface User {
    id: string;
    name: string;
    email: string;

    prayerGroupIds: string[];
}

export interface PrayerGroup {
    id: string;
    name: string;
    description: string;

    sessionIds: string[];

    isDefault: boolean;
}

export interface PrayerGroupsProps {
    prayerGroups: PrayerGroup[];

    sessions: PrayerSession[];

    users: User[];

    newGroupName: string;

    newGroupDescription: string;

    onNameChange: (
        value: string
    ) => void;

    onDescriptionChange: (
        value: string
    ) => void;

    onCreate: () => void;

    onAssignSession: (
        prayerGroupId: string,
        sessionId: string
    ) => void;

    onRemoveSession: (
        prayerGroupId: string,
        sessionId: string
    ) => void;

    onAssignUser: (
        prayerGroupId: string,
        userId: string
    ) => void;

    onRemoveUser: (
        prayerGroupId: string,
        userId: string
    ) => void;
}