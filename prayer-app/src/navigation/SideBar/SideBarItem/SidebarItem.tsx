import {
    Button,
    Flex,
    Typography,
} from "../../../../primitives";

import type { SidebarItemProps } from "./SidebarItem.types";

export const SidebarItem = ({
    label,
    icon,
    badge,
    active = false,
    disabled = false,
    collapsed = false,
    onClick,
}: SidebarItemProps) => {
    return (
        <Button
            variant={active ? "primary" : "ghost"}
            fullWidth
            disabled={disabled}
            onClick={onClick}
            className="justify-start"
        >
            <Flex
                align="center"
                justify="between"
                className="w-full"
            >
                <Flex
                    align="center"
                    gap="md"
                >
                    {icon}

                    {!collapsed && (
                        <Typography variant="body-sm">
                            {label}
                        </Typography>
                    )}
                </Flex>

                {!collapsed && badge}
            </Flex>
        </Button>
    );
};

export default SidebarItem;