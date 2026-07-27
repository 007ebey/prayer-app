// New file generated from SessionEditor.tsx
import {
    FormEvent,
    HTMLAttributes,
    ReactNode,
} from "react";

export interface SessionEditorValues {
    title: string;

    description: string;

    hostId: string;

    backgroundImage?: string;

    startDate: string;

    endDate: string;

    isPublic: boolean;

    isActive: boolean;
}

export interface SessionEditorProps {

    className?: string;

    heading?: ReactNode;

    description?: ReactNode;

    values: SessionEditorValues;

    loading?: boolean;

    onValuesChange?: (
        values: SessionEditorValues
    ) => void;

    onSubmit?: (
        event: FormEvent<HTMLFormElement>
    ) => void;

    onCancel?: () => void;
}