import {
    Button,
    Divider,
    Flex,
} from "../../../primitives";

import type {
    TabItem,
    TabsProps,
} from "./Tabs.types";

const Tabs = ({
    items,
    activeTab,
    onTabChange,
    fullWidth = false,
}: TabsProps) => {
    return (
        <div className="w-full">
            <Flex
                gap="none"
                className={
                    fullWidth
                        ? "w-full"
                        : undefined
                }
            >
                {items.map(
                    (
                        item: TabItem
                    ) => (
                        <Button
                            key={item.id}
                            variant={
                                item.id ===
                                activeTab
                                    ? "primary"
                                    : "ghost"
                            }
                            disabled={
                                item.disabled
                            }
                            fullWidth={
                                fullWidth
                            }
                            onClick={() =>
                                onTabChange?.(
                                    item
                                )
                            }
                            className="rounded-none border-b-2 border-transparent data-[active=true]:border-primary"
                            data-active={
                                item.id ===
                                activeTab
                            }
                        >
                            <Flex
                                align="center"
                                gap="sm"
                            >
                                {item.icon}

                                {item.label}

                                {item.badge}
                            </Flex>
                        </Button>
                    )
                )}
            </Flex>

            <Divider />
        </div>
    );
};

export default Tabs;