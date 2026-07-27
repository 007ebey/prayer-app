import { useEffect } from "react";

import {
    Card,
    Divider,
    Stack,
    Button,
} from "../../../primitives";

import type {
    ContextMenuProps,
    ContextMenuItem,
} from "./ContextMenu.types";

const ContextMenu = ({
    open,
    x,
    y,
    items,
    onClose,
}: ContextMenuProps) => {
    useEffect(() => {
        if (!open) {
            return;
        }

        const handleClick = () => {
            onClose();
        };

        window.addEventListener(
            "click",
            handleClick
        );

        return () => {
            window.removeEventListener(
                "click",
                handleClick
            );
        };
    }, [open, onClose]);

    if (!open) {
        return null;
    }

    return (
        <Card
            className="fixed z-50 min-w-56 p-2 shadow-xl"
            style={{
                left: x,
                top: y,
            }}
        >
            <Stack gap="xs">
                {items.map(
                    (
                        item: ContextMenuItem,
                        index
                    ) => (
                        <div key={item.id}>
                            <Button
                                variant="ghost"
                                fullWidth
                                disabled={
                                    item.disabled
                                }
                                onClick={() => {
                                    item.onClick?.();

                                    onClose();
                                }}
                                className={`justify-start ${
                                    item.danger
                                        ? "text-destructive"
                                        : ""
                                }`}
                            >
                                <span className="flex items-center gap-3">
                                    {item.icon}

                                    {item.label}
                                </span>
                            </Button>

                            {index !==
                                items.length -
                                    1 && (
                                <Divider />
                            )}
                        </div>
                    )
                )}
            </Stack>
        </Card>
    );
};

export default ContextMenu;