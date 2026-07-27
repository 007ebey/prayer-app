import { forwardRef } from "react";

import {
    Stack,
    Typography,
} from "..";

import { cn } from "../../utils/cn";

import { radioVariants } from "./Radio.styles";

import type { RadioProps } from "./Radio.types";

const Radio = forwardRef<
    HTMLInputElement,
    RadioProps
>(
    (
        {
            id,
            label,
            helperText,
            error,
            className,
            ...props
        },
        ref
    ) => {
        return (
            <Stack gap="xs">

                <label
                    htmlFor={id}
                    className="flex cursor-pointer items-center gap-3"
                >

                    <input
                        ref={ref}
                        id={id}
                        type="radio"
                        className={cn(
                            radioVariants({
                                error: !!error,
                            }),
                            className
                        )}
                        {...props}
                    />

                    {label && (
                        <Typography
                            variant="body-sm"
                        >
                            {label}
                        </Typography>
                    )}

                </label>

                {helperText && !error && (
                    <Typography
                        variant="caption"
                        className="text-muted-foreground"
                    >
                        {helperText}
                    </Typography>
                )}

                {error && (
                    <Typography
                        variant="caption"
                        className="text-destructive"
                    >
                        {error}
                    </Typography>
                )}

            </Stack>
        );
    }
);

Radio.displayName = "Radio";

export default Radio;