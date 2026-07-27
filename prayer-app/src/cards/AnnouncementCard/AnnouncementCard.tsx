// New file generated from AnnouncementCard.tsx
import {
    Badge,
    Button,
    Card,
    Flex,
    Image,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { announcementCardVariants } from "./AnnouncementCard.styles";

import type { AnnouncementCardProps } from "./AnnouncementCard.types";

const priorityColor = {
    low: "neutral",
    normal: "info",
    high: "warning",
    urgent: "danger",
} as const;

const AnnouncementCard = ({
    heading,
    description,
    image,
    author,
    publishedAt,
    priority = "normal",
    actions,
    className,
    ...props
}: AnnouncementCardProps) => {
    return (
        <Card
            className={cn(
                announcementCardVariants({
                    priority,
                }),
                className
            )}
            {...props}
        >
            {image && (
                <Image
                    src={image}
                    alt=""
                    className="h-48 w-full"
                    fit="cover"
                />
            )}

            <Stack
                gap="md"
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

                        {publishedAt && (
                            <Typography
                                variant="caption"
                                className="text-muted-foreground"
                            >
                                {publishedAt}
                            </Typography>
                        )}

                    </Stack>

                    <Badge
                        variant="soft"
                        color={
                            priorityColor[
                                priority
                            ]
                        }
                    >
                        {priority}
                    </Badge>
                </Flex>

                {description && (
                    <Typography
                        variant="body"
                        className="text-muted-foreground"
                    >
                        {description}
                    </Typography>
                )}

                {(author || actions) && (
                    <Flex
                        justify="between"
                        align="center"
                    >
                        {author && (
                            <Typography
                                variant="body-sm"
                            >
                                {author}
                            </Typography>
                        )}

                        {actions}
                    </Flex>
                )}
            </Stack>
        </Card>
    );
};

export default AnnouncementCard;