import { useEffect } from "react";

import {
    Button,
    Card,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import type {
    NotificationItem,
    NotificationPopupProps,
} from "./NotificationPopup.types";

const NotificationPopup = ({
    open,
    notifications,
    title = "Notifications",
    emptyMessage = "You're all caught up.",
    onClose,
    onViewAll,
    onMarkAllRead,
}: NotificationPopupProps) => {

    useEffect(() => {
        if (!open) {
            return;
        }

        const handleClick = () => {
            onClose?.();
        };

        window.addEventListener(
            "click",
            handleClick
        );

        return () =>
            window.removeEventListener(
                "click",
                handleClick
            );
    }, [open, onClose]);

    if (!open) {
        return null;
    }

    return (
        <Card className="absolute right-0 top-full z-50 mt-2 w-96 shadow-xl">
            <Stack gap="none">

                <Flex
                    justify="between"
                    align="center"
                    className="p-4"
                >
                    <Typography variant="title">
                        {title}
                    </Typography>

                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={onMarkAllRead}
                    >
                        Mark all read
                    </Button>
                </Flex>

                <Divider />

                <div className="max-h-96 overflow-y-auto">

                    {notifications.length === 0 ? (
                        <div className="p-6 text-center">
                            <Typography variant="body-sm">
                                {emptyMessage}
                            </Typography>
                        </div>
                    ) : (
                        notifications.map(
                            (
                                notification: NotificationItem
                            ) => (
                                <button
                                    key={
                                        notification.id
                                    }
                                    type="button"
                                    onClick={
                                        notification.onClick
                                    }
                                    className="flex w-full items-start gap-3 p-4 text-left hover:bg-muted"
                                >
                                    {notification.icon}

                                    <Stack
                                        gap="xs"
                                        className="flex-1"
                                    >
                                        <Flex
                                            justify="between"
                                            align="start"
                                        >
                                            <Typography
                                                variant="body-sm"
                                                weight={
                                                    notification.unread
                                                        ? "semibold"
                                                        : "normal"
                                                }
                                            >
                                                {
                                                    notification.title
                                                }
                                            </Typography>

                                            {notification.unread && (
                                                <span className="mt-1 h-2 w-2 rounded-full bg-primary" />
                                            )}
                                        </Flex>

                                        {notification.description && (
                                            <Typography variant="caption">
                                                {
                                                    notification.description
                                                }
                                            </Typography>
                                        )}

                                        {notification.timestamp && (
                                            <Typography
                                                variant="caption"
                                                className="text-muted-foreground"
                                            >
                                                {
                                                    notification.timestamp
                                                }
                                            </Typography>
                                        )}
                                    </Stack>
                                </button>
                            )
                        )
                    )}

                </div>

                <Divider />

                <div className="p-3">
                    <Button
                        variant="outline"
                        fullWidth
                        onClick={onViewAll}
                    >
                        View All Notifications
                    </Button>
                </div>

            </Stack>
        </Card>
    );
};

export default NotificationPopup;