import {
    Plus,
} from "lucide-react";

import {
    Badge,
    Button,
    Input,
    Stack,
    Surface,
    Typography,
} from "../../../primitives";

import type {
    PrayerPointsProps,
} from "./PrayerPoints.types";

const PrayerPoints = ({
    prayerPoints,
    sessions,

    title,
    description,

    onTitleChange,
    onDescriptionChange,
    onCreate,
}: PrayerPointsProps) => {

    return (

        <Surface
            elevation="raised"
            padding="xl"
            radius="xl"
        >

            <Stack gap="lg">

                {/* Header */}

                <div>

                    <Typography variant="h2">
                        Prayer Points
                    </Typography>

                    <Typography
                        variant="body-sm"
                        color="muted"
                    >
                        Create reusable prayer points
                        and associate them with prayer
                        sessions.
                    </Typography>

                </div>


                {/* Create */}

                <Stack gap="md">

                    <Typography variant="h3">
                        Create Prayer Point
                    </Typography>

                    <div className="grid gap-4 md:grid-cols-2">

                        <Input
                            value={title}
                            placeholder="Prayer point title"
                            onChange={(event) =>
                                onTitleChange(
                                    event.target.value
                                )
                            }
                        />

                        <Input
                            value={description}
                            placeholder="Prayer point description"
                            onChange={(event) =>
                                onDescriptionChange(
                                    event.target.value
                                )
                            }
                        />

                    </div>

                    <div>

                        <Button
                            leftIcon={
                                <Plus className="h-4 w-4" />
                            }
                            onClick={onCreate}
                        >
                            Add Prayer Point
                        </Button>

                    </div>

                </Stack>


                {/* Library */}

                <div className="border-t border-border pt-6">

                    <Stack gap="md">

                        <div>

                            <Typography variant="h3">
                                Prayer Point Library
                            </Typography>

                            <Typography
                                variant="body-sm"
                                color="muted"
                            >
                                Reusable prayer points
                                available for prayer sessions.
                            </Typography>

                        </div>


                        {prayerPoints.length === 0 ? (

                            <Surface
                                bordered
                                padding="lg"
                                radius="lg"
                            >

                                <Typography
                                    variant="body-sm"
                                    color="muted"
                                >
                                    No prayer points have
                                    been created yet.
                                </Typography>

                            </Surface>

                        ) : (

                            <div className="grid gap-4">

                                {prayerPoints.map(
                                    (prayerPoint) => {

                                        const associatedSessions =
                                            sessions.filter(
                                                (session) =>
                                                    session
                                                        .prayerPointIds
                                                        .includes(
                                                            prayerPoint.id
                                                        )
                                            );

                                        return (

                                            <Surface
                                                key={
                                                    prayerPoint.id
                                                }
                                                bordered
                                                padding="lg"
                                                radius="lg"
                                            >

                                                <Stack gap="md">

                                                    <div className="flex flex-wrap items-start justify-between gap-4">

                                                        <Stack gap="xs">

                                                            <Typography variant="title">
                                                                {
                                                                    prayerPoint.title
                                                                }
                                                            </Typography>

                                                            <Typography
                                                                variant="body-sm"
                                                                color="muted"
                                                            >
                                                                {
                                                                    prayerPoint.description
                                                                }
                                                            </Typography>

                                                        </Stack>

                                                        <Badge color="info">

                                                            {
                                                                associatedSessions.length
                                                            }{" "}

                                                            {
                                                                associatedSessions.length === 1
                                                                    ? "session"
                                                                    : "sessions"
                                                            }

                                                        </Badge>

                                                    </div>


                                                    {/* Associations */}

                                                    <div>

                                                        <Typography
                                                            variant="caption"
                                                            color="muted"
                                                        >
                                                            USED IN
                                                        </Typography>

                                                        {associatedSessions.length === 0 ? (

                                                            <Typography
                                                                variant="body-sm"
                                                                color="muted"
                                                                className="mt-2"
                                                            >
                                                                Not assigned
                                                                to any prayer
                                                                session.
                                                            </Typography>

                                                        ) : (

                                                            <div className="mt-2 flex flex-wrap gap-2">

                                                                {associatedSessions.map(
                                                                    (session) => (

                                                                        <Badge
                                                                            key={
                                                                                session.id
                                                                            }
                                                                            color="secondary"
                                                                        >
                                                                            {
                                                                                session.title
                                                                            }
                                                                        </Badge>

                                                                    )
                                                                )}

                                                            </div>

                                                        )}

                                                    </div>

                                                </Stack>

                                            </Surface>

                                        );

                                    }
                                )}

                            </div>

                        )}

                    </Stack>

                </div>

            </Stack>

        </Surface>

    );

};

export default PrayerPoints;