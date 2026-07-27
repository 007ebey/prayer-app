// New file generated from BulkActions.tsx
import { HTMLAttributes, ReactNode } from "react";

export interface BulkAction {
    id: string;

    label: ReactNode;

    icon?: ReactNode;

    disabled?: boolean;

    onClick?: () => void;
}

export interface BulkActionsProps
    extends HTMLAttributes<HTMLDivElement> {

    selectedCount: number;

    actions: BulkAction[];

    onClearSelection?: () => void;
}