// New file generated from SettingsDialog.tsx
import {
    HTMLAttributes,
} from "react";

export interface SettingsValues {

    notifications: boolean;

    prayerReminders: boolean;

    sessionSounds: boolean;

    autoJoin: boolean;

    darkMode: boolean;

}

export interface SettingsDialogProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "onChange"
    > {

    open: boolean;

    values: SettingsValues;

    loading?: boolean;

    onClose: () => void;

    onSave: (
        values: SettingsValues
    ) => void;

}