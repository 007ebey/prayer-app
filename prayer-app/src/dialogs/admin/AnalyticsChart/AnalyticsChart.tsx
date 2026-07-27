// New file generated from AnalyticsChart.tsx
import {
    Card,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { analyticsChartVariants } from "./AnalyticsChart.styles";

import type { AnalyticsChartProps } from "./AnalyticsChart.types";

const AnalyticsChart = ({
    heading,
    description,
    actions,
    data,
    loading = false,
    emptyState,
    className,
}: AnalyticsChartProps & {
    className?: string;
}) => {
    return (
        <Card
            className={cn(
                analyticsChartVariants(),
                className
            )}
        >
            {(heading ||
                description ||
                actions) && (
                <Flex
                    justify="between"
                    align="start"
                    className="mb-6"
                >
                    <Stack gap="xs">
                        {heading && (
                            <Typography variant="h4">
                                {heading}
                            </Typography>
                        )}

                        {description && (
                            <Typography
                                variant="caption"
                                className="text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}
                    </Stack>

                    {actions}
                </Flex>
            )}

            {loading ? (
                <div className="flex h-72 items-center justify-center">
                    <Typography>
                        Loading...
                    </Typography>
                </div>
            ) : data.length === 0 ? (
                emptyState ?? (
                    <div className="flex h-72 items-center justify-center">
                        <Typography>
                            No data available
                        </Typography>
                    </div>
                )
            ) : (
                <div className="flex h-72 items-center justify-center rounded-md border border-dashed">
                    <Typography variant="caption">
                        Chart placeholder
                    </Typography>
                </div>
            )}
        </Card>
    );
};

export default AnalyticsChart;