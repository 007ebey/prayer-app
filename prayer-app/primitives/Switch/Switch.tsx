import { forwardRef } from "react";

import {
    Stack,
    Typography,
} from "..";

import { cn } from "../../utils/cn";

import {
    switchThumbVariants,
    switchTrackVariants,
    switchVariants,
} from "./Switch.styles";

import type { SwitchProps } from "./Switch.types";

const Switch = forwardRef<
    HTMLInputElement,
    SwitchProps
>(
    (
        {
            id,
            label,
            helperText,
            error,
            checked = false,
            disabled = false,
            onCheckedChange,
            className,
            ...props
        },
        ref
    ) => {
        return (
            <Stack gap="xs">

                <label
                    htmlFor={id}
                    className="flex cursor-pointer items-center justify-between gap-4"
                >
                    {label && (
                        <Typography variant="body">
                            {label}
                        </Typography>
                    )}

                    <span
                        className={cn(
                            switchTrackVariants({
                                checked,
                                disabled,
                            })
                        )}
                    >
                        <input
                            ref={ref}
                            id={id}
                            type="checkbox"
                            checked={checked}
                            disabled={disabled}
                            className={cn(
                                switchVariants(),
                                className
                            )}
                            onChange={(event) =>
                                onCheckedChange?.(
                                    event.target.checked
                                )
                            }
                            {...props}
                        />

                        <span
                            className={cn(
                                switchThumbVariants({
                                    checked,
                                })
                            )}
                        />
                    </span>

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

Switch.displayName = "Switch";

export default Switch;