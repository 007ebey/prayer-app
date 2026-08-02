import {
    ArrowLeft,
    Church,
    Users,
    Heart,
} from "lucide-react";

import {
    Stack,
    Surface,
    Typography,
    Badge,
    Avatar,
    AvatarGroup,
    Button,
} from "./../../../primitives";
import { useState } from "react";
import PrayerParticipantList from "../../PrayerParticipantList";

interface PrayerWallProps {
    onLeave: () => void;
}



const prayerPoints = [
    {
        id: 1,
        title: "Pray for families",
        description:
            "Pray for unity, restoration, wisdom, and peace in families.",
    },
    {
        id: 2,
        title: "Pray for those who need healing",
        description:
            "Remember those facing sickness, pain, and difficult circumstances.",
    },
    {
        id: 3,
        title: "Pray for young people",
        description:
            "Pray for students and young adults as they make important decisions.",
    },
    {
        id: 4,
        title: "Pray for our communities",
        description:
            "Pray for peace, compassion, and opportunities to serve others.",
    },
];

const PrayerWall = ({
    onLeave,
}: PrayerWallProps) => {

    const [showParticipants, setShowParticipants] =
        useState(false);

    const [participants, setParticipants] =
        useState([
            {
                id: "1",
                name: "Pradeep Ebey",
                greeted: false,
            },
            {
                id: "2",
                name: "John Samuel",
                greeted: false,
            },
            {
                id: "3",
                name: "Anna Mary",
                greeted: false,
            },
            {
                id: "4",
                name: "Robert K",
                greeted: false,
            },
        ]);

    const handleSayHi = (
        participantId: string
    ) => {

        setParticipants((current) =>
            current.map((participant) =>
                participant.id === participantId
                    ? {
                        ...participant,
                        greeted: true,
                    }
                    : participant
            )
        );

    };


    return (
        <Stack gap="2xl">

            {/* Top bar */}

            <div className="flex items-center justify-between gap-4">

                <Button
                    variant="ghost"
                    size="sm"
                    onClick={onLeave}
                >
                    <ArrowLeft className="h-4 w-4" />

                    Leave Prayer
                </Button>

                <Badge color="success">
                    LIVE
                </Badge>

            </div>


            {/* Session Header */}

            <Surface
                elevation="raised"
                padding="xl"
                radius="xl"
            >
                <Stack gap="lg">

                    <div className="flex items-start justify-between gap-6">

                        <Stack gap="sm">

                            <div className="flex items-center gap-2 text-primary">

                                <Church className="h-5 w-5" />

                                <span className="text-sm font-medium">
                                    Prayer Wall
                                </span>

                            </div>

                            <Typography variant="display">
                                Evening Prayer Meeting
                            </Typography>

                            <Typography
                                variant="subtitle"
                                color="muted"
                            >
                                Pray together with believers
                                joining from around the world.
                            </Typography>

                        </Stack>

                    </div>

                    {/* Participants */}

                    <PrayerParticipantList
                        participants={participants}
                        onSayHi={handleSayHi}
                    />

                </Stack>
            </Surface>


            {/* Prayer points */}

            <Stack gap="lg">

                <div>

                    <Typography variant="h2">
                        Prayer Points
                    </Typography>

                    <Typography
                        variant="body"
                        color="muted"
                    >
                        Take time with each prayer point.
                    </Typography>

                </div>


                <div className="grid gap-4">

                    {prayerPoints.map(
                        (prayer, index) => (

                            <Surface
                                key={prayer.id}
                                elevation="raised"
                                padding="lg"
                                radius="lg"
                            >

                                <div className="flex gap-4">

                                    {/* Number */}

                                    <div className="
                                        flex h-10 w-10
                                        shrink-0
                                        items-center
                                        justify-center
                                        rounded-full
                                        bg-primary
                                        text-primary-foreground
                                        font-semibold
                                    ">
                                        {index + 1}
                                    </div>


                                    {/* Content */}

                                    <Stack gap="sm">

                                        <Typography variant="h3">
                                            {prayer.title}
                                        </Typography>

                                        <Typography
                                            variant="body"
                                            color="muted"
                                        >
                                            {prayer.description}
                                        </Typography>

                                    </Stack>

                                </div>

                            </Surface>

                        )
                    )}

                </div>

            </Stack>


            {/* Bottom prayer action */}

            <Surface
                elevation="raised"
                padding="lg"
                radius="lg"
            >

                <div className="flex flex-wrap items-center justify-between gap-4">

                    <div className="flex items-center gap-3">

                        <Heart className="h-5 w-5 text-primary" />

                        <Stack gap="xs">

                            <Typography variant="h3">
                                Continue in prayer
                            </Typography>

                            <Typography
                                variant="body-sm"
                                color="muted"
                            >
                                Stay connected with everyone
                                praying right now.
                            </Typography>

                        </Stack>

                    </div>

                    <Button
                        variant="outline"
                        onClick={onLeave}
                    >
                        Leave Prayer
                    </Button>

                </div>

            </Surface>

        </Stack>
    );
};

export default PrayerWall;