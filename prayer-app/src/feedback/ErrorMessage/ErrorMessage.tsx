// New file generated from ErrorMessage.tsx
import {
    AlertCircle,
} from "lucide-react";

import {
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    errorMessageVariants,
    iconVariants,
} from "./ErrorMessage.styles";

import type {
    ErrorMessageProps,
} from "./ErrorMessage.types";

const ErrorMessage = ({
    heading,
    message,
    icon,
    actions,
    variant = "inline",
    className,
    ...props
}: ErrorMessageProps) => {

    return (
        <div
            role="alert"
            className={cn(
                errorMessageVariants({
                    variant,
                }),
                className
            )}
            {...props}
        >
            <div
                className={iconVariants()}
            >
                {icon ?? (
                    <AlertCircle
                        size={18}
                    />
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
                        className="text-destructive"
                    >
                        {heading}
                    </Typography>
                )}

                <Typography
                    variant="body-sm"
                    className="text-destructive"
                >
                    {message}
                </Typography>

                {actions}

            </Stack>

        </div>
    );
};

export default ErrorMessage;