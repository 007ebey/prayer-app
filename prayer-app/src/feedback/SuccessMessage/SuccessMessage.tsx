// New file generated from SuccessMessage.tsx
import {
    CheckCircle2,
} from "lucide-react";

import {
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    headingVariants,
    iconVariants,
    messageVariants,
    successMessageVariants,
} from "./SuccessMessage.styles";

import type {
    SuccessMessageProps,
} from "./SuccessMessage.types";

const SuccessMessage = ({
    heading,
    message,
    icon,
    actions,
    variant = "inline",
    className,
    ...props
}: SuccessMessageProps) => {

    return (
        <div
            role="status"
            aria-live="polite"
            className={cn(
                successMessageVariants({
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
                    <CheckCircle2
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
                        className={headingVariants()}
                    >
                        {heading}
                    </Typography>
                )}

                <Typography
                    variant="body-sm"
                    className={messageVariants()}
                >
                    {message}
                </Typography>

                {actions}

            </Stack>

        </div>
    );
};

export default SuccessMessage;