// New file generated from ShareDialog.tsx
// New file generated from ShareDialog.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type ShareOption =
    | "copy"
    | "email"
    | "whatsapp"
    | "telegram"
    | "facebook"
    | "x"
    | "linkedin";

export interface ShareDialogProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "onChange"
    > {

    open: boolean;

    onClose: () => void;

    url: string;

    heading?: ReactNode;

    description?: ReactNode;

    enabledOptions?: ShareOption[];

    onShare?: (
        option: ShareOption
    ) => void;
}