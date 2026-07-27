// New file generated from OTPInput.tsx
import {
    forwardRef,
    useEffect,
    useRef,
    useState,
} from "react";

import { cn } from "../../../utils/cn";

import {
    FormGroup,
} from "../FormGroup";

import {
    otpInputVariants,
} from "./OTPInput.styles";

import type {
    OTPInputProps,
} from "./OTPInput.types";

const OTPInput = forwardRef<
    HTMLDivElement,
    OTPInputProps
>(({
    label,
    helperText,
    error,
    value,
    defaultValue = "",
    onChange,
    length = 6,
    disabled = false,
    autoFocus = false,
    numericOnly = true,
    fullWidth = true,
    className,
    ...props
}, ref) => {

    const controlled =
        value !== undefined;

    const [
        internalValue,
        setInternalValue,
    ] = useState(
        defaultValue.padEnd(
            length,
            " "
        )
    );

    const otp =
        (
            controlled
                ? value
                : internalValue
        )
            .padEnd(
                length,
                " "
            )
            .split("");

    const refs =
        useRef<
            HTMLInputElement[]
        >([]);

    useEffect(() => {

        if (
            autoFocus
        ) {

            refs.current[0]?.focus();

        }

    }, [
        autoFocus,
    ]);

    const update = (
        next: string[],
    ) => {

        const joined =
            next.join("");

        if (
            !controlled
        ) {

            setInternalValue(
                joined
            );

        }

        onChange?.(
            joined.trimEnd()
        );

    };

    return (

        <FormGroup
            ref={ref}
            label={label}
            helperText={helperText}
            error={error}
            fullWidth={fullWidth}
        >

            <div
                className={cn(
                    "flex gap-3",
                    className
                )}
                {...props}
            >

                {otp.map(
                    (
                        character,
                        index,
                    ) => (

                        <input
                            key={
                                index
                            }
                            ref={node => {

                                if (
                                    node
                                ) {

                                    refs.current[
                                        index
                                    ] =
                                        node;

                                }

                            }}
                            className={cn(
                                otpInputVariants({

                                    state:
                                        disabled
                                            ? "disabled"
                                            : error
                                            ? "error"
                                            : "default",

                                })
                            )}
                            value={
                                character.trim()
                            }
                            disabled={
                                disabled
                            }
                            maxLength={
                                1
                            }
                            inputMode={
                                numericOnly
                                    ? "numeric"
                                    : "text"
                            }
                            onChange={e => {

                                let next =
                                    e.target.value;

                                if (
                                    numericOnly
                                ) {

                                    next =
                                        next.replace(
                                            /\D/g,
                                            ""
                                        );

                                }

                                const copy =
                                    [
                                        ...otp,
                                    ];

                                copy[
                                    index
                                ] =
                                    next;

                                update(
                                    copy
                                );

                                if (
                                    next &&
                                    index <
                                        length -
                                            1
                                ) {

                                    refs.current[
                                        index +
                                            1
                                    ]?.focus();

                                }

                            }}
                            onKeyDown={e => {

                                if (
                                    e.key ===
                                        "Backspace" &&
                                    !otp[
                                        index
                                    ] &&
                                    index >
                                        0
                                ) {

                                    refs.current[
                                        index -
                                            1
                                    ]?.focus();

                                }

                            }}
                            onPaste={e => {

                                e.preventDefault();

                                let pasted =
                                    e.clipboardData.getData(
                                        "text"
                                    );

                                if (
                                    numericOnly
                                ) {

                                    pasted =
                                        pasted.replace(
                                            /\D/g,
                                            ""
                                        );

                                }

                                const chars =
                                    pasted
                                        .slice(
                                            0,
                                            length
                                        )
                                        .split("");

                                while (
                                    chars.length <
                                    length
                                ) {

                                    chars.push(
                                        " "
                                    );

                                }

                                update(
                                    chars
                                );

                            }}
                        />

                    )
                )}

            </div>

        </FormGroup>

    );

});

OTPInput.displayName =
    "OTPInput";

export default OTPInput;