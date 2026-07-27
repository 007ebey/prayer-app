// New file generated from InviteDialog.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface InviteDialogProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title" | "onSubmit"
    > {

    open: boolean;

    onClose: () => void;

    onInvite: (
        values: {
            emails: string;
            message: string;
        }
    ) => void;

    heading?: ReactNode;

    description?: ReactNode;

    defaultMessage?: string;

    loading?: boolean;
}