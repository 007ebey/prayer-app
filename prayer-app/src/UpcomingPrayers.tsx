import { useState } from "react";

import {
    CalendarDays,
    Clock,
    ArrowRight,
    ChevronUp,
    Users,
} from "lucide-react";

import {
    Stack,
    Surface,
    Typography,
    Badge,
    Button,
} from "./../primitives";

const upcomingPrayers = [
    {
        id: 1,
        title: "Morning Prayer",
        description: "Start the day together in prayer.",
        day: "Tomorrow",
        time: "6:00 AM",
        duration: "30 min",
        interested: 18,
    },
    {
        id: 2,
        title: "Lunch Prayer",
        description: "Take a moment in the middle of the day to pray.",
        day: "Tomorrow",
        time: "1:00 PM",
        duration: "20 min",
        interested: 12,
    },
    {
        id: 3,
        title: "Evening Intercession",
        description: "Come together and intercede for others.",
        day: "Tomorrow",
        time: "7:00 PM",
        duration: "45 min",
        interested: 34,
    },
    {
        id: 4,
        title: "Prayer for Families",
        description: "Pray for marriages, children, parents, and homes.",
        day: "Friday",
        time: "8:00 PM",
        duration: "45 min",
        interested: 27,
    },
    {
        id: 5,
        title: "Healing Prayer",
        description: "A focused time of prayer for healing and restoration.",
        day: "Saturday",
        time: "6:30 PM",
        duration: "60 min",
        interested: 41,
    },
    {
        id: 6,
        title: "Prayer for the Nation",
        description: "Intercede together for leaders, communities, and the nation.",
        day: "Sunday",
        time: "7:30 PM",
        duration: "45 min",
        interested: 52,
    },
    {
        id: 7,
        title: "Students Prayer",
        description: "Pray for students, schools, universities, and their future.",
        day: "Monday",
        time: "7:00 AM",
        duration: "30 min",
        interested: 16,
    },
    {
        id: 8,
        title: "Workplace Prayer",
        description: "Pray for wisdom, opportunities, workplaces, and careers.",
        day: "Tuesday",
        time: "7:00 AM",
        duration: "30 min",
        interested: 23,
    },
    {
        id: 9,
        title: "Night Prayer",
        description: "End the day in worship, thanksgiving, and prayer.",
        day: "Wednesday",
        time: "10:00 PM",
        duration: "30 min",
        interested: 29,
    },
];

const INITIAL_VISIBLE_COUNT = 3;

const UpcomingPrayers = () => {
    const [showAll, setShowAll] = useState(false);

    const visiblePrayers = showAll
        ? upcomingPrayers
        : upcomingPrayers.slice(
              0,
              INITIAL_VISIBLE_COUNT
          );

    const hiddenCount =
        upcomingPrayers.length -
        INITIAL_VISIBLE_COUNT;

    return (
        <Stack gap="lg">

            {/* Section Header */}

            <div className="flex items-end justify-between gap-4">

                <Stack gap="xs">

                    <Typography variant="h2">
                        Upcoming Prayers
                    </Typography>

                    <Typography
                        variant="body"
                        color="muted"
                    >
                        Prayer sessions scheduled next.
                    </Typography>

                </Stack>

                {upcomingPrayers.length >
                    INITIAL_VISIBLE_COUNT && (

                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={() =>
                            setShowAll(
                                (previous) => !previous
                            )
                        }
                    >
                        {showAll
                            ? "Show less"
                            : `View all (${upcomingPrayers.length})`}

                        {showAll ? (
                            <ChevronUp className="h-4 w-4" />
                        ) : (
                            <ArrowRight className="h-4 w-4" />
                        )}

                    </Button>

                )}

            </div>


            {/* Sessions */}

            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">

                {visiblePrayers.map((prayer) => (

                    <Surface
                        key={prayer.id}
                        elevation="raised"
                        padding="lg"
                        radius="lg"
                    >

                        <Stack gap="md">

                            {/* Status */}

                            <div className="flex items-center justify-between gap-4">

                                <Badge>
                                    UPCOMING
                                </Badge>

                                <div className="flex items-center gap-1.5 text-muted-foreground">

                                    <Users className="h-4 w-4" />

                                    <span className="text-sm">
                                        {prayer.interested}
                                    </span>

                                </div>

                            </div>


                            {/* Title */}

                            <Stack gap="xs">

                                <Typography variant="h3">
                                    {prayer.title}
                                </Typography>

                                <Typography
                                    variant="body-sm"
                                    color="muted"
                                >
                                    {prayer.description}
                                </Typography>

                            </Stack>


                            {/* Schedule */}

                            <div className="flex flex-wrap items-center gap-4 text-muted-foreground">

                                <div className="flex items-center gap-2">

                                    <CalendarDays className="h-4 w-4" />

                                    <span className="text-sm">
                                        {prayer.day}
                                    </span>

                                </div>

                                <div className="flex items-center gap-2">

                                    <Clock className="h-4 w-4" />

                                    <span className="text-sm">
                                        {prayer.time}
                                    </span>

                                </div>

                            </div>


                            {/* Duration */}

                            <Typography
                                variant="body-sm"
                                color="muted"
                            >
                                {prayer.duration}
                            </Typography>

                        </Stack>

                    </Surface>

                ))}

            </div>


            {/* Expanded state footer */}

            {showAll && hiddenCount > 0 && (

                <div className="flex justify-center">

                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={() =>
                            setShowAll(false)
                        }
                    >
                        Show less

                        <ChevronUp className="h-4 w-4" />
                    </Button>

                </div>

            )}

        </Stack>
    );
};

export default UpcomingPrayers;