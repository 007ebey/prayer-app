// New file generated from SessionTable.tsx
import {
    Pencil,
    Play,
    Square,
    Trash2,
    Eye,
} from "lucide-react";

import {
    Avatar,
    Badge,
    Button,
    Card,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { sessionTableVariants } from "./SessionTable.styles";

import type {
    Session,
    SessionTableProps,
} from "./SessionTable.types";

const statusColor = (
    status: Session["status"]
) => {
    switch (status) {
        case "live":
            return "success";

        case "scheduled":
            return "warning";

        case "completed":
            return "info";

        case "cancelled":
            return "danger";

        default:
            return "neutral";
    }
};

const SessionTable = ({
    sessions,
    loading = false,
    emptyState,
    onView,
    onEdit,
    onDelete,
    onStart,
    onEnd,
    className,
    ...props
}: SessionTableProps) => {

    return (
        <Card
            className={cn(
                sessionTableVariants(),
                className
            )}
            {...props}
        >
            {loading ? (
                <div className="p-6">
                    <Typography>
                        Loading sessions...
                    </Typography>
                </div>
            ) : sessions.length === 0 ? (
                emptyState ?? (
                    <div className="p-6">
                        <Typography>
                            No sessions found.
                        </Typography>
                    </div>
                )
            ) : (
                <div className="overflow-x-auto">
                    <table className="min-w-full">
                        <thead className="border-b bg-muted/50">
                            <tr>
                                <th className="px-6 py-3 text-left">
                                    Session
                                </th>

                                <th className="px-6 py-3 text-left">
                                    Host
                                </th>

                                <th className="px-6 py-3 text-center">
                                    Participants
                                </th>

                                <th className="px-6 py-3 text-left">
                                    Schedule
                                </th>

                                <th className="px-6 py-3 text-center">
                                    Status
                                </th>

                                <th className="px-6 py-3 text-right">
                                    Actions
                                </th>
                            </tr>
                        </thead>

                        <tbody>
                            {sessions.map((session) => (
                                <tr
                                    key={session.id}
                                    className="border-b last:border-0"
                                >
                                    <td className="px-6 py-4">
                                        <Typography
                                            variant="body"
                                            weight="medium"
                                        >
                                            {session.title}
                                        </Typography>
                                    </td>

                                    <td className="px-6 py-4">
                                        <div className="flex items-center gap-3">
                                            <Avatar
                                                name={session.host.name}
                                                {...(session.host.avatar
                                                    ? { src: session.host.avatar }
                                                    : {})}
                                                size="sm"
                                            />
                                            <Typography
                                                variant="body-sm"
                                            >
                                                {
                                                    session.host.name
                                                }
                                            </Typography>
                                        </div>
                                    </td>

                                    <td className="px-6 py-4 text-center">
                                        {session.participants}
                                    </td>

                                    <td className="px-6 py-4">
                                        <Typography
                                            variant="body-sm"
                                        >
                                            {session.startTime}
                                        </Typography>

                                        <Typography
                                            variant="caption"
                                            className="text-muted-foreground"
                                        >
                                            {session.endTime}
                                        </Typography>
                                    </td>

                                    <td className="px-6 py-4 text-center">
                                        <Badge
                                            variant="soft"
                                            color={statusColor(
                                                session.status
                                            )}
                                        >
                                            {session.status}
                                        </Badge>
                                    </td>

                                    <td className="px-6 py-4">
                                        <div className="flex justify-end gap-2">

                                            {onView && (
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    onClick={() =>
                                                        onView(
                                                            session
                                                        )
                                                    }
                                                >
                                                    <Eye size={16} />
                                                </Button>
                                            )}

                                            {session.status ===
                                                "scheduled" &&
                                                onStart && (
                                                    <Button
                                                        variant="ghost"
                                                        size="icon"
                                                        onClick={() =>
                                                            onStart(
                                                                session
                                                            )
                                                        }
                                                    >
                                                        <Play size={16} />
                                                    </Button>
                                                )}

                                            {session.status ===
                                                "live" &&
                                                onEnd && (
                                                    <Button
                                                        variant="ghost"
                                                        size="icon"
                                                        onClick={() =>
                                                            onEnd(
                                                                session
                                                            )
                                                        }
                                                    >
                                                        <Square size={16} />
                                                    </Button>
                                                )}

                                            {onEdit && (
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    onClick={() =>
                                                        onEdit(
                                                            session
                                                        )
                                                    }
                                                >
                                                    <Pencil size={16} />
                                                </Button>
                                            )}

                                            {onDelete && (
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    onClick={() =>
                                                        onDelete(
                                                            session
                                                        )
                                                    }
                                                >
                                                    <Trash2 size={16} />
                                                </Button>
                                            )}

                                        </div>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </Card>
    );
};

export default SessionTable;