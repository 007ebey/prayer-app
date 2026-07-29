import type { ReactNode } from "react";

export type PopoverSide =
    | "top"
    | "right"
    | "bottom"
    | "left";

export interface PopoverProps {

    trigger: ReactNode;

    children: ReactNode;

    open?: boolean;

    defaultOpen?: boolean;

    onOpenChange?: (
        open: boolean
    ) => void;

    side?: PopoverSide;

    align?: "start" | "center" | "end";

}