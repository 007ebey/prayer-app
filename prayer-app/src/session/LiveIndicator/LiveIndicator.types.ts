// New file generated from LiveIndicator.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type LiveIndicatorState =
    | "live"
    | "connecting"
    | "offline";

export interface LiveIndicatorProps
    extends HTMLAttributes<HTMLSpanElement> {

    state?: LiveIndicatorState;

    label?: ReactNode;

}