import {
    useMemo,
    useState,
} from "react";

import {
    CalendarDays,
    Clock,
    Plus,
    Search,
    Trash2,
    UserCheck,
    UserX,
    Users,
} from "lucide-react";

import {
    Badge,
    Button,
    Input,
    Stack,
    Surface,
    Typography,
} from "../../../primitives";

type SessionStatus =
    | "scheduled"
    | "cancelled";

interface PrayerSession {
    id: string;
    title: string;
    date: string;
    time: string;
    status: SessionStatus;
}

interface Participant {
    id: string;
    name: string;
    email: string;
    hasAccess: boolean;
}

const initialSessions: PrayerSession[] = [
    {
        id: "session-1",
        title: "Morning Prayer",
        date: "2026-08-03",
        time: "06:00",
        status: "scheduled",
    },
    {
        id: "session-2",
        title: "Evening Prayer",
        date: "2026-08-03",
        time: "19:30",
        status: "scheduled",
    },
];

const initialParticipants: Participant[] = [
    {
        id: "participant-1",
        name: "John Samuel",
        email: "john@example.com",
        hasAccess: true,
    },
    {
        id: "participant-2",
        name: "Anna Mary",
        email: "anna@example.com",
        hasAccess: true,
    },
    {
        id: "participant-3",
        name: "Robert K",
        email: "robert@example.com",
        hasAccess: false,
    },
];

const AdminPage = () => {

    const [sessions, setSessions] =
        useState(initialSessions);

    const [participants, setParticipants] =
        useState(initialParticipants);

    const [title, setTitle] =
        useState("");

    const [date, setDate] =
        useState("");

    const [time, setTime] =
        useState("");

    const [search, setSearch] =
        useState("");

    const filteredParticipants =
        useMemo(() => {

            const query =
                search
                    .trim()
                    .toLowerCase();

            if (!query) {
                return participants;
            }

            return participants.filter(
                (participant) =>
                    participant.name
                        .toLowerCase()
                        .includes(query) ||
                    participant.email
                        .toLowerCase()
                        .includes(query)
            );

        }, [
            participants,
            search,
        ]);

    const handleScheduleSession = () => {

        if (
            !title.trim() ||
            !date ||
            !time
        ) {
            return;
        }

        const session: PrayerSession = {
            id: crypto.randomUUID(),
            title: title.trim(),
            date,
            time,
            status: "scheduled",
        };

        setSessions((current) => [
            ...current,
            session,
        ]);

        setTitle("");
        setDate("");
        setTime("");
    };

    const handleDeleteSession = (
        sessionId: string
    ) => {

        setSessions((current) =>
            current.filter(
                (session) =>
                    session.id !==
                    sessionId
            )
        );

    };

    const handleAccessChange = (
        participantId: string
    ) => {

        setParticipants((current) =>
            current.map(
                (participant) =>
                    participant.id ===
                    participantId
                        ? {
                              ...participant,
                              hasAccess:
                                  !participant.hasAccess,
                          }
                        : participant
            )
        );

    };

    return (

        <Stack gap="2xl">

            {/* Page header */}

            <div>

                <Typography variant="h1">
                    Admin
                </Typography>

                <Typography
                    variant="body"
                    color="muted"
                >
                    Schedule prayer sessions and
                    manage participant access.
                </Typography>

            </div>


            {/* Summary */}

            <div className="grid gap-4 sm:grid-cols-2">

                <Surface
                    bordered
                    padding="lg"
                    radius="lg"
                >

                    <div className="flex items-center gap-4">

                        <CalendarDays className="h-6 w-6 text-primary" />

                        <div>

                            <Typography variant="h3">
                                {sessions.length}
                            </Typography>

                            <Typography
                                variant="body-sm"
                                color="muted"
                            >
                                Scheduled sessions
                            </Typography>

                        </div>

                    </div>

                </Surface>

                <Surface
                    bordered
                    padding="lg"
                    radius="lg"
                >

                    <div className="flex items-center gap-4">

                        <Users className="h-6 w-6 text-primary" />

                        <div>

                            <Typography variant="h3">
                                {
                                    participants.filter(
                                        (participant) =>
                                            participant.hasAccess
                                    ).length
                                }
                            </Typography>

                            <Typography
                                variant="body-sm"
                                color="muted"
                            >
                                Participants with access
                            </Typography>

                        </div>

                    </div>

                </Surface>

            </div>


            {/* Session management */}

            <Surface
                elevation="raised"
                padding="xl"
                radius="xl"
            >

                <Stack gap="lg">

                    <div>

                        <Typography variant="h2">
                            Schedule Prayer Session
                        </Typography>

                        <Typography
                            variant="body-sm"
                            color="muted"
                        >
                            Create upcoming prayer
                            sessions for participants.
                        </Typography>

                    </div>


                    {/* Scheduler */}

                    <div className="grid gap-4 md:grid-cols-3">

                        <Input
                            value={title}
                            placeholder="Session title"
                            onChange={(event) =>
                                setTitle(
                                    event.target.value
                                )
                            }
                        />

                        <Input
                            type="date"
                            value={date}
                            onChange={(event) =>
                                setDate(
                                    event.target.value
                                )
                            }
                        />

                        <Input
                            type="time"
                            value={time}
                            onChange={(event) =>
                                setTime(
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
                            onClick={
                                handleScheduleSession
                            }
                        >
                            Schedule Session
                        </Button>

                    </div>


                    {/* Scheduled sessions */}

                    <div className="border-t border-border pt-6">

                        <Stack gap="md">

                            <Typography variant="h3">
                                Upcoming Sessions
                            </Typography>

                            {sessions.length === 0 ? (

                                <Typography
                                    variant="body-sm"
                                    color="muted"
                                >
                                    No prayer sessions
                                    scheduled.
                                </Typography>

                            ) : (

                                sessions.map(
                                    (session) => (

                                        <div
                                            key={
                                                session.id
                                            }
                                            className="
                                                flex
                                                flex-wrap
                                                items-center
                                                justify-between
                                                gap-4
                                                rounded-lg
                                                border
                                                border-border
                                                p-4
                                            "
                                        >

                                            <Stack gap="xs">

                                                <div className="flex items-center gap-2">

                                                    <Typography variant="title">
                                                        {
                                                            session.title
                                                        }
                                                    </Typography>

                                                    <Badge color="info">
                                                        Scheduled
                                                    </Badge>

                                                </div>

                                                <div className="flex flex-wrap gap-4 text-muted-foreground">

                                                    <span className="flex items-center gap-2 text-sm">

                                                        <CalendarDays className="h-4 w-4" />

                                                        {
                                                            session.date
                                                        }

                                                    </span>

                                                    <span className="flex items-center gap-2 text-sm">

                                                        <Clock className="h-4 w-4" />

                                                        {
                                                            session.time
                                                        }

                                                    </span>

                                                </div>

                                            </Stack>

                                            <Button
                                                variant="ghost"
                                                size="sm"
                                                onClick={() =>
                                                    handleDeleteSession(
                                                        session.id
                                                    )
                                                }
                                            >
                                                <Trash2 className="h-4 w-4" />

                                                Remove
                                            </Button>

                                        </div>

                                    )
                                )

                            )}

                        </Stack>

                    </div>

                </Stack>

            </Surface>


            {/* Participant access */}

            <Surface
                elevation="raised"
                padding="xl"
                radius="xl"
            >

                <Stack gap="lg">

                    <div>

                        <Typography variant="h2">
                            Participant Access
                        </Typography>

                        <Typography
                            variant="body-sm"
                            color="muted"
                        >
                            Control which users can
                            access prayer sessions.
                        </Typography>

                    </div>


                    {/* Search */}

                    <div className="max-w-md">

                        <Input
                            value={search}
                            placeholder="Search participants..."
                            leftIcon={
                                <Search className="h-4 w-4" />
                            }
                            onChange={(event) =>
                                setSearch(
                                    event.target.value
                                )
                            }
                        />

                    </div>


                    {/* Participants */}

                    <div className="space-y-3">

                        {filteredParticipants.map(
                            (participant) => (

                                <div
                                    key={
                                        participant.id
                                    }
                                    className="
                                        flex
                                        flex-wrap
                                        items-center
                                        justify-between
                                        gap-4
                                        rounded-lg
                                        border
                                        border-border
                                        p-4
                                    "
                                >

                                    <div>

                                        <Typography variant="title">
                                            {
                                                participant.name
                                            }
                                        </Typography>

                                        <Typography
                                            variant="body-sm"
                                            color="muted"
                                        >
                                            {
                                                participant.email
                                            }
                                        </Typography>

                                    </div>

                                    <div className="flex items-center gap-3">

                                        <Badge
                                            color={
                                                participant.hasAccess
                                                    ? "success"
                                                    : "danger"
                                            }
                                        >
                                            {participant.hasAccess
                                                ? "Access granted"
                                                : "No access"}
                                        </Badge>

                                        <Button
                                            variant={
                                                participant.hasAccess
                                                    ? "outline"
                                                    : "primary"
                                            }
                                            size="sm"
                                            onClick={() =>
                                                handleAccessChange(
                                                    participant.id
                                                )
                                            }
                                        >
                                            {participant.hasAccess ? (
                                                <UserX className="h-4 w-4" />
                                            ) : (
                                                <UserCheck className="h-4 w-4" />
                                            )}

                                            {participant.hasAccess
                                                ? "Revoke"
                                                : "Grant Access"}
                                        </Button>

                                    </div>

                                </div>

                            )
                        )}

                    </div>

                </Stack>

            </Surface>

        </Stack>

    );

};

export default AdminPage;