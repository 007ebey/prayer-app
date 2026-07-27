import { ChevronDown } from "lucide-react";

import {
    Avatar,
    Button,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import type { UserMenuProps } from "./UserMenu.types";

const UserMenu = ({
    name = "Guest",
    subtitle,
    avatar,
    showName = true,
    onClick,
}: UserMenuProps) => {
    return (
        <Button
            variant="ghost"
            onClick={onClick}
            className="h-auto px-2 py-1"
        >
            <Flex
                align="center"
                gap="sm"
            >
                {avatar ?? (
                    <Avatar
                        name={name}
                    />
                )}

                {showName && (
                    <Stack gap="none">
                        <Typography
                            variant="body-sm"
                            weight="medium"
                        >
                            {name}
                        </Typography>

                        {subtitle && (
                            <Typography
                                variant="caption"
                                className="text-muted-foreground"
                            >
                                {subtitle}
                            </Typography>
                        )}
                    </Stack>
                )}

                <ChevronDown
                    size={16}
                    className="text-muted-foreground"
                />
            </Flex>
        </Button>
    );
};

export default UserMenu;