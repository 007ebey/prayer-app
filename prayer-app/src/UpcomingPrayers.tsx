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

interface PrayerSession {
    id: string;
    prayerGroupID: string;
    title: string;
    date: string;
    time: string;
    duration: number;
    prayerPointIDs?: string[];
}

interface UpcomingPrayersProps {
    prayers: PrayerSession[];
}

const INITIAL_VISIBLE_COUNT = 3;

const formatDate = (date: string): string => {
    return new Date(date).toLocaleDateString("en-IN", {
        weekday: "long",
    });
};

const UpcomingPrayers = ({
    prayers,
}: UpcomingPrayersProps) => {
    const [showAll, setShowAll] = useState(false);

    const visiblePrayers = showAll
        ? prayers
        : prayers.slice(0, INITIAL_VISIBLE_COUNT);

    const hiddenCount = Math.max(
        prayers.length - INITIAL_VISIBLE_COUNT,
        0
    );

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

                {prayers.length > INITIAL_VISIBLE_COUNT && (
                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={() =>
                            setShowAll((previous) => !previous)
                        }
                    >
                        {showAll
                            ? "Show less"
                            : `View all (${prayers.length})`}

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
                                    Join us for a time of prayer
                                    and fellowship.
                                </Typography>

                            </Stack>

                            {/* Schedule */}

                            <div className="flex flex-wrap items-center gap-4 text-muted-foreground">

                                <div className="flex items-center gap-2">

                                    <CalendarDays className="h-4 w-4" />

                                    <span className="text-sm">
                                        {formatDate(prayer.date)}
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
                                {prayer.duration} min
                            </Typography>

                        </Stack>

                    </Surface>

                ))}

            </div>

            {/* Empty state */}

            {prayers.length === 0 && (
                <Surface
                    elevation="raised"
                    padding="lg"
                    radius="lg"
                >
                    <Typography
                        variant="body"
                        color="muted"
                    >
                        No upcoming prayer sessions.
                    </Typography>
                </Surface>
            )}

            {/* Expanded state footer */}

            {showAll && hiddenCount > 0 && (
                <div className="flex justify-center">

                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={() => setShowAll(false)}
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