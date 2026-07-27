import { forwardRef } from "react";

import {
    Stack,
    Typography,
} from "..";

import { cn } from "../../utils/cn";

import { textareaVariants } from "./Textarea.styles";

import type { TextareaProps } from "./Textarea.types";

const Textarea = forwardRef<
    HTMLTextAreaElement,
    TextareaProps
>(
    (
        {
            label,
            helperText,
            error,
            variant,
            resize,
            fullWidth,
            className,
            id,
            ...props
        },
        ref
    ) => {
        return (
            <Stack
                gap="xs"
                className={
                    fullWidth
                        ? "w-full"
                        : undefined
                }
            >
                {label && (
                    <label
                        htmlFor={id}
                    >
                        <Typography
                            variant="body-sm"
                            weight="medium"
                        >
                            {label}
                        </Typography>
                    </label>
                )}

                <textarea
                    ref={ref}
                    id={id}
                    className={cn(
                        textareaVariants({
                            variant,
                            resize,
                            error: !!error,
                            fullWidth,
                        }),
                        className
                    )}
                    {...props}
                />

                {error ? (
                    <Typography
                        variant="caption"
                        className="text-destructive"
                    >
                        {error}
                    </Typography>
                ) : (
                    helperText && (
                        <Typography
                            variant="caption"
                            className="text-muted-foreground"
                        >
                            {helperText}
                        </Typography>
                    )
                )}
            </Stack>
        );
    }
);

Textarea.displayName = "Textarea";

export default Textarea;