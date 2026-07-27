// New file generated from ActivityCard.tsx
import {
    Card,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { activityCardVariants } from "./ActivityCard.styles";

import type { ActivityCardProps } from "./ActivityCard.types";

const ActivityCard = ({
    icon,
    title,
    description,
    timestamp,
    status,
    actions,
    className,
    ...props
}: ActivityCardProps) => {
    return (
        <Card
            className={cn(
                activityCardVariants({
                    status,
                }),
                className
            )}
            {...props}
        >
            <Flex
                align="start"
                justify="between"
                gap="md"
            >
                <Flex
                    align="start"
                    gap="md"
                    className="flex-1"
                >
                    {icon && (
                        <div className="flex h-10 w-10 items-center justify-center rounded-full bg-muted">
                            {icon}
                        </div>
                    )}

                    <Stack
                        gap="xs"
                        className="flex-1"
                    >
                        <Typography
                            variant="body"
                            weight="semibold"
                        >
                            {title}
                        </Typography>

                        {description && (
                            <Typography
                                variant="body-sm"
                                className="text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}

                        {timestamp && (
                            <Typography
                                variant="caption"
                                className="text-muted-foreground"
                            >
                                {timestamp}
                            </Typography>
                        )}
                    </Stack>
                </Flex>

                {actions}
            </Flex>
        </Card>
    );
};

export default ActivityCard;