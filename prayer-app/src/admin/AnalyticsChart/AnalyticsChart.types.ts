// New file generated from AnalyticsChart.tsx
import { ReactNode } from "react";

export type AnalyticsChartType =
    | "line"
    | "bar"
    | "area"
    | "pie";

export interface AnalyticsChartData {
    label: string;

    value: number;
}

export interface AnalyticsChartProps {
    heading?: ReactNode;

    description?: ReactNode;

    data: AnalyticsChartData[];

    type?: AnalyticsChartType;

    actions?: ReactNode;

    loading?: boolean;

    emptyState?: ReactNode;
}