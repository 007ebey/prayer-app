import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { InputProps } from "./Input.types";

import { inputVariants } from "./Input.styles";

const Input = forwardRef<
    HTMLInputElement,
    InputProps
>(

(
    {
        label,
        helperText,
        error,
        variant,
        leftIcon,
        rightIcon,
        inputSize = "md",
        fullWidth = true,
        className,
        id,
        ...props
    },

    ref

) => {

    return (

        <div
            className={cn(
                fullWidth && "w-full"
            )}
        >

            {label && (

                <label

                    htmlFor={id}

                    className="mb-2 block text-sm font-medium text-slate-700"

                >

                    {label}

                </label>

            )}

            <div className="relative">

                {leftIcon && (

                    <div
                        className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400"
                    >

                        {leftIcon}

                    </div>

                )}

                <input

                    ref={ref}

                    id={id}

                    className={cn(

                        inputVariants({

                            variant,

                            size: inputSize,

                            hasError: !!error

                        }),

                        leftIcon && "pl-10",

                        rightIcon && "pr-10",

                        className

                    )}

                    {...props}

                />

                {rightIcon && (

                    <div
                        className="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400"
                    >

                        {rightIcon}

                    </div>

                )}

            </div>

            {error ? (

                <p className="mt-2 text-sm text-red-600">

                    {error}

                </p>

            ) : helperText ? (

                <p className="mt-2 text-sm text-slate-500">

                    {helperText}

                </p>

            ) : null}

        </div>

    );

});

Input.displayName = "Input";

export default Input;