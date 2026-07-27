// New file generated from SessionCard.tsx
import {
    Calendar,
    Clock,
    Radio,
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

import { sessionCardVariants } from "./SessionCard.styles";

import type { SessionCardProps } from "./SessionCard.types";

const statusColors = {
    live: "success",
    upcoming: "primary",
    completed: "neutral",
    cancelled: "danger",
} as const;

const SessionCard = ({
    heading,
    description,
    image,
    host,
    startTime,
    endTime,
    duration,
    participantCount,
    maxParticipants,
    tags,
    status = "upcoming",
    actions,
    className,
    ...props
}: SessionCardProps) => {

    return (
        <Card
            className={cn(
                sessionCardVariants({
                    status,
                }),
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

                        <Typography variant="h4">
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

                    <Badge
                        variant="soft"
                        color={statusColors[status]}
                    >
                        {status === "live" && (
                            <Radio
                                size={10}
                                className="mr-1 fill-current"
                            />
                        )}

                        {status}
                    </Badge>

                </Flex>

                <Stack gap="sm">

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

                    {startTime && (
                        <Flex
                            gap="sm"
                            align="center"
                        >
                            <Calendar size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                {startTime}

                                {endTime &&
                                    <> - {endTime}</>}
                            </Typography>
                        </Flex>
                    )}

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

                    {participantCount !==
                        undefined && (
                        <Flex
                            gap="sm"
                            align="center"
                        >
                            <Users size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                {participantCount}

                                {maxParticipants &&
                                    ` / ${maxParticipants}`}

                                {" participants"}
                            </Typography>
                        </Flex>
                    )}

                </Stack>

                {tags}

                {actions ?? (
                    <Button
                        fullWidth
                        variant={
                            status === "live"
                                ? "primary"
                                : "outline"
                        }
                    >
                        {status === "live"
                            ? "Join Session"
                            : status ===
                                "upcoming"
                              ? "View Details"
                              : "View Summary"}
                    </Button>
                )}

            </Stack>

        </Card>
    );
};

export default SessionCard;