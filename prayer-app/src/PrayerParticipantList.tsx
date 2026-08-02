import { useState } from "react";

import {
    Users,
    Hand,
} from "lucide-react";

import {
    Avatar,
    AvatarGroup,
    Button,
    Stack,
    Typography,
} from "../primitives";

export interface PrayerParticipant {
    id: string;

    name: string;

    greeted: boolean;

    avatarUrl?: string;
}

export interface PrayerParticipantListProps {
    participants: PrayerParticipant[];

    onSayHi?: (
        participantId: string
    ) => void;

    total?: number;

    previewCount?: number;
}

const PrayerParticipantList = ({
    participants,
    onSayHi,
    total,
    previewCount = 4,
}: PrayerParticipantListProps) => {

    const [
        showParticipants,
        setShowParticipants,
    ] = useState(false);

    const participantCount =
        total ?? participants.length;

    const previewParticipants =
        participants.slice(
            0,
            previewCount
        );

    return (

        <Stack gap="md">

            {/* Participant summary */}

            <div className="flex flex-wrap items-center gap-3">

                <AvatarGroup>

                    {previewParticipants.map(
                        (participant) => (

                            <Avatar
                                key={
                                    participant.id
                                }
                                name={
                                    participant.name
                                }
                                src={
                                    participant.avatarUrl
                                }
                            />

                        )
                    )}

                </AvatarGroup>


                {/* Count */}

                <div className="flex items-center gap-2 text-muted-foreground">

                    <Users className="h-4 w-4" />

                    <Typography
                        variant="body-sm"
                        color="muted"
                    >
                        {participantCount}{" "}
                        {participantCount === 1
                            ? "person"
                            : "people"}{" "}
                        praying
                    </Typography>

                </div>


                {/* View all */}

                {participants.length > 0 && (

                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={() =>
                            setShowParticipants(
                                (current) =>
                                    !current
                            )
                        }
                    >
                        {showParticipants
                            ? "Hide"
                            : "View all"}
                    </Button>

                )}

            </div>


            {/* Expanded participant list */}

            {showParticipants && (

                <div className="grid gap-3 md:grid-cols-2">

                    {participants.map(
                        (participant) => (

                            <div
                                key={
                                    participant.id
                                }
                                className="
                                    flex
                                    items-center
                                    justify-between
                                    gap-4
                                    rounded-lg
                                    border
                                    border-border
                                    p-3
                                "
                            >

                                {/* Participant */}

                                <div className="flex min-w-0 items-center gap-3">

                                    <Avatar
                                        name={
                                            participant.name
                                        }
                                        src={
                                            participant.avatarUrl
                                        }
                                    />

                                    <div className="min-w-0">

                                        <Typography
                                            variant="body-sm"
                                            className="truncate"
                                        >
                                            {
                                                participant.name
                                            }
                                        </Typography>

                                        <Typography
                                            variant="caption"
                                            color="muted"
                                        >
                                            Praying now
                                        </Typography>

                                    </div>

                                </div>


                                {/* Say Hi */}

                                <Button
                                    variant={
                                        participant.greeted
                                            ? "ghost"
                                            : "outline"
                                    }
                                    size="sm"
                                    disabled={
                                        participant.greeted
                                    }
                                    onClick={() =>
                                        onSayHi?.(
                                            participant.id
                                        )
                                    }
                                >

                                    <Hand className="h-4 w-4" />

                                    {participant.greeted
                                        ? "Hi sent"
                                        : "Say Hi"}

                                </Button>

                            </div>

                        )
                    )}

                </div>

            )}

        </Stack>

    );

};

export default PrayerParticipantList;