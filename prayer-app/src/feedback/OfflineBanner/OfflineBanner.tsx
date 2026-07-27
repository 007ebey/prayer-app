// New file generated from OfflineBanner.tsx
import {
    WifiOff,
} from "lucide-react";

import {
    Flex,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    offlineBannerVariants,
} from "./OfflineBanner.styles";

import type {
    OfflineBannerProps,
} from "./OfflineBanner.types";

const OfflineBanner = ({
    offline,
    heading = "You're offline",
    description = "Some features may be unavailable until your connection is restored.",
    icon,
    actions,
    position = "top",
    className,
    ...props
}: OfflineBannerProps) => {

    if (!offline) {
        return null;
    }

    return (
        <div
            role="status"
            aria-live="polite"
            className={cn(
                offlineBannerVariants({
                    position,
                }),
                className
            )}
            {...props}
        >
            <Flex
                align="center"
                justify="between"
                className="mx-auto max-w-7xl gap-4 px-4 py-3"
            >
                <Flex
                    align="center"
                    gap="md"
                    className="min-w-0 flex-1"
                >
                    {icon ?? (
                        <WifiOff
                            size={20}
                            className="shrink-0"
                        />
                    )}

                    <div className="min-w-0">
                        <Typography
                            variant="body"
                            weight="semibold"
                        >
                            {heading}
                        </Typography>

                        <Typography
                            variant="body-sm"
                        >
                            {description}
                        </Typography>
                    </div>
                </Flex>

                {actions}
            </Flex>
        </div>
    );
};

export default OfflineBanner;