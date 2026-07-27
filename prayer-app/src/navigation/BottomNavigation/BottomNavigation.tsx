import {
    Button,
    Card,
    Typography,
} from "../../../primitives";

import type {
    BottomNavigationItem,
    BottomNavigationProps,
} from "./BottomNavigation.types";

const BottomNavigation = ({
    items,
    activeItem,
    onItemClick,
}: BottomNavigationProps) => {
    return (
        <Card className="fixed bottom-0 left-0 right-0 rounded-none border-t px-2 py-2 shadow-lg">
            <nav className="flex items-center justify-around">
                {items.map((item) => {
                    const active =
                        item.id === activeItem;

                    return (
                        <Button
                            key={item.id}
                            variant="ghost"
                            onClick={() =>
                                onItemClick?.(
                                    item
                                )
                            }
                            className={`relative flex h-auto flex-1 flex-col items-center gap-1 py-2 ${
                                active
                                    ? "text-primary"
                                    : ""
                            }`}
                        >
                            <div className="relative">
                                {item.icon}

                                {item.badge &&
                                    item.badge > 0 && (
                                        <span className="absolute -right-2 -top-2 flex h-5 min-w-5 items-center justify-center rounded-full bg-destructive px-1 text-[10px] font-medium text-destructive-foreground">
                                            {item.badge}
                                        </span>
                                    )}
                            </div>

                            <Typography
                                variant="caption"
                                weight={
                                    active
                                        ? "semibold"
                                        : "normal"
                                }
                            >
                                {item.label}
                            </Typography>
                        </Button>
                    );
                })}
            </nav>
        </Card>
    );
};

export default BottomNavigation;