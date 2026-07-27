// New file generated from ProgressBar.tsx
import {
    Flex,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    progressIndicatorVariants,
    progressVariants,
} from "./ProgressBar.styles";

import type {
    ProgressBarProps,
} from "./ProgressBar.types";

const ProgressBar = ({
    value,
    max = 100,
    showLabel = false,
    label,
    animated = true,
    striped = false,
    size = "md",
    tone = "primary",
    className,
    ...props
}: ProgressBarProps) => {

    const percentage = Math.min(
        100,
        Math.max(
            0,
            (value / max) * 100
        )
    );

    return (
        <div
            className={cn(
                "w-full",
                className
            )}
            {...props}
        >
            {(label || showLabel) && (
                <Flex
                    justify="between"
                    className="mb-2"
                >
                    <Typography
                        variant="body-sm"
                    >
                        {label}
                    </Typography>

                    {showLabel && (
                        <Typography
                            variant="body-sm"
                        >
                            {Math.round(
                                percentage
                            )}
                            %
                        </Typography>
                    )}
                </Flex>
            )}

            <div
                role="progressbar"
                aria-valuemin={0}
                aria-valuemax={max}
                aria-valuenow={value}
                className={progressVariants({
                    size,
                })}
            >
                <div
                    className={progressIndicatorVariants({
                        tone,
                        animated,
                        striped,
                    })}
                    style={{
                        width: `${percentage}%`,
                    }}
                />
            </div>
        </div>
    );
};

export default ProgressBar;