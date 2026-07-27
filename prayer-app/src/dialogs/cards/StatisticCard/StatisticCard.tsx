// New file generated from StatisticCard.tsx
import {
    ArrowDownRight,
    ArrowUpRight,
    Minus,
} from "lucide-react";

import {
    Card,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { statisticsCardVariants } from "./StatisticCard.styles";

import type {
    StatisticsCardProps,
} from "./StatisticCard.types";

const trendIcons = {
    up: ArrowUpRight,
    down: ArrowDownRight,
    neutral: Minus,
};

const trendColors = {
    up: "text-success",
    down: "text-destructive",
    neutral: "text-muted-foreground",
};

const StatisticsCard = ({
    heading,
    value,
    icon,
    subtitle,
    trend,
    trendValue,
    footer,
    actions,
    className,
    ...props
}: StatisticsCardProps) => {

    const TrendIcon =
        trend
            ? trendIcons[trend]
            : null;

    return (
        <Card
            className={cn(
                statisticsCardVariants(),
                className
            )}
            {...props}
        >
            <Stack
                gap="lg"
                className="p-6"
            >

                <Flex
                    justify="between"
                    align="start"
                >

                    <Stack gap="xs">

                        <Typography
                            variant="body-sm"
                            className="text-muted-foreground"
                        >
                            {heading}
                        </Typography>

                        <Typography
                            variant="display-sm"
                            weight="bold"
                        >
                            {value}
                        </Typography>

                    </Stack>

                    {icon && (
                        <div className="flex h-12 w-12 items-center justify-center rounded-lg bg-muted">
                            {icon}
                        </div>
                    )}

                </Flex>

                {(subtitle ||
                    trend) && (
                    <Flex
                        justify="between"
                        align="center"
                    >

                        {subtitle && (
                            <Typography
                                variant="caption"
                                className="text-muted-foreground"
                            >
                                {subtitle}
                            </Typography>
                        )}

                        {trend &&
                            TrendIcon && (
                                <Flex
                                    gap="xs"
                                    align="center"
                                    className={
                                        trendColors[
                                            trend
                                        ]
                                    }
                                >
                                    <TrendIcon
                                        size={16}
                                    />

                                    {trendValue && (
                                        <Typography
                                            variant="caption"
                                        >
                                            {
                                                trendValue
                                            }
                                        </Typography>
                                    )}
                                </Flex>
                            )}

                    </Flex>
                )}

                {(footer ||
                    actions) && (
                    <Flex
                        justify="between"
                        align="center"
                    >
                        {footer}

                        {actions}
                    </Flex>
                )}

            </Stack>
        </Card>
    );
};

export default StatisticsCard;