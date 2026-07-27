// New file generated from PrayerCard.tsx
import {
    Heart,
    MessageCircle,
    User,
} from "lucide-react";

import {
    Badge,
    Button,
    Card,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { prayerCardVariants } from "./PrayerCard.styles";

import type { PrayerCardProps } from "./PrayerCard.types";

const priorityColors = {
    low: "neutral",
    normal: "primary",
    high: "warning",
    urgent: "danger",
} as const;

const statusColors = {
    active: "primary",
    answered: "success",
    archived: "neutral",
} as const;

const PrayerCard = ({
    heading,
    description,
    author,
    createdAt,
    category,
    priority = "normal",
    status = "active",
    prayerCount = 0,
    commentCount = 0,
    isPraying = false,
    actions,
    className,
    ...props
}: PrayerCardProps) => {

    return (
        <Card
            className={cn(
                prayerCardVariants({
                    status,
                }),
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
                            variant="h4"
                        >
                            {heading}
                        </Typography>

                        {description && (
                            <Typography
                                variant="body"
                                className="text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}

                    </Stack>

                    <Stack gap="xs">

                        <Badge
                            variant="soft"
                            color={statusColors[status]}
                            size="sm"
                        >
                            {status}
                        </Badge>

                        <Badge
                            variant="outline"
                            color={priorityColors[priority]}
                            size="sm"
                        >
                            {priority}
                        </Badge>

                    </Stack>
                </Flex>

                {(author || createdAt || category) && (
                    <Flex
                        gap="md"
                        wrap="wrap"
                    >
                        {author && (
                            <Typography
                                variant="caption"
                            >
                                <User
                                    size={14}
                                    className="mr-1 inline"
                                />
                                {author}
                            </Typography>
                        )}

                        {createdAt && (
                            <Typography
                                variant="caption"
                            >
                                {createdAt}
                            </Typography>
                        )}

                        {category && (
                            <Badge
                                variant="soft"
                                color="info"
                                size="sm"
                            >
                                {category}
                            </Badge>
                        )}
                    </Flex>
                )}

                <Flex
                    justify="between"
                    align="center"
                >
                    <Flex gap="lg">

                        <Typography
                            variant="body-sm"
                        >
                            🙏 {prayerCount}
                        </Typography>

                        <Typography
                            variant="body-sm"
                        >
                            <MessageCircle
                                size={16}
                                className="mr-1 inline"
                            />
                            {commentCount}
                        </Typography>

                    </Flex>

                    {actions ?? (
                        <Button
                            variant={
                                isPraying
                                    ? "primary"
                                    : "outline"
                            }
                            size="sm"
                        >
                            <Heart
                                size={16}
                            />
                            {isPraying
                                ? "Praying"
                                : "Pray"}
                        </Button>
                    )}
                </Flex>
            </Stack>
        </Card>
    );
};

export default PrayerCard;