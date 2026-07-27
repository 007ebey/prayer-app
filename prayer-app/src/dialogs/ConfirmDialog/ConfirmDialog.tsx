// New file generated from ConfirmDialog.tsx
import {
    useEffect,
} from "react";

import {
    AlertTriangle,
} from "lucide-react";

import {
    Button,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    backdropVariants,
    dialogVariants,
    iconVariants,
} from "./ConfirmDialog.styles";

import type {
    ConfirmDialogProps,
} from "./ConfirmDialog.types";

const ConfirmDialog = ({
    open,
    onClose,
    onConfirm,
    heading,
    description,
    icon,
    confirmLabel = "Confirm",
    cancelLabel = "Cancel",
    tone = "default",
    loading = false,
    closeOnBackdrop = true,
    closeOnEscape = true,
    className,
    ...props
}: ConfirmDialogProps) => {

    useEffect(() => {

        if (
            !open ||
            !closeOnEscape
        ) {
            return;
        }

        const handler = (
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
                    dialogVariants(),
                    className
                )}
                {...props}
            >
                <Stack
                    gap="lg"
                    className="p-6"
                >

                    <Flex justify="center">

                        <div
                            className={iconVariants({
                                tone,
                            })}
                        >
                            {icon ?? (
                                <AlertTriangle
                                    size={28}
                                />
                            )}
                        </div>

                    </Flex>

                    <Stack
                        gap="sm"
                        align="center"
                    >

                        <Typography
                            variant="h3"
                            className="text-center"
                        >
                            {heading}
                        </Typography>

                        {description && (
                            <Typography
                                variant="body"
                                className="text-center text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}

                    </Stack>

                    <Divider />

                    <Flex
                        gap="md"
                        justify="end"
                    >

                        <Button
                            variant="outline"
                            onClick={onClose}
                            disabled={loading}
                        >
                            {cancelLabel}
                        </Button>

                        <Button
                            variant={
                                tone === "danger"
                                    ? "danger"
                                    : "primary"
                            }
                            onClick={onConfirm}
                            loading={loading}
                        >
                            {confirmLabel}
                        </Button>

                    </Flex>

                </Stack>

            </div>
        </>
    );
};

export default ConfirmDialog;