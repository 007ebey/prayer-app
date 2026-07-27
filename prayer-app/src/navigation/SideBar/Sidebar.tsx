import {
    Card,
    Divider,
    Stack,
} from "../../../primitives";

import SidebarItem from "./SideBarItem/SidebarItem";

import type { SidebarProps } from "./Sidebar.types";

const Sidebar = ({
    items,
    selectedItem,
    footer,
    onItemClick,
}: SidebarProps) => {
    return (
        <Card className="flex h-screen w-72 flex-col rounded-none">
            <Stack
                gap="md"
                className="h-full p-4"
            >
                <Stack
                    gap="xs"
                    className="flex-1"
                >
                    {items.map((item) => (
                        <SidebarItem
                            key={item.id}
                            {...item}
                            active={item.id === selectedItem}
                            onClick={() => onItemClick?.(item)}
                        />
                    ))}
                </Stack>

                {footer && (
                    <>
                        <Divider />

                        {footer}
                    </>
                )}
            </Stack>
        </Card>
    );
};

export default Sidebar;