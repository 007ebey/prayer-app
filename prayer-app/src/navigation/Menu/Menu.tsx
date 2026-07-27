import {
    Button,
    Divider,
    Stack,
    Flex,
} from "../../../primitives";

import type {
    MenuItem,
    MenuProps,
} from "./Menu.types";

const Menu = ({
    items,
    direction = "vertical",
    fullWidth = true,
}: MenuProps) => {
    const Container =
        direction === "vertical"
            ? Stack
            : Flex;

    return (
        <Container
            gap="xs"
            className={
                direction === "horizontal"
                    ? "items-center"
                    : undefined
            }
        >
            {items.map(
                (
                    item: MenuItem,
                    index
                ) => (
                    <div
                        key={item.id}
                        className={
                            fullWidth
                                ? "w-full"
                                : undefined
                        }
                    >
                        <Button
                            variant={
                                item.selected
                                    ? "primary"
                                    : "ghost"
                            }
                            fullWidth={fullWidth}
                            disabled={
                                item.disabled
                            }
                            onClick={
                                item.onClick
                            }
                            className={`justify-between ${
                                item.danger
                                    ? "text-destructive"
                                    : ""
                            }`}
                        >
                            <span className="flex items-center gap-3">
                                {item.icon}

                                {item.label}
                            </span>

                            {item.badge}
                        </Button>

                        {direction ===
                            "vertical" &&
                            index !==
                                items.length -
                                    1 && (
                                <Divider />
                            )}
                    </div>
                )
            )}
        </Container>
    );
};

export default Menu;