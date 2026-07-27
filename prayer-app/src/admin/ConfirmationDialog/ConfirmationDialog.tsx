// New file generated from ConfirmationDialog.tsx
import {
    Button,
    Card,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { confirmationDialogVariants } from "./ConfirmationDialog.styles";

import type { ConfirmationDialogProps } from "./ConfirmationDialog.types";

const ConfirmationDialog = ({
    open,
    heading,
    description,
    confirmLabel = "Confirm",
    cancelLabel = "Cancel",
    intent = "default",
    loading = false,
    icon,
    onConfirm,
    onCancel,
    className,
    ...props
}: ConfirmationDialogProps) => {
    if (!open) {
        return null;
    }

    const confirmVariant =
        intent === "danger"
            ? "danger"
            : "primary";

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4">

            <Card
                className={cn(
                    confirmationDialogVariants({
                        intent,
                    }),
                    className
                )}
                {...props}
            >
                <Stack
                    gap="lg"
                    className="p-6"
                >
                    {icon && (
                        <Flex justify="center">
                            {icon}
                        </Flex>
                    )}

                    <Stack
                        gap="sm"
                        align="center"
                    >
                        <Typography
                            variant="h3"
                            align="center"
                        >
                            {heading}
                        </Typography>

                        {description && (
                            <Typography
                                variant="body"
                                align="center"
                                className="text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}
                    </Stack>

                    <Divider />

                    <Flex
                        justify="end"
                        gap="sm"
                    >
                        <Button
                            variant="outline"
                            onClick={onCancel}
                            disabled={loading}
                        >
                            {cancelLabel}
                        </Button>

                        <Button
                            variant={confirmVariant}
                            loading={loading}
                            onClick={onConfirm}
                        >
                            {confirmLabel}
                        </Button>
                    </Flex>
                </Stack>
            </Card>

        </div>
    );
};

export default ConfirmationDialog;