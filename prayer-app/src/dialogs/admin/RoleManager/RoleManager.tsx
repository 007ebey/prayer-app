// New file generated from RoleManager.tsx
import {
    Pencil,
    Shield,
    Trash2,
    Users,
} from "lucide-react";

import {
    Badge,
    Button,
    Card,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { roleManagerVariants } from "./RoleManager.styles";

import type {
    Role,
    RoleManagerProps,
} from "./RoleManager.types";

const RoleManager = ({
    heading,
    description,
    roles,
    loading = false,
    emptyState,
    onCreateRole,
    onEditRole,
    onDeleteRole,
    onManagePermissions,
    className,
    ...props
}: RoleManagerProps) => {
    return (
        <Card
            className={cn(
                roleManagerVariants(),
                className
            )}
            {...props}
        >
            <Stack
                gap="lg"
                className="p-6"
            >
                <Flex
                    justify="between"
                    align="center"
                >
                    <Stack gap="xs">
                        {heading && (
                            <Typography variant="h4">
                                {heading}
                            </Typography>
                        )}

                        {description && (
                            <Typography
                                variant="caption"
                                className="text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}
                    </Stack>

                    {onCreateRole && (
                        <Button
                            onClick={onCreateRole}
                        >
                            Add Role
                        </Button>
                    )}
                </Flex>

                <Divider />

                {loading ? (
                    <Typography>
                        Loading roles...
                    </Typography>
                ) : roles.length === 0 ? (
                    emptyState ?? (
                        <Typography>
                            No roles available.
                        </Typography>
                    )
                ) : (
                    <Stack gap="md">
                        {roles.map(
                            (role: Role) => (
                                <Card
                                    key={role.id}
                                    className="p-4"
                                >
                                    <Flex
                                        justify="between"
                                        align="center"
                                    >
                                        <Stack gap="xs">

                                            <Flex
                                                gap="sm"
                                                align="center"
                                            >
                                                <Typography
                                                    variant="body"
                                                    weight="semibold"
                                                >
                                                    {role.name}
                                                </Typography>

                                                {role.system && (
                                                    <Badge
                                                        variant="soft"
                                                        color="info"
                                                    >
                                                        System
                                                    </Badge>
                                                )}
                                            </Flex>

                                            {role.description && (
                                                <Typography
                                                    variant="caption"
                                                    className="text-muted-foreground"
                                                >
                                                    {role.description}
                                                </Typography>
                                            )}

                                            {role.userCount !==
                                                undefined && (
                                                <Flex
                                                    gap="xs"
                                                    align="center"
                                                >
                                                    <Users
                                                        size={
                                                            14
                                                        }
                                                    />

                                                    <Typography
                                                        variant="caption"
                                                    >
                                                        {
                                                            role.userCount
                                                        }{" "}
                                                        users
                                                    </Typography>
                                                </Flex>
                                            )}
                                        </Stack>

                                        <Flex gap="sm">

                                            {onManagePermissions && (
                                                <Button
                                                    variant="outline"
                                                    size="sm"
                                                    leftIcon={
                                                        <Shield
                                                            size={
                                                                16
                                                            }
                                                        />
                                                    }
                                                    onClick={() =>
                                                        onManagePermissions(
                                                            role
                                                        )
                                                    }
                                                >
                                                    Permissions
                                                </Button>
                                            )}

                                            {onEditRole && (
                                                <Button
                                                    variant="outline"
                                                    size="sm"
                                                    leftIcon={
                                                        <Pencil
                                                            size={
                                                                16
                                                            }
                                                        />
                                                    }
                                                    onClick={() =>
                                                        onEditRole(
                                                            role
                                                        )
                                                    }
                                                >
                                                    Edit
                                                </Button>
                                            )}

                                            {onDeleteRole &&
                                                !role.system && (
                                                    <Button
                                                        variant="danger"
                                                        size="sm"
                                                        leftIcon={
                                                            <Trash2
                                                                size={
                                                                    16
                                                                }
                                                            />
                                                        }
                                                        onClick={() =>
                                                            onDeleteRole(
                                                                role
                                                            )
                                                        }
                                                    >
                                                        Delete
                                                    </Button>
                                                )}
                                        </Flex>
                                    </Flex>
                                </Card>
                            )
                        )}
                    </Stack>
                )}
            </Stack>
        </Card>
    );
};

export default RoleManager;