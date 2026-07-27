import {
    forwardRef,
    useEffect,
    useRef,
} from "react";

import { cn } from "../../utils/cn";

import {
    Stack,
    Typography,
} from "..";

import { checkboxVariants } from "./Checkbox.styles";

import type { CheckboxProps } from "./Checkbox.types";

const Checkbox = forwardRef<
    HTMLInputElement,
    CheckboxProps
>(
    (
        {
            label,
            helperText,
            error,
            indeterminate = false,
            className,
            id,
            ...props
        },
        ref
    ) => {
        const internalRef =
            useRef<HTMLInputElement>(null);

        useEffect(() => {
            if (internalRef.current) {
                internalRef.current.indeterminate =
                    indeterminate;
            }
        }, [indeterminate]);

        return (
            <Stack gap="xs">
                <label
                    htmlFor={id}
                    className="flex cursor-pointer items-center gap-3"
                >
                    <input
                        id={id}
                        ref={(node) => {
                            internalRef.current =
                                node;

                            if (
                                typeof ref ===
                                "function"
                            ) {
                                ref(node);
                            } else if (
                                ref
                            ) {
                                ref.current =
                                    node;
                            }
                        }}
                        type="checkbox"
                        className={cn(
                            checkboxVariants({
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

Checkbox.displayName = "Checkbox";

export default Checkbox;