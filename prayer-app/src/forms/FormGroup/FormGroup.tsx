// New file generated from FormGroup.tsx
import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    formGroupVariants,
} from "./FormGroup.styles";

import type {
    FormGroupProps,
} from "./FormGroup.types";

const FormGroup = forwardRef<
    HTMLDivElement,
    FormGroupProps
>(({
    label,
    helperText,
    error,
    required = false,
    optional = false,
    disabled = false,
    fullWidth = true,
    children,
    className,
    ...props
}, ref) => {

    return (

        <div
            ref={ref}
            className={cn(
                formGroupVariants({
                    fullWidth,
                }),
                className
            )}
            {...props}
        >

            {label && (

                <label
                    className={cn(
                        "text-sm font-medium",
                        disabled &&
                            "opacity-60"
                    )}
                >

                    {label}

                    {required && (

                        <span className="ml-1 text-destructive">

                            *

                        </span>

                    )}

                    {!required &&
                        optional && (

                            <span className="ml-2 text-xs text-muted-foreground">

                                (Optional)

                            </span>

                        )}

                </label>

            )}

            {children}

            {error ? (

                <p
                    className="text-sm text-destructive"
                    role="alert"
                >

                    {error}

                </p>

            ) : helperText ? (

                <p className="text-sm text-muted-foreground">

                    {helperText}

                </p>

            ) : null}

        </div>

    );

});

FormGroup.displayName =
    "FormGroup";

export default FormGroup;