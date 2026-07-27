// New file generated from UserRoleSelector.tsx
import {
    Card,
    Checkbox,
    Divider,
    Radio,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { userRoleSelectorVariants } from "./UserRoleSelector.styles";

import type {
    UserRole,
    UserRoleSelectorProps,
} from "./UserRoleSelector.types";

const UserRoleSelector = ({
    heading,
    description,
    roles,
    selectedRoleIds,
    multiple = true,
    loading = false,
    onSelectionChange,
    className,
    ...props
}: UserRoleSelectorProps) => {

    const toggleRole = (
        role: UserRole
    ) => {

        if (multiple) {

            const selected =
                selectedRoleIds.includes(
                    role.id
                );

            onSelectionChange?.(
                selected
                    ? selectedRoleIds.filter(
                          (id) =>
                              id !== role.id
                      )
                    : [
                          ...selectedRoleIds,
                          role.id,
                      ]
            );

            return;
        }

        onSelectionChange?.([
            role.id,
        ]);
    };

    return (
        <Card
            className={cn(
                userRoleSelectorVariants(),
                className
            )}
            {...props}
        >
            <Stack gap="lg">

                {(heading ||
                    description) && (
                    <>
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

                        <Divider />
                    </>
                )}

                {loading ? (
                    <Typography>
                        Loading roles...
                    </Typography>
                ) : (
                    <Stack gap="md">

                        {roles.map(
                            (role) => {

                                const checked =
                                    selectedRoleIds.includes(
                                        role.id
                                    );

                                return (
                                    <div
                                        key={
                                            role.id
                                        }
                                        className="rounded-lg border p-4"
                                    >
                                        {multiple ? (
                                            <Checkbox
                                                checked={
                                                    checked
                                                }
                                                disabled={
                                                    role.disabled
                                                }
                                                label={
                                                    role.name
                                                }
                                                helperText={
                                                    role.description
                                                }
                                                onChange={() =>
                                                    toggleRole(
                                                        role
                                                    )
                                                }
                                            />
                                        ) : (
                                            <Radio
                                                checked={
                                                    checked
                                                }
                                                disabled={
                                                    role.disabled
                                                }
                                                label={
                                                    role.name
                                                }
                                                helperText={
                                                    role.description
                                                }
                                                onChange={() =>
                                                    toggleRole(
                                                        role
                                                    )
                                                }
                                            />
                                        )}
                                    </div>
                                );
                            }
                        )}

                    </Stack>
                )}

            </Stack>
        </Card>
    );
};

export default UserRoleSelector;