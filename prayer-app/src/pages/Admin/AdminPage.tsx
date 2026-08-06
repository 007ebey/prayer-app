import {
    useMemo,
    useState,
} from "react";

import {
    CalendarDays,
    Check,
    Church,
    Clock,
    Plus,
    Search,
    ShieldCheck,
    Trash2,
    UserRound,
    Users,
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
import { PrayerPoints } from "../PrayerPoints";

import { PrayerGroups } from "../PrayerGroups";


/* -------------------------------------------------------------------------- */
/* Types                                                                      */
/* -------------------------------------------------------------------------- */

type AdminSection =
    | "sessions"
    | "prayer-points"
    | "access";

type AdminTab =
    | "users"
    | "groups"
    | "prayer-groups";

type Permission =
    | "prayer.session.view"
    | "prayer.session.join"
    | "prayer.session.host"
    | "session.manage"
    | "user.manage"
    | "group.manage";

interface PrayerSession {
    id: string;

    title: string;

    date: string;

    time: string;

    prayerPointIds: string[];
}

interface PrayerPoint {
    id: string;

    title: string;

    description: string;
}

interface User {
    id: string;

    name: string;

    email: string;

    groupIds: string[];

    prayerGroupIds: string[];
}

interface AccessGroup {
    id: string;

    name: string;

    description: string;

    permissions: Permission[];
}

interface PrayerGroup {
    id: string;

    name: string;

    description: string;

    sessionIds: string[];

    isDefault: boolean;
}

/* -------------------------------------------------------------------------- */
/* Dummy Data                                                                 */
/* -------------------------------------------------------------------------- */

const initialSessions: PrayerSession[] = [
    {
        id: "session-1",
        title: "Morning Prayer",
        date: "2026-08-03",
        time: "06:00",
        prayerPointIds: [
            "prayer-1",
            "prayer-2",
        ],
    },

    {
        id: "session-2",
        title: "Evening Prayer",
        date: "2026-08-03",
        time: "19:30",
        prayerPointIds: [
            "prayer-1",
            "prayer-3",
            "prayer-4",
        ],
    },
];

const initialPrayerPoints: PrayerPoint[] = [
    {
        id: "prayer-1",
        title: "Pray for families",
        description:
            "Pray for unity, restoration, wisdom, and peace in families.",
    },

    {
        id: "prayer-2",
        title: "Pray for healing",
        description:
            "Remember those facing sickness, pain, and difficult circumstances.",
    },

    {
        id: "prayer-3",
        title: "Pray for young people",
        description:
            "Pray for students and young adults as they make important decisions.",
    },

    {
        id: "prayer-4",
        title: "Pray for our communities",
        description:
            "Pray for peace, compassion, and opportunities to serve others.",
    },
];

const initialGroups: AccessGroup[] = [
    {
        id: "members",
        name: "Members",
        description:
            "Standard access to prayer sessions.",
        permissions: [
            "prayer.session.view",
            "prayer.session.join",
        ],
    },

    {
        id: "hosts",
        name: "Prayer Hosts",
        description:
            "Can participate in and host prayer sessions.",
        permissions: [
            "prayer.session.view",
            "prayer.session.join",
            "prayer.session.host",
        ],
    },

    {
        id: "admins",
        name: "Administrators",
        description:
            "Can manage sessions, users, and access groups.",
        permissions: [
            "session.manage",
            "user.manage",
            "group.manage",
        ],
    },
];

const initialUsers: User[] = [
    {
        id: "user-1",
        name: "John Samuel",
        email: "john@example.com",

        groupIds: [
            "members",
        ],

        prayerGroupIds: [
            "general-participants",
        ],
    },

    {
        id: "user-2",
        name: "Anna Mary",
        email: "anna@example.com",

        groupIds: [
            "members",
            "hosts",
        ],

        prayerGroupIds: [
            "general-participants",
            "youth",
        ],
    },

    {
        id: "user-3",
        name: "Robert K",
        email: "robert@example.com",

        groupIds: [],

        prayerGroupIds: [
            "general-participants",
        ],
    },

    {
        id: "user-4",
        name: "David Thomas",
        email: "david@example.com",

        groupIds: [
            "admins",
        ],

        prayerGroupIds: [
            "general-participants",
            "leaders",
        ],
    },
];

const initialPrayerGroups: PrayerGroup[] = [
    {
        id: "general-participants",

        name: "General Participants",

        description:
            "Default prayer group with access to all general prayer sessions.",

        sessionIds: [],

        isDefault: true,
    },

    {
        id: "youth",

        name: "Youth Prayer",

        description:
            "Prayer sessions for the youth prayer community.",

        sessionIds: [
            "session-2",
        ],

        isDefault: false,
    },

    {
        id: "leaders",

        name: "Leadership Prayer",

        description:
            "Private prayer sessions for ministry leaders.",

        sessionIds: [],

        isDefault: false,
    },
];


const permissionLabels: Record<
    Permission,
    string
> = {
    "prayer.session.view":
        "View prayer sessions",

    "prayer.session.join":
        "Join prayer sessions",

    "prayer.session.host":
        "Host prayer sessions",

    "session.manage":
        "Manage prayer sessions",

    "user.manage":
        "Manage users",

    "group.manage":
        "Manage access groups",
};


const permissions =
    Object.keys(
        permissionLabels
    ) as Permission[];



/* -------------------------------------------------------------------------- */
/* Component                                                                  */
/* -------------------------------------------------------------------------- */

const AdminPage = () => {

    /* ---------------------------------------------------------------------- */
    /* Session State                                                          */
    /* ---------------------------------------------------------------------- */

    const [
        sessions,
        setSessions,
    ] = useState(
        initialSessions
    );

    const [
        sessionTitle,
        setSessionTitle,
    ] = useState("");

    const [
        sessionDate,
        setSessionDate,
    ] = useState("");

    const [
        sessionTime,
        setSessionTime,
    ] = useState("");


    /* ---------------------------------------------------------------------- */
    /* Access State                                                           */
    /* ---------------------------------------------------------------------- */

    const [
        users,
        setUsers,
    ] = useState(
        initialUsers
    );

    const [
        groups,
        setGroups,
    ] = useState(
        initialGroups
    );

    const [
        accessTab,
        setAccessTab,
    ] = useState<AdminTab>(
        "users"
    );

    const [
        search,
        setSearch,
    ] = useState("");

    const [
        newGroupName,
        setNewGroupName,
    ] = useState("");

    const [
        prayerPoints,
        setPrayerPoints,
    ] = useState(
        initialPrayerPoints
    );

    const [
        prayerPointTitle,
        setPrayerPointTitle,
    ] = useState("");

    const [
        prayerPointDescription,
        setPrayerPointDescription,
    ] = useState("");


    /* ---------------------------------------------------------------------- */
    /* Derived State                                                          */
    /* ---------------------------------------------------------------------- */

    const filteredUsers =
        useMemo(() => {

            const query =
                search
                    .trim()
                    .toLowerCase();

            if (!query) {
                return users;
            }

            return users.filter(
                (user) =>
                    user.name
                        .toLowerCase()
                        .includes(query) ||
                    user.email
                        .toLowerCase()
                        .includes(query)
            );

        }, [
            users,
            search,
        ]);


    /* ---------------------------------------------------------------------- */
    /* Session Actions                                                        */
    /* ---------------------------------------------------------------------- */

    const handleScheduleSession = () => {

        if (
            !sessionTitle.trim() ||
            !sessionDate ||
            !sessionTime
        ) {
            return;
        }

        const session: PrayerSession = {
            id: crypto.randomUUID(),

            title:
                sessionTitle.trim(),

            date:
                sessionDate,

            time:
                sessionTime,
        };

        setSessions(
            (current) => [
                ...current,
                session,
            ]
        );

        setSessionTitle("");
        setSessionDate("");
        setSessionTime("");

    };


    const handleDeleteSession = (
        sessionId: string
    ) => {

        setSessions(
            (current) =>
                current.filter(
                    (session) =>
                        session.id !==
                        sessionId
                )
        );

    };


    /* ---------------------------------------------------------------------- */
    /* User Group Actions                                                     */
    /* ---------------------------------------------------------------------- */

    const handleAssignGroup = (
        userId: string,
        groupId: string
    ) => {

        setUsers(
            (current) =>
                current.map(
                    (user) => {

                        if (
                            user.id !==
                            userId
                        ) {
                            return user;
                        }

                        if (
                            user.groupIds.includes(
                                groupId
                            )
                        ) {
                            return user;
                        }

                        return {
                            ...user,

                            groupIds: [
                                ...user.groupIds,
                                groupId,
                            ],
                        };

                    }
                )
        );

    };


    const handleRemoveGroup = (
        userId: string,
        groupId: string
    ) => {

        setUsers(
            (current) =>
                current.map(
                    (user) =>
                        user.id ===
                            userId
                            ? {
                                ...user,

                                groupIds:
                                    user.groupIds.filter(
                                        (id) =>
                                            id !==
                                            groupId
                                    ),
                            }
                            : user
                )
        );

    };


    /* ---------------------------------------------------------------------- */
    /* Group Actions                                                          */
    /* ---------------------------------------------------------------------- */

    const handleCreateGroup = () => {

        const name =
            newGroupName.trim();

        if (!name) {
            return;
        }

        const group: AccessGroup = {
            id: crypto.randomUUID(),

            name,

            description:
                "Custom access group.",

            permissions: [],
        };

        setGroups(
            (current) => [
                ...current,
                group,
            ]
        );

        setNewGroupName("");

    };


    const handleTogglePermission = (
        groupId: string,
        permission: Permission
    ) => {

        setGroups(
            (current) =>
                current.map(
                    (group) => {

                        if (
                            group.id !==
                            groupId
                        ) {
                            return group;
                        }

                        const hasPermission =
                            group.permissions.includes(
                                permission
                            );

                        return {
                            ...group,

                            permissions:
                                hasPermission
                                    ? group.permissions.filter(
                                        (item) =>
                                            item !==
                                            permission
                                    )
                                    : [
                                        ...group.permissions,
                                        permission,
                                    ],
                        };

                    }
                )
        );

    };

    const handleCreatePrayerPoint = () => {

        const title =
            prayerPointTitle.trim();

        const description =
            prayerPointDescription.trim();

        if (!title) {
            return;
        }

        const prayerPoint: PrayerPoint = {
            id: crypto.randomUUID(),

            title,

            description,
        };

        setPrayerPoints(
            (current) => [
                ...current,
                prayerPoint,
            ]
        );

        setPrayerPointTitle("");
        setPrayerPointDescription("");

    };

    // const handleAssignPrayerPoint = (
    //     sessionId: string,
    //     prayerPointId: string
    // ) => {

    //     setSessions(
    //         (current) =>
    //             current.map(
    //                 (session) => {

    //                     if (
    //                         session.id !==
    //                         sessionId
    //                     ) {
    //                         return session;
    //                     }

    //                     if (
    //                         session.prayerPointIds.includes(
    //                             prayerPointId
    //                         )
    //                     ) {
    //                         return session;
    //                     }

    //                     return {
    //                         ...session,

    //                         prayerPointIds: [
    //                             ...session.prayerPointIds,
    //                             prayerPointId,
    //                         ],
    //                     };

    //                 }
    //             )
    //     );

    // };

    // const handleRemovePrayerPoint = (
    //     sessionId: string,
    //     prayerPointId: string
    // ) => {

    //     setSessions(
    //         (current) =>
    //             current.map(
    //                 (session) =>
    //                     session.id === sessionId
    //                         ? {
    //                             ...session,

    //                             prayerPointIds:
    //                                 session.prayerPointIds.filter(
    //                                     (id) =>
    //                                         id !==
    //                                         prayerPointId
    //                                 ),
    //                         }
    //                         : session
    //             )
    //     );

    // };

    const handleAssignUserToPrayerGroup = (
        prayerGroupId: string,
        userId: string
    ) => {

        setUsers(
            (current) =>
                current.map((user) => {

                    if (user.id !== userId) {
                        return user;
                    }

                    if (
                        user.prayerGroupIds.includes(
                            prayerGroupId
                        )
                    ) {
                        return user;
                    }

                    return {
                        ...user,

                        prayerGroupIds: [
                            ...user.prayerGroupIds,
                            prayerGroupId,
                        ],
                    };

                })
        );

    };


    const handleRemoveUserFromPrayerGroup = (
        prayerGroupId: string,
        userId: string
    ) => {

        setUsers(
            (current) =>
                current.map((user) => {

                    if (user.id !== userId) {
                        return user;
                    }

                    return {
                        ...user,

                        prayerGroupIds:
                            user.prayerGroupIds.filter(
                                (id) =>
                                    id !== prayerGroupId
                            ),
                    };

                })
        );

    };

    /* ---------------------------------------------------------------------- */
    /* Prayer Group Actions                                                   */
    /* ---------------------------------------------------------------------- */

    const [
        prayerGroups,
        setPrayerGroups,
    ] = useState<PrayerGroup[]>(
        initialPrayerGroups
    );

    const [
        newPrayerGroupName,
        setNewPrayerGroupName,
    ] = useState("");

    const [
        newPrayerGroupDescription,
        setNewPrayerGroupDescription,
    ] = useState("");

    const handleCreatePrayerGroup = () => {

        const name =
            newPrayerGroupName.trim();

        if (!name) {
            return;
        }

        setPrayerGroups(
            (current) => [
                ...current,
                {
                    id: crypto.randomUUID(),

                    name,

                    description:
                        newPrayerGroupDescription.trim(),

                    sessionIds: [],

                    isDefault: false,
                },
            ]
        );

        setNewPrayerGroupName("");

        setNewPrayerGroupDescription("");

    };

    const handleAssignSessionToPrayerGroup = (
        prayerGroupId: string,
        sessionId: string
    ) => {

        setPrayerGroups(
            (current) =>
                current.map((group) => {

                    if (
                        group.id !== prayerGroupId ||
                        group.isDefault
                    ) {
                        return group;
                    }

                    if (
                        group.sessionIds.includes(
                            sessionId
                        )
                    ) {
                        return group;
                    }

                    return {
                        ...group,

                        sessionIds: [
                            ...group.sessionIds,
                            sessionId,
                        ],
                    };

                })
        );

    };

    const handleRemoveSessionFromPrayerGroup = (
        prayerGroupId: string,
        sessionId: string
    ) => {

        setPrayerGroups(
            (current) =>
                current.map((group) => {

                    if (
                        group.id !== prayerGroupId ||
                        group.isDefault
                    ) {
                        return group;
                    }

                    return {
                        ...group,

                        sessionIds:
                            group.sessionIds.filter(
                                (id) =>
                                    id !== sessionId
                            ),
                    };

                })
        );

    };

    /* ---------------------------------------------------------------------- */
    /* UI                                                                     */
    /* ---------------------------------------------------------------------- */

    return (

        <div className="min-h-screen bg-background">

            <div className="mx-auto max-w-7xl p-8">

                <Stack gap="2xl">


                    {/* Page Header */}

                    <div>

                        <Typography variant="h1">
                            Administration
                        </Typography>

                        <Typography
                            variant="body"
                            color="muted"
                        >
                            Manage prayer sessions,
                            users, groups, and access.
                        </Typography>

                    </div>


                    {/* Summary */}

                    <div className="grid gap-4 md:grid-cols-3">

                        <Surface
                            bordered
                            padding="lg"
                            radius="lg"
                        >

                            <div className="flex items-center gap-4">

                                <CalendarDays className="h-6 w-6 text-primary" />

                                <div>

                                    <Typography variant="h3">
                                        {
                                            sessions.length
                                        }
                                    </Typography>

                                    <Typography
                                        variant="body-sm"
                                        color="muted"
                                    >
                                        Scheduled Sessions
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

                                <UserRound className="h-6 w-6 text-primary" />

                                <div>

                                    <Typography variant="h3">
                                        {
                                            users.length
                                        }
                                    </Typography>

                                    <Typography
                                        variant="body-sm"
                                        color="muted"
                                    >
                                        Users
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

                                <ShieldCheck className="h-6 w-6 text-primary" />

                                <div>

                                    <Typography variant="h3">
                                        {
                                            groups.length
                                        }
                                    </Typography>

                                    <Typography
                                        variant="body-sm"
                                        color="muted"
                                    >
                                        Access Groups
                                    </Typography>

                                </div>

                            </div>

                        </Surface>

                    </div>


                    {/* ====================================================== */}
                    {/* SESSION MANAGEMENT                                     */}
                    {/* ====================================================== */}

                    <Surface
                        elevation="raised"
                        padding="xl"
                        radius="xl"
                    >

                        <Stack gap="lg">

                            <div>

                                <Typography variant="h2">
                                    Prayer Sessions
                                </Typography>

                                <Typography
                                    variant="body-sm"
                                    color="muted"
                                >
                                    Schedule and manage
                                    upcoming prayer sessions.
                                </Typography>

                            </div>


                            {/* Scheduler */}

                            <div className="grid gap-4 md:grid-cols-3">

                                <Input
                                    value={
                                        sessionTitle
                                    }
                                    placeholder="Session title"
                                    onChange={(
                                        event
                                    ) =>
                                        setSessionTitle(
                                            event.target
                                                .value
                                        )
                                    }
                                />

                                <Input
                                    type="date"
                                    value={
                                        sessionDate
                                    }
                                    onChange={(
                                        event
                                    ) =>
                                        setSessionDate(
                                            event.target
                                                .value
                                        )
                                    }
                                />

                                <Input
                                    type="time"
                                    value={
                                        sessionTime
                                    }
                                    onChange={(
                                        event
                                    ) =>
                                        setSessionTime(
                                            event.target
                                                .value
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


                            {/* Upcoming */}

                            <div className="border-t border-border pt-6">

                                <Stack gap="md">

                                    <Typography variant="h3">
                                        Upcoming Sessions
                                    </Typography>

                                    {sessions.length ===
                                        0 ? (

                                        <Typography
                                            variant="body-sm"
                                            color="muted"
                                        >
                                            No sessions
                                            scheduled.
                                        </Typography>

                                    ) : (

                                        sessions.map(
                                            (
                                                session
                                            ) => (

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

                                                        <Typography variant="title">
                                                            {
                                                                session.title
                                                            }
                                                        </Typography>

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


                    {/* ====================================================== */}
                    {/* ACCESS MANAGEMENT                                      */}
                    {/* ====================================================== */}

                    <Surface
                        elevation="raised"
                        padding="xl"
                        radius="xl"
                    >

                        <Stack gap="lg">

                            <div>

                                <Typography variant="h2">
                                    Access Management
                                </Typography>

                                <Typography
                                    variant="body-sm"
                                    color="muted"
                                >
                                    Assign users to groups
                                    and control the permissions
                                    granted to each group.
                                </Typography>

                            </div>


                            {/* Tabs */}

                            <div className="flex gap-2 border-b border-border">

                                <Button
                                    variant={
                                        accessTab ===
                                            "users"
                                            ? "primary"
                                            : "ghost"
                                    }
                                    onClick={() =>
                                        setAccessTab(
                                            "users"
                                        )
                                    }
                                >
                                    <Users className="h-4 w-4" />

                                    Users
                                </Button>


                                <Button
                                    variant={
                                        accessTab ===
                                            "groups"
                                            ? "primary"
                                            : "ghost"
                                    }
                                    onClick={() =>
                                        setAccessTab(
                                            "groups"
                                        )
                                    }
                                >
                                    <ShieldCheck className="h-4 w-4" />

                                    Groups
                                </Button>

                                <Button
                                    variant={
                                        accessTab === "prayer-groups"
                                            ? "primary"
                                            : "ghost"
                                    }
                                    onClick={() =>
                                        setAccessTab("prayer-groups")
                                    }
                                >
                                    <Church className="h-4 w-4" />

                                    Prayer Groups
                                </Button>

                            </div>


                            {/* ================================================== */}
                            {/* USERS                                              */}
                            {/* ================================================== */}

                            {accessTab ===
                                "users" && (

                                    <Stack gap="lg">


                                        {/* Search */}

                                        <div className="max-w-md">

                                            <Input
                                                value={
                                                    search
                                                }
                                                placeholder="Search users..."
                                                leftIcon={
                                                    <Search className="h-4 w-4" />
                                                }
                                                onChange={(
                                                    event
                                                ) =>
                                                    setSearch(
                                                        event
                                                            .target
                                                            .value
                                                    )
                                                }
                                            />

                                        </div>


                                        {/* User List */}

                                        <div className="space-y-4">

                                            {filteredUsers.map(
                                                (
                                                    user
                                                ) => {

                                                    const availableGroups =
                                                        groups.filter(
                                                            (
                                                                group
                                                            ) =>
                                                                !user.groupIds.includes(
                                                                    group.id
                                                                )
                                                        );

                                                    return (

                                                        <Surface
                                                            key={
                                                                user.id
                                                            }
                                                            bordered
                                                            padding="lg"
                                                            radius="lg"
                                                        >

                                                            <Stack gap="md">

                                                                <div>

                                                                    <Typography variant="title">
                                                                        {
                                                                            user.name
                                                                        }
                                                                    </Typography>

                                                                    <Typography
                                                                        variant="body-sm"
                                                                        color="muted"
                                                                    >
                                                                        {
                                                                            user.email
                                                                        }
                                                                    </Typography>

                                                                </div>


                                                                {/* Assigned Groups */}

                                                                <div className="flex flex-wrap gap-2">

                                                                    {user
                                                                        .groupIds
                                                                        .length ===
                                                                        0 ? (

                                                                        <Typography
                                                                            variant="body-sm"
                                                                            color="muted"
                                                                        >
                                                                            No groups
                                                                            assigned.
                                                                        </Typography>

                                                                    ) : (

                                                                        user.groupIds.map(
                                                                            (
                                                                                groupId
                                                                            ) => {

                                                                                const group =
                                                                                    groups.find(
                                                                                        (
                                                                                            item
                                                                                        ) =>
                                                                                            item.id ===
                                                                                            groupId
                                                                                    );

                                                                                if (
                                                                                    !group
                                                                                ) {
                                                                                    return null;
                                                                                }

                                                                                return (

                                                                                    <div
                                                                                        key={
                                                                                            group.id
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
                                                                                            {
                                                                                                group.name
                                                                                            }
                                                                                        </span>

                                                                                        <button
                                                                                            type="button"
                                                                                            className="
                                                                                            text-muted-foreground
                                                                                            hover:text-foreground
                                                                                        "
                                                                                            onClick={() =>
                                                                                                handleRemoveGroup(
                                                                                                    user.id,
                                                                                                    group.id
                                                                                                )
                                                                                            }
                                                                                        >
                                                                                            <X className="h-3 w-3" />
                                                                                        </button>

                                                                                    </div>

                                                                                );

                                                                            }
                                                                        )

                                                                    )}

                                                                </div>


                                                                {/* Available Groups */}

                                                                {availableGroups.length >
                                                                    0 && (

                                                                        <div>

                                                                            <Typography
                                                                                variant="caption"
                                                                                color="muted"
                                                                            >
                                                                                Assign group
                                                                            </Typography>

                                                                            <div className="mt-2 flex flex-wrap gap-2">

                                                                                {availableGroups.map(
                                                                                    (
                                                                                        group
                                                                                    ) => (

                                                                                        <Button
                                                                                            key={
                                                                                                group.id
                                                                                            }
                                                                                            variant="outline"
                                                                                            size="sm"
                                                                                            onClick={() =>
                                                                                                handleAssignGroup(
                                                                                                    user.id,
                                                                                                    group.id
                                                                                                )
                                                                                            }
                                                                                        >
                                                                                            <Plus className="h-3 w-3" />

                                                                                            {
                                                                                                group.name
                                                                                            }

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

                                )}


                            {/* ================================================== */}
                            {/* GROUPS                                             */}
                            {/* ================================================== */}

                            {accessTab ===
                                "groups" && (

                                    <Stack gap="lg">


                                        {/* Create Group */}

                                        <div className="flex max-w-xl gap-3">

                                            <Input
                                                value={
                                                    newGroupName
                                                }
                                                placeholder="New group name"
                                                onChange={(
                                                    event
                                                ) =>
                                                    setNewGroupName(
                                                        event
                                                            .target
                                                            .value
                                                    )
                                                }
                                            />

                                            <Button
                                                onClick={
                                                    handleCreateGroup
                                                }
                                            >
                                                <Plus className="h-4 w-4" />

                                                Create
                                            </Button>

                                        </div>


                                        {/* Groups */}

                                        <div className="grid gap-4 lg:grid-cols-2">

                                            {groups.map(
                                                (
                                                    group
                                                ) => {

                                                    const memberCount =
                                                        users.filter(
                                                            (
                                                                user
                                                            ) =>
                                                                user.groupIds.includes(
                                                                    group.id
                                                                )
                                                        ).length;

                                                    return (

                                                        <Surface
                                                            key={
                                                                group.id
                                                            }
                                                            bordered
                                                            padding="lg"
                                                            radius="lg"
                                                        >

                                                            <Stack gap="lg">

                                                                {/* Header */}

                                                                <div className="flex items-start justify-between gap-4">

                                                                    <div>

                                                                        <Typography variant="title">
                                                                            {
                                                                                group.name
                                                                            }
                                                                        </Typography>

                                                                        <Typography
                                                                            variant="body-sm"
                                                                            color="muted"
                                                                        >
                                                                            {
                                                                                group.description
                                                                            }
                                                                        </Typography>

                                                                    </div>

                                                                    <Badge color="info">
                                                                        {
                                                                            memberCount
                                                                        }{" "}
                                                                        users
                                                                    </Badge>

                                                                </div>


                                                                {/* Permissions */}

                                                                <Stack gap="sm">

                                                                    <Typography variant="body-sm">
                                                                        Permissions
                                                                    </Typography>

                                                                    {permissions.map(
                                                                        (
                                                                            permission
                                                                        ) => {

                                                                            const enabled =
                                                                                group.permissions.includes(
                                                                                    permission
                                                                                );

                                                                            return (

                                                                                <button
                                                                                    key={
                                                                                        permission
                                                                                    }
                                                                                    type="button"
                                                                                    className="
                                                                                    flex
                                                                                    items-center
                                                                                    justify-between
                                                                                    gap-4
                                                                                    rounded-md
                                                                                    border
                                                                                    border-border
                                                                                    p-3
                                                                                    text-left
                                                                                "
                                                                                    onClick={() =>
                                                                                        handleTogglePermission(
                                                                                            group.id,
                                                                                            permission
                                                                                        )
                                                                                    }
                                                                                >

                                                                                    <span className="text-sm">
                                                                                        {
                                                                                            permissionLabels[
                                                                                            permission
                                                                                            ]
                                                                                        }
                                                                                    </span>

                                                                                    <span
                                                                                        className={`
                                                                                        flex
                                                                                        h-5
                                                                                        w-5
                                                                                        items-center
                                                                                        justify-center
                                                                                        rounded
                                                                                        border
                                                                                        ${enabled
                                                                                                ? "bg-primary text-primary-foreground"
                                                                                                : "border-border"
                                                                                            }
                                                                                    `}
                                                                                    >

                                                                                        {enabled && (
                                                                                            <Check className="h-3 w-3" />
                                                                                        )}

                                                                                    </span>

                                                                                </button>

                                                                            );

                                                                        }
                                                                    )}

                                                                </Stack>

                                                            </Stack>

                                                        </Surface>

                                                    );

                                                }
                                            )}

                                        </div>

                                    </Stack>

                                )}

                            {/* ================================================== */}
                            {/* PRAYER GROUPS                                      */}
                            {/* ================================================== */}

                            {accessTab === "prayer-groups" && (
                                <PrayerGroups
                                    prayerGroups={prayerGroups}
                                    sessions={sessions}
                                    users={users}

                                    newGroupName={newPrayerGroupName}
                                    newGroupDescription={
                                        newPrayerGroupDescription
                                    }

                                    onNameChange={
                                        setNewPrayerGroupName
                                    }
                                    onDescriptionChange={
                                        setNewPrayerGroupDescription
                                    }

                                    onCreate={
                                        handleCreatePrayerGroup
                                    }

                                    onAssignSession={
                                        handleAssignSessionToPrayerGroup
                                    }

                                    onRemoveSession={
                                        handleRemoveSessionFromPrayerGroup
                                    }

                                    onAssignUser={
                                        handleAssignUserToPrayerGroup
                                    }

                                    onRemoveUser={
                                        handleRemoveUserFromPrayerGroup
                                    }
                                />
                            )}

                        </Stack>

                    </Surface>

                    {/* ====================================================== */}
                    {/* Prayer Points Management                               */}
                    {/* ====================================================== */}

                    <Surface
                        elevation="raised"
                        padding="xl"
                        radius="xl"
                    >
                        <PrayerPoints
                            prayerPoints={prayerPoints}
                            sessions={sessions}

                            title={prayerPointTitle}
                            description={prayerPointDescription}

                            onTitleChange={
                                setPrayerPointTitle
                            }
                            onDescriptionChange={
                                setPrayerPointDescription
                            }

                            onCreate={
                                handleCreatePrayerPoint
                            }
                        />
                    </Surface>


                </Stack>

            </div>

        </div>

    );

};

export default AdminPage;