// New file generated from UserTable.tsx
import {
    Eye,
    Pencil,
    Shield,
    Trash2,
} from "lucide-react";

import {
    Avatar,
    Badge,
    Button,
    Card,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { userTableVariants } from "./UserTable.styles";

import type {
    User,
    UserRole,
    UserTableProps,
} from "./UserTable.types";

const UserTable = ({
    users,
    loading = false,
    emptyState,
    onView,
    onEdit,
    onDelete,
    onManageRoles,
    className,
    ...props
}: UserTableProps) => {

    return (
        <Card
            className={cn(
                userTableVariants(),
                className
            )}
            {...props}
        >
            {loading ? (
                <div className="p-6">
                    <Typography>
                        Loading users...
                    </Typography>
                </div>
            ) : users.length === 0 ? (
                emptyState ?? (
                    <div className="p-6">
                        <Typography>
                            No users found.
                        </Typography>
                    </div>
                )
            ) : (
                <div className="overflow-x-auto">

                    <table className="min-w-full">

                        <thead className="border-b bg-muted/50">

                            <tr>

                                <th className="px-6 py-3 text-left">
                                    User
                                </th>

                                <th className="px-6 py-3 text-left">
                                    Email
                                </th>

                                <th className="px-6 py-3 text-left">
                                    Roles
                                </th>

                                <th className="px-6 py-3 text-center">
                                    Status
                                </th>

                                <th className="px-6 py-3 text-left">
                                    Last Active
                                </th>

                                <th className="px-6 py-3 text-right">
                                    Actions
                                </th>

                            </tr>

                        </thead>

                        <tbody>

                            {users.map((user) => (

                                <tr
                                    key={user.id}
                                    className="border-b last:border-0"
                                >

                                    <td className="px-6 py-4">

                                        <div className="flex items-center gap-3">

                                            <Avatar
                                                name={user.name}
                                                size="sm"
                                                {...(user.avatar
                                                    ? { src: user.avatar }
                                                    : {})}
                                                {...(user.status
                                                    ? { status: user.status }
                                                    : {})}
                                            />

                                            <Typography
                                                variant="body"
                                                weight="medium"
                                            >
                                                {user.name}
                                            </Typography>

                                        </div>

                                    </td>

                                    <td className="px-6 py-4">

                                        <Typography variant="body-sm">
                                            {user.email}
                                        </Typography>

                                    </td>

                                    <td className="px-6 py-4">

                                        <div className="flex flex-wrap gap-2">

                                            {user.roles.map(
                                                (
                                                    role: UserRole
                                                ) => (
                                                    <Badge
                                                        key={
                                                            role.id
                                                        }
                                                        variant="soft"
                                                        color="primary"
                                                        size="sm"
                                                    >
                                                        {
                                                            role.name
                                                        }
                                                    </Badge>
                                                )
                                            )}

                                        </div>

                                    </td>

                                    <td className="px-6 py-4 text-center">

                                        <Badge
                                            variant="soft"
                                            color={
                                                user.status ===
                                                    "online"
                                                    ? "success"
                                                    : user.status ===
                                                        "away"
                                                        ? "warning"
                                                        : user.status ===
                                                            "busy"
                                                            ? "danger"
                                                            : "neutral"
                                            }
                                            size="sm"
                                        >
                                            {user.status ??
                                                "offline"}
                                        </Badge>

                                    </td>

                                    <td className="px-6 py-4">

                                        <Typography variant="body-sm">
                                            {user.lastActive ??
                                                "-"}
                                        </Typography>

                                    </td>

                                    <td className="px-6 py-4">

                                        <div className="flex justify-end gap-2">

                                            {onView && (
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    onClick={() =>
                                                        onView(
                                                            user
                                                        )
                                                    }
                                                >
                                                    <Eye size={16} />
                                                </Button>
                                            )}

                                            {onManageRoles && (
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    onClick={() =>
                                                        onManageRoles(
                                                            user
                                                        )
                                                    }
                                                >
                                                    <Shield size={16} />
                                                </Button>
                                            )}

                                            {onEdit && (
                                                <Button
                                                    variant="ghost"
                                                    size="icon"
                                                    onClick={() =>
                                                        onEdit(
                                                            user
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
                                                            user
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

export default UserTable;