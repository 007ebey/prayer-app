import {
    Plus,
    X,
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
    PrayerGroupsProps,
} from "./PrayerGroups.types";


const PrayerGroups = ({
    prayerGroups,
    sessions,
    users,

    newGroupName,
    newGroupDescription,

    onNameChange,
    onDescriptionChange,

    onCreate,

    onAssignSession,
    onRemoveSession,

    onAssignUser,
    onRemoveUser,
}: PrayerGroupsProps) => {

    return (

        <Stack gap="lg">

            {/* Header */}

            <div>

                <Typography variant="h3">
                    Prayer Groups
                </Typography>

                <Typography
                    variant="body-sm"
                    color="muted"
                >
                    Control which prayer sessions
                    are visible to different groups
                    of participants.
                </Typography>

            </div>


            {/* Create */}

            <Surface
                bordered
                padding="lg"
                radius="lg"
            >

                <Stack gap="md">

                    <Typography variant="title">
                        Create Prayer Group
                    </Typography>

                    <div className="grid gap-4 md:grid-cols-2">

                        <Input
                            value={
                                newGroupName
                            }
                            placeholder="Prayer group name"
                            onChange={(event) =>
                                onNameChange(
                                    event.target.value
                                )
                            }
                        />

                        <Input
                            value={
                                newGroupDescription
                            }
                            placeholder="Description"
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
                            Create Prayer Group
                        </Button>

                    </div>

                </Stack>

            </Surface>


            {/* Prayer Groups */}

            <div className="grid gap-4">

                {prayerGroups.map(
                    (group) => {

                        const assignedUsers =
                            users.filter(
                                (user) =>
                                    user.prayerGroupIds.includes(
                                        group.id
                                    )
                            );

                        const availableUsers =
                            users.filter(
                                (user) =>
                                    !user.prayerGroupIds.includes(
                                        group.id
                                    )
                            );

                        const assignedSessions =
                            sessions.filter(
                                (session) =>
                                    group.sessionIds.includes(
                                        session.id
                                    )
                            );

                        const availableSessions =
                            sessions.filter(
                                (session) =>
                                    !group.sessionIds.includes(
                                        session.id
                                    )
                            );

                        return (

                            <Surface
                                key={group.id}
                                bordered
                                padding="lg"
                                radius="lg"
                            >

                                <Stack gap="lg">

                                    {/* Group Header */}

                                    <div className="flex flex-wrap items-start justify-between gap-4">

                                        <div>

                                            <div className="flex flex-wrap items-center gap-2">

                                                <Typography variant="title">
                                                    {group.name}
                                                </Typography>

                                                {group.isDefault && (

                                                    <Badge color="success">
                                                        Default
                                                    </Badge>

                                                )}

                                            </div>

                                            <Typography
                                                variant="body-sm"
                                                color="muted"
                                            >
                                                {group.description}
                                            </Typography>

                                        </div>

                                        <Badge color="info">

                                            {group.isDefault
                                                ? "All general sessions"
                                                : `${assignedSessions.length} sessions`}

                                        </Badge>

                                    </div>


                                    {/* Default group */}

                                    {group.isDefault ? (

                                        <Surface
                                            variant="muted"
                                            padding="md"
                                            radius="md"
                                        >

                                            <Typography
                                                variant="body-sm"
                                                color="muted"
                                            >
                                                Members of this group
                                                can view all general
                                                prayer sessions.
                                            </Typography>

                                        </Surface>

                                    ) : (

                                        <Stack gap="md">

                                            {/* Assigned Sessions */}

                                            <div>

                                                <Typography
                                                    variant="caption"
                                                    color="muted"
                                                >
                                                    ASSIGNED SESSIONS
                                                </Typography>

                                                {assignedSessions.length === 0 ? (

                                                    <Typography
                                                        variant="body-sm"
                                                        color="muted"
                                                        className="mt-2"
                                                    >
                                                        No sessions assigned.
                                                    </Typography>

                                                ) : (

                                                    <div className="mt-2 flex flex-wrap gap-2">

                                                        {assignedSessions.map(
                                                            (session) => (

                                                                <div
                                                                    key={
                                                                        session.id
                                                                    }
                                                                    className="
                                                                        flex
                                                                        items-center
                                                                        gap-2
                                                                        rounded-full
                                                                        border
                                                                        border-border
                                                                        px-3
                                                                        py-1
                                                                    "
                                                                >

                                                                    <span className="text-sm">
                                                                        {session.title}
                                                                    </span>

                                                                    <button
                                                                        type="button"
                                                                        onClick={() =>
                                                                            onRemoveSession(
                                                                                group.id,
                                                                                session.id
                                                                            )
                                                                        }
                                                                    >
                                                                        <X className="h-3 w-3" />
                                                                    </button>

                                                                </div>

                                                            )
                                                        )}

                                                    </div>

                                                )}

                                            </div>


                                            {/* Available Sessions */}

                                            {availableSessions.length > 0 && (

                                                <div>

                                                    <Typography
                                                        variant="caption"
                                                        color="muted"
                                                    >
                                                        ASSIGN SESSION
                                                    </Typography>

                                                    <div className="mt-2 flex flex-wrap gap-2">

                                                        {availableSessions.map(
                                                            (session) => (

                                                                <Button
                                                                    key={
                                                                        session.id
                                                                    }
                                                                    variant="outline"
                                                                    size="sm"
                                                                    onClick={() =>
                                                                        onAssignSession(
                                                                            group.id,
                                                                            session.id
                                                                        )
                                                                    }
                                                                >

                                                                    <Plus className="h-3 w-3" />

                                                                    {session.title}

                                                                </Button>

                                                            )
                                                        )}

                                                    </div>

                                                </div>

                                            )}

                                        </Stack>

                                    )}

                                </Stack>

                                {/* Users */}

                                <Stack gap="md">

                                    <div className="flex items-center justify-between gap-4">

                                        <div>

                                            <Typography variant="body-sm">
                                                Members
                                            </Typography>

                                            <Typography
                                                variant="caption"
                                                color="muted"
                                            >
                                                Users belonging to this
                                                prayer group.
                                            </Typography>

                                        </div>

                                        <Badge color="info">
                                            {assignedUsers.length}{" "}
                                            {assignedUsers.length === 1
                                                ? "user"
                                                : "users"}
                                        </Badge>

                                    </div>


                                    {/* Assigned Users */}

                                    {assignedUsers.length === 0 ? (

                                        <Typography
                                            variant="body-sm"
                                            color="muted"
                                        >
                                            No users assigned.
                                        </Typography>

                                    ) : (

                                        <div className="flex flex-wrap gap-2">

                                            {assignedUsers.map(
                                                (user) => (

                                                    <div
                                                        key={user.id}
                                                        className="
                            flex
                            items-center
                            gap-2
                            rounded-full
                            border
                            border-border
                            px-3
                            py-1
                        "
                                                    >

                                                        <span className="text-sm">
                                                            {user.name}
                                                        </span>

                                                        {!group.isDefault && (

                                                            <button
                                                                type="button"
                                                                className="
                                    text-muted-foreground
                                    transition-colors
                                    hover:text-foreground
                                "
                                                                onClick={() =>
                                                                    onRemoveUser(
                                                                        group.id,
                                                                        user.id
                                                                    )
                                                                }
                                                                aria-label={
                                                                    `Remove ${user.name} from ${group.name}`
                                                                }
                                                            >
                                                                <X className="h-3 w-3" />
                                                            </button>

                                                        )}

                                                    </div>

                                                )
                                            )}

                                        </div>

                                    )}


                                    {/* Assign Users */}

                                    {!group.isDefault &&
                                        availableUsers.length > 0 && (

                                            <div>

                                                <Typography
                                                    variant="caption"
                                                    color="muted"
                                                >
                                                    ASSIGN USER
                                                </Typography>

                                                <div className="mt-2 flex flex-wrap gap-2">

                                                    {availableUsers.map(
                                                        (user) => (

                                                            <Button
                                                                key={user.id}
                                                                variant="outline"
                                                                size="sm"
                                                                onClick={() =>
                                                                    onAssignUser(
                                                                        group.id,
                                                                        user.id
                                                                    )
                                                                }
                                                            >

                                                                <Plus className="h-3 w-3" />

                                                                {user.name}

                                                            </Button>

                                                        )
                                                    )}

                                                </div>

                                            </div>

                                        )}

                                </Stack>

                            </Surface>

                        );

                    }
                )}

            </div>

        </Stack>

    );

};

export default PrayerGroups;