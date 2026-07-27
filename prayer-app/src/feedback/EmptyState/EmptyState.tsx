// New file generated from EmptyState.tsx
import {
    Inbox,
} from "lucide-react";

import {
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    emptyStateVariants,
    illustrationVariants,
} from "./EmptyState.styles";

import type {
    EmptyStateProps,
} from "./EmptyState.types";

const EmptyState = ({
    illustration,
    icon,
    heading,
    description,
    actions,
    footer,
    size = "md",
    className,
    ...props
}: EmptyStateProps) => {

    return (
        <Stack
            align="center"
            className={cn(
                emptyStateVariants({
                    size,
                }),
                className
            )}
            {...props}
        >
            {illustration ?? (
                <div
                    className={illustrationVariants({
                        size,
                    })}
                >
                    {icon ?? (
                        <Inbox
                            size={
                                size === "lg"
                                    ? 48
                                    : size === "md"
                                      ? 36
                                      : 24
                            }
                        />
                    )}
                </div>
            )}

            <Stack
                gap="sm"
                align="center"
                className="max-w-md"
            >
                <Typography
                    variant="h3"
                >
                    {heading}
                </Typography>

                {description && (
                    <Typography
                        variant="body"
                        className="text-muted-foreground"
                    >
                        {description}
                    </Typography>
                )}
            </Stack>

            {actions}

            {footer}

        </Stack>
    );
};

export default EmptyState;