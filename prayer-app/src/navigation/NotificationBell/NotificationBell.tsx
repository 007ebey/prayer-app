import { Bell } from "lucide-react";

import { Button } from "../../../primitives";

import type { NotificationBellProps } from "./NotificationBell.types";

const NotificationBell = ({
    icon,
    count = 0,
    maxCount = 99,
    showZero = false,
    onClick,
}: NotificationBellProps) => {
    const showBadge =
        showZero || count > 0;

    const badgeText =
        count > maxCount
            ? `${maxCount}+`
            : count;

    return (
        <div className="relative inline-flex">
            <Button
                variant="ghost"
                size="icon"
                onClick={onClick}
                aria-label="Notifications"
            >
                {icon ?? <Bell size={20} />}
            </Button>

            {showBadge && (
                <span
                    className="
                        absolute
                        -right-1
                        -top-1
                        flex
                        min-h-5
                        min-w-5
                        items-center
                        justify-center
                        rounded-full
                        bg-destructive
                        px-1
                        text-[10px]
                        font-semibold
                        text-destructive-foreground
                    "
                >
                    {badgeText}
                </span>
            )}
        </div>
    );
};

export default NotificationBell;