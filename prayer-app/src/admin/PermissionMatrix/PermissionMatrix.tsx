// New file generated from PermissionMatrix.tsx
import {
    Card,
    Checkbox,
    Divider,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { permissionMatrixVariants } from "./PermissionMatrix.styles";

import type { PermissionMatrixProps } from "./PermissionMatrix.types";

const PermissionMatrix = ({
    heading,
    description,
    roles,
    permissions,
    values,
    loading = false,
    onPermissionChange,
    className,
    ...props
}: PermissionMatrixProps) => {
    return (
        <Card
            className={cn(
                permissionMatrixVariants(),
                className
            )}
            {...props}
        >
            {(heading || description) && (
                <>
                    <Stack
                        gap="xs"
                        className="p-6"
                    >
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

                    <Divider />
                </>
            )}

            {loading ? (
                <div className="p-6">
                    <Typography>
                        Loading permissions...
                    </Typography>
                </div>
            ) : (
                <div className="overflow-x-auto">
                    <table className="min-w-full border-collapse">
                        <thead>
                            <tr className="border-b">
                                <th className="p-4 text-left">
                                    <Typography
                                        variant="body-sm"
                                        weight="semibold"
                                    >
                                        Permission
                                    </Typography>
                                </th>

                                {roles.map((role) => (
                                    <th
                                        key={role.id}
                                        className="p-4 text-center"
                                    >
                                        <Typography
                                            variant="body-sm"
                                            weight="semibold"
                                        >
                                            {role.label}
                                        </Typography>
                                    </th>
                                ))}
                            </tr>
                        </thead>

                        <tbody>
                            {permissions.map(
                                (permission) => (
                                    <tr
                                        key={
                                            permission.id
                                        }
                                        className="border-b last:border-b-0"
                                    >
                                        <td className="p-4">
                                            <Typography variant="body-sm">
                                                {
                                                    permission.label
                                                }
                                            </Typography>
                                        </td>

                                        {roles.map(
                                            (
                                                role
                                            ) => (
                                                <td
                                                    key={
                                                        role.id
                                                    }
                                                    className="p-4 text-center"
                                                >
                                                    <Checkbox
                                                        checked={
                                                            values[
                                                                role.id
                                                            ]?.[
                                                                permission.id
                                                            ] ??
                                                            false
                                                        }
                                                        onChange={(
                                                            event
                                                        ) =>
                                                            onPermissionChange?.(
                                                                role.id,
                                                                permission.id,
                                                                event
                                                                    .target
                                                                    .checked
                                                            )
                                                        }
                                                    />
                                                </td>
                                            )
                                        )}
                                    </tr>
                                )
                            )}
                        </tbody>
                    </table>
                </div>
            )}
        </Card>
    );
};

export default PermissionMatrix;