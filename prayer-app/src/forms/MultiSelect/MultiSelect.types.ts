// New file generated from MultiSelect.tsx
import {
    ReactNode,
} from "react";

export interface MultiSelectOption {

    value: string;

    label: ReactNode;

    disabled?: boolean;

}

export interface MultiSelectProps {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    placeholder?: string;

    options: MultiSelectOption[];

    value?: string[];

    onChange?(
        values: string[],
    ): void;

    disabled?: boolean;

    searchable?: boolean;

    clearable?: boolean;

    maxSelected?: number;

    fullWidth?: boolean;

}