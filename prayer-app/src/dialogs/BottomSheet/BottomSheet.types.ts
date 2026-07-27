// New file generated from BottomSheet.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type BottomSheetSize =
    | "sm"
    | "md"
    | "lg"
    | "full";

export interface BottomSheetProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    open: boolean;

    onClose: () => void;

    heading?: ReactNode;

    description?: ReactNode;

    footer?: ReactNode;

    children?: ReactNode;

    closeOnBackdrop?: boolean;

    closeOnEscape?: boolean;

    showHandle?: boolean;

    size?: BottomSheetSize;
}