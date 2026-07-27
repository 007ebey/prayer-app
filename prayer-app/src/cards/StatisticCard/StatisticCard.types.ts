// New file generated from StatisticCard.tsx
import {
    HTMLAttributes,
    ReactNode,
} from "react";

export type StatisticsTrend =
    | "up"
    | "down"
    | "neutral";

export interface StatisticsCardProps
    extends HTMLAttributes<HTMLDivElement> {

    heading: ReactNode;

    value: ReactNode;

    icon?: ReactNode;

    subtitle?: ReactNode;

    trend?: StatisticsTrend;

    trendValue?: ReactNode;

    footer?: ReactNode;

    actions?: ReactNode;
}