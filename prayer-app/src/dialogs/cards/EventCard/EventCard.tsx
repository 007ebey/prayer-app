// New file generated from EventCard.tsx
import {
    Calendar,
    Clock,
    MapPin,
    Users,
} from "lucide-react";

import {
    Badge,
    Card,
    Flex,
    Image,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { eventCardVariants } from "./EventCard.styles";

import type { EventCardProps } from "./EventCard.types";

const badgeColors = {
    upcoming: "primary",
    live: "success",
    completed: "neutral",
    cancelled: "danger",
} as const;

const EventCard = ({
    image,
    heading,
    description,
    startDate,
    endDate,
    location,
    organizer,
    attendees,
    status = "upcoming",
    actions,
    className,
    ...props
}: EventCardProps) => {

    return (
        <Card
            className={cn(
                eventCardVariants({
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

                    <Badge
                        variant="soft"
                        color={badgeColors[status]}
                    >
                        {status}
                    </Badge>

                </Flex>

                <Stack gap="sm">

                    <Flex gap="sm" align="center">

                        <Calendar size={16} />

                        <Typography
                            variant="body-sm"
                        >
                            {startDate}

                            {endDate &&
                                <> - {endDate}</>}
                        </Typography>

                    </Flex>

                    {location && (

                        <Flex gap="sm" align="center">

                            <MapPin size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                {location}
                            </Typography>

                        </Flex>

                    )}

                    {organizer && (

                        <Flex gap="sm" align="center">

                            <Clock size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                Hosted by {organizer}
                            </Typography>

                        </Flex>

                    )}

                    {attendees !== undefined && (

                        <Flex gap="sm" align="center">

                            <Users size={16} />

                            <Typography
                                variant="body-sm"
                            >
                                {attendees} attending
                            </Typography>

                        </Flex>

                    )}

                </Stack>

                {actions}

            </Stack>

        </Card>
    );
};

export default EventCard;