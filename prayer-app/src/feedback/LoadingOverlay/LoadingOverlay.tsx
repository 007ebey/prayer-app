// New file generated from LoadingOverlay.tsx
import {
    Loader2,
} from "lucide-react";

import {
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    containerVariants,
    overlayVariants,
    spinnerVariants,
} from "./LoadingOverlay.styles";

import type {
    LoadingOverlayProps,
} from "./LoadingOverlay.types";

const LoadingOverlay = ({
    loading,
    message = "Loading...",
    spinner,
    backdrop = true,
    blur = true,
    fullScreen = false,
    size = "md",
    className,
    ...props
}: LoadingOverlayProps) => {

    if (!loading) {
        return null;
    }

    return (
        <div
            role="status"
            aria-live="polite"
            aria-busy="true"
            className={cn(
                overlayVariants({
                    backdrop,
                    blur,
                    fullScreen,
                }),
                className
            )}
            {...props}
        >
            <Stack
                align="center"
                gap="md"
                className={containerVariants({
                    size,
                })}
            >
                {spinner ?? (
                    <Loader2
                        className={spinnerVariants({
                            size,
                        })}
                    />
                )}

                {message && (
                    <Typography
                        variant="body"
                    >
                        {message}
                    </Typography>
                )}
            </Stack>
        </div>
    );
};

export default LoadingOverlay;