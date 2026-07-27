// New file generated from PasswordInput.tsx
import {
    forwardRef,
    useState,
} from "react";

import {
    Eye,
    EyeOff,
} from "lucide-react";

import { cn } from "../../../utils/cn";

import {
    FormGroup,
} from "../FormGroup";

import {
    passwordInputVariants,
} from "./PasswordInput.styles";

import type {
    PasswordInputProps,
} from "./PasswordInput.types";

const PasswordInput = forwardRef<
    HTMLInputElement,
    PasswordInputProps
>(({
    label,
    helperText,
    error,
    leftIcon,
    showToggle = true,
    variant,
    inputSize = "md",
    fullWidth = true,
    className,
    disabled = false,
    ...props
}, ref) => {

    const [visible, setVisible] =
        useState(false);

    return (

        <FormGroup
            label={label}
            helperText={helperText}
            error={error}
            disabled={disabled}
            fullWidth={fullWidth}
        >

            <div
                className={cn(
                    passwordInputVariants({

                        variant,

                        inputSize,

                        error:
                            error != null,

                        disabled,

                        fullWidth,

                    }),
                    className
                )}
            >

                {leftIcon && (

                    <span className="mr-2 flex items-center text-muted-foreground">

                        {leftIcon}

                    </span>

                )}

                <input
                    {...props}
                    ref={ref}
                    disabled={disabled}
                    type={
                        visible
                            ? "text"
                            : "password"
                    }
                    className="
                        flex-1
                        bg-transparent
                        outline-none
                        placeholder:text-muted-foreground
                    "
                />

                {showToggle && (

                    <button
                        type="button"
                        tabIndex={-1}
                        disabled={disabled}
                        onClick={() =>
                            setVisible(
                                previous =>
                                    !previous
                            )
                        }
                        className="
                            ml-2
                            flex
                            items-center
                            text-muted-foreground
                            hover:text-foreground
                            disabled:pointer-events-none
                        "
                        aria-label={
                            visible
                                ? "Hide password"
                                : "Show password"
                        }
                    >

                        {visible
                            ? (
                                <EyeOff
                                    size={18}
                                />
                            )
                            : (
                                <Eye
                                    size={18}
                                />
                            )}

                    </button>

                )}

            </div>

        </FormGroup>

    );

});

PasswordInput.displayName =
    "PasswordInput";

export default PasswordInput;