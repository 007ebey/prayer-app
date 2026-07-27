// New file generated from Checkbox.tsx
import {
    forwardRef,
    useEffect,
    useRef,
} from "react";

import {
    Check,
    Minus,
} from "lucide-react";

import {
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    checkboxContainerVariants,
    checkboxVariants,
} from "./Checkbox.styles";

import type {
    CheckboxProps,
} from "./Checkbox.types";

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
        checked,
        disabled,
        fullWidth = true,
        className,
        id,
        ...props
    },
    ref
) => {

    const inputRef =
        useRef<HTMLInputElement>(null);

    useEffect(() => {

        if (inputRef.current) {

            inputRef.current.indeterminate =
                indeterminate;

        }

    }, [indeterminate]);

    const mergedRef = (
        node: HTMLInputElement
    ) => {

        inputRef.current = node;

        if (
            typeof ref === "function"
        ) {
            ref(node);
        }
        else if (ref) {
            ref.current = node;
        }

    };

    return (

        <label
            htmlFor={id}
            className={cn(
                checkboxContainerVariants({
                    fullWidth,
                }),
                className
            )}
        >

            <span className="relative">

                <input
                    ref={mergedRef}
                    id={id}
                    type="checkbox"
                    checked={checked}
                    disabled={disabled}
                    className="peer sr-only"
                    {...props}
                />

                <span
                    className={cn(
                        checkboxVariants({
                            checked:
                                Boolean(checked) ||
                                indeterminate,
                            error:
                                Boolean(error),
                        })
                    )}
                >

                    {indeterminate ? (

                        <Minus
                            size={14}
                        />

                    ) : checked ? (

                        <Check
                            size={14}
                        />

                    ) : null}

                </span>

            </span>

            <Stack
                gap="xs"
                className="min-w-0"
            >

                {label && (

                    <Typography
                        variant="body"
                    >
                        {label}
                    </Typography>

                )}

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

        </label>

    );

});

Checkbox.displayName =
    "Checkbox";

export default Checkbox;