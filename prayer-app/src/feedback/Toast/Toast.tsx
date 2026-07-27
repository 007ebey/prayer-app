// New file generated from Toast.tsx
import {
    AlertCircle,
    CheckCircle2,
    Info,
    TriangleAlert,
    X,
} from "lucide-react";

import {
    Button,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    iconVariants,
    toastVariants,
} from "./Toast.styles";

import type {
    ToastProps,
} from "./Toast.types";

const icons = {
    info: Info,
    success: CheckCircle2,
    warning: TriangleAlert,
    danger: AlertCircle,
};

const Toast = ({
    heading,
    description,
    icon,
    actions,
    dismissible = true,
    onDismiss,
    tone = "info",
    className,
    ...props
}: ToastProps) => {

    const Icon = icons[tone];

    return (
        <div
            role="status"
            aria-live="polite"
            className={cn(
                toastVariants({
                    tone,
                }),
                className
            )}
            {...props}
        >
            <div
                className={iconVariants({
                    tone,
                })}
            >
                {icon ?? (
                    <Icon size={20} />
                )}
            </div>

            <Stack
                gap="xs"
                className="min-w-0 flex-1"
            >
                {heading && (
                    <Typography
                        variant="body"
                        weight="semibold"
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

                {actions}
            </Stack>

            {dismissible && (
                <Button
                    variant="ghost"
                    size="iconSm"
                    onClick={onDismiss}
                    aria-label="Dismiss notification"
                >
                    <X size={16} />
                </Button>
            )}
        </div>
    );
};

export default Toast;