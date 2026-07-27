// New file generated from Alert.tsx
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

import { cn } from "../../utils/cn";

import {
    alertVariants,
} from "./Alert.styles";

import type {
    AlertProps,
} from "./Alert.types";

const defaultIcons = {
    info: Info,
    success: CheckCircle2,
    warning: TriangleAlert,
    danger: AlertCircle,
};

const Alert = ({
    heading,
    description,
    icon,
    actions,
    dismissible = false,
    onDismiss,
    variant = "soft",
    tone = "info",
    className,
    ...props
}: AlertProps) => {

    const Icon =
        defaultIcons[tone];

    return (
        <div
            role="alert"
            className={cn(
                alertVariants({
                    variant,
                    tone,
                }),
                className
            )}
            {...props}
        >
            <div className="mt-0.5 shrink-0">
                {icon ?? (
                    <Icon
                        size={22}
                    />
                )}
            </div>

            <Stack
                gap="xs"
                className="min-w-0 flex-1"
            >
                {heading && (
                    <Typography
                        variant="title"
                        weight="semibold"
                    >
                        {heading}
                    </Typography>
                )}

                {description && (
                    <Typography
                        variant="body-sm"
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
                    aria-label="Dismiss alert"
                >
                    <X
                        size={16}
                    />
                </Button>
            )}
        </div>
    );
};

export default Alert;