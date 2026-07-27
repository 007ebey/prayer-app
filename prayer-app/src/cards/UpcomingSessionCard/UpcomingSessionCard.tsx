// New file generated from UpcomingSessionCard.tsx
import {
    Bell,
    Calendar,
    Clock,
    User,
    Users,
} from "lucide-react";

import {
    Badge,
    Button,
    Card,
    Flex,
    Image,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { upcomingSessionCardVariants } from "./UpcomingSessionCard.styles";

import type {
    UpcomingSessionCardProps,
} from "./UpcomingSessionCard.types";

const UpcomingSessionCard = ({
    heading,
    description,
    image,
    host,
    startDate,
    duration,
    countdown,
    participantCount,
    maxParticipants,
    reminderEnabled = false,
    tags,
    actions,
    className,
    ...props
}: UpcomingSessionCardProps) => {

    return (
        <Card
            className={cn(
                upcomingSessionCardVariants(),
                className
            )}
            {...props}
        >

            {image && (
                <Image
                    src={image}
                    alt=""
                    aspectRatio="video"
                    fit="cover"
                    radius="none"
                />
            )}

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
                                variant="body-sm"
                                className="text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}

                    </Stack>

                    {countdown && (
                        <Badge
                            variant="soft"
                            color="warning"
                        >
                            {countdown}
                        </Badge>
                    )}

                </Flex>

                <Stack gap="sm">

                    <Flex
                        gap="sm"
                        align="center"
                    >
                        <Calendar size={16} />

                        <Typography
                            variant="body-sm"
                        >
                            {startDate}
                        </Typography>
                    </Flex>

                    {duration && (
                        <Flex
                            gap="sm"
                            align="center"
                        >
                            <Clock size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                {duration}
                            </Typography>
                        </Flex>
                    )}

                    {host && (
                        <Flex
                            gap="sm"
                            align="center"
                        >
                            <User size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                {host}
                            </Typography>
                        </Flex>
                    )}

                    {participantCount != null && (
                        <Flex
                            gap="sm"
                            align="center"
                        >
                            <Users size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                {participantCount}

                                {maxParticipants != null &&
                                    ` / ${maxParticipants}`}

                                {" registered"}
                            </Typography>
                        </Flex>
                    )}

                </Stack>

                {tags}

                <Flex
                    justify="between"
                    align="center"
                >

                    <Button
                        variant={
                            reminderEnabled
                                ? "primary"
                                : "outline"
                        }
                        size="sm"
                    >
                        <Bell
                            size={16}
                        />

                        {reminderEnabled
                            ? "Reminder Set"
                            : "Remind Me"}
                    </Button>

                    {actions ?? (
                        <Button
                            size="sm"
                        >
                            View Session
                        </Button>
                    )}

                </Flex>

            </Stack>

        </Card>
    );
};

export default UpcomingSessionCard;