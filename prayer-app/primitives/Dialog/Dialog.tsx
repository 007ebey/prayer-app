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
} from "../../primitives";

import { cn } from "../../utils/cn";

import {
    backdropVariants,
    dialogVariants,
} from "./Dialog.styles";

import type {
    DialogProps,
} from "./Dialog.types";

const Dialog = ({
    open,
    onClose,
    heading,
    description,
    children,
    footer,
    size = "md",
    showCloseButton = true,
    closeOnBackdrop = true,
    closeOnEscape = true,
    className,
    ...props
}: DialogProps) => {

    useEffect(() => {

        if (
            !open ||
            !closeOnEscape
        ) {
            return;
        }

        const handleKeyDown = (
            event: KeyboardEvent
        ) => {

            if (
                event.key === "Escape"
            ) {
                onClose();
            }

        };

        window.addEventListener(
            "keydown",
            handleKeyDown
        );

        return () =>
            window.removeEventListener(
                "keydown",
                handleKeyDown
            );

    }, [
        open,
        closeOnEscape,
        onClose,
    ]);

    useEffect(() => {

        if (!open) {
            return;
        }

        const previousOverflow =
            document.body.style.overflow;

        document.body.style.overflow =
            "hidden";

        return () => {
            document.body.style.overflow =
                previousOverflow;
        };

    }, [open]);

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
                    dialogVariants({
                        size,
                    }),
                    className
                )}
                {...props}
            >
                {(heading ||
                    description ||
                    showCloseButton) && (
                    <>
                        <Flex
                            justify="between"
                            align="start"
                            className="p-6"
                        >
                            <Stack gap="xs">

                                {heading && (
                                    <Typography
                                        variant="h3"
                                    >
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

                            {showCloseButton && (
                                <Button
                                    variant="ghost"
                                    size="icon"
                                    onClick={onClose}
                                >
                                    <X
                                        size={18}
                                    />
                                </Button>
                            )}

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

export default Dialog;