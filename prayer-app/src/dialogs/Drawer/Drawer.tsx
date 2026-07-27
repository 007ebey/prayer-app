// New file generated from Drawer.tsx
import {
    useEffect,
} from "react";

import {
    X,
} from "lucide-react";

import {
    Button,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    backdropVariants,
    drawerVariants,
} from "./Drawer.styles";

import type {
    DrawerProps,
} from "./Drawer.types";

const Drawer = ({
    open,
    onClose,
    heading,
    description,
    footer,
    children,
    side = "right",
    size = "md",
    closeOnBackdrop = true,
    closeOnEscape = true,
    className,
    ...props
}: DrawerProps) => {

    useEffect(() => {

        if (!open || !closeOnEscape) {
            return;
        }

        const handler = (
            event: KeyboardEvent
        ) => {

            if (event.key === "Escape") {
                onClose();
            }

        };

        window.addEventListener(
            "keydown",
            handler
        );

        return () =>
            window.removeEventListener(
                "keydown",
                handler
            );

    }, [
        open,
        closeOnEscape,
        onClose,
    ]);

    if (!open) {
        return null;
    }

    return (
        <>
            <div
                className={backdropVariants()}
                onClick={
                    closeOnBackdrop
                        ? onClose
                        : undefined
                }
            />

            <div
                role="dialog"
                aria-modal="true"
                className={cn(
                    drawerVariants({
                        side,
                        size,
                    }),
                    className
                )}
                {...props}
            >
                {(heading ||
                    description) && (
                    <>
                        <Flex
                            justify="between"
                            align="start"
                            className="p-6"
                        >
                            <Stack gap="xs">

                                {heading && (
                                    <Typography variant="h3">
                                        {heading}
                                    </Typography>
                                )}

                                {description && (
                                    <Typography
                                        variant="body-sm"
                                        className="text-muted-foreground"
                                    >
                                        {description}
                                    </Typography>
                                )}

                            </Stack>

                            <Button
                                variant="ghost"
                                size="icon"
                                onClick={onClose}
                            >
                                <X size={18} />
                            </Button>

                        </Flex>

                        <Divider />
                    </>
                )}

                <div className="flex-1 overflow-auto p-6">
                    {children}
                </div>

                {footer && (
                    <>
                        <Divider />

                        <div className="p-6">
                            {footer}
                        </div>
                    </>
                )}
            </div>
        </>
    );
};

export default Drawer;