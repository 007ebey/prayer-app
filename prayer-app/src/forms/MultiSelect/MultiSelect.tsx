// New file generated from MultiSelect.tsx
import {
    useMemo,
    useState,
} from "react";

import {
    Check,
    ChevronDown,
    X,
} from "lucide-react";

import {
    Badge,
    Checkbox,
    FormGroup,
    Input,
    Popover,
    Stack,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    multiSelectVariants,
} from "./MultiSelect.styles";

import type {
    MultiSelectOption,
    MultiSelectProps,
} from "./MultiSelect.types";

const MultiSelect = ({
    label,
    helperText,
    error,
    placeholder = "Select...",
    options,
    value = [],
    onChange,
    searchable = true,
    clearable = true,
    disabled = false,
    fullWidth = true,
}: MultiSelectProps) => {

    const [search, setSearch] =
        useState("");

    const filtered =
        useMemo(
            () =>
                options.filter(
                    option =>
                        option.label
                            ?.toString()
                            .toLowerCase()
                            .includes(
                                search.toLowerCase()
                            )
                ),
            [
                options,
                search,
            ]
        );

    const toggle = (
        option: MultiSelectOption,
    ) => {

        if (
            option.disabled
        ) {

            return;

        }

        if (
            value.includes(
                option.value
            )
        ) {

            onChange?.(
                value.filter(
                    v =>
                        v !==
                        option.value
                )
            );

            return;

        }

        onChange?.([
            ...value,
            option.value,
        ]);

    };

    return (

        <FormGroup
            label={label}
            helperText={helperText}
            error={error}
            fullWidth={fullWidth}
        >

            <Popover
                trigger={

                    <div
                        className={cn(
                            multiSelectVariants()
                        )}
                    >

                        {value.length ===
                        0 ? (

                            <span className="text-muted-foreground">

                                {
                                    placeholder
                                }

                            </span>

                        ) : (

                            value.map(
                                selected => {

                                    const option =
                                        options.find(
                                            o =>
                                                o.value ===
                                                selected
                                        );

                                    if (
                                        !option
                                    ) {

                                        return null;

                                    }

                                    return (

                                        <Badge
                                            key={
                                                selected
                                            }
                                        >

                                            {
                                                option.label
                                            }

                                        </Badge>

                                    );

                                }
                            )

                        )}

                        <ChevronDown
                            size={18}
                            className="ml-auto"
                        />

                    </div>

                }
            >

                <Stack
                    gap="sm"
                    className="w-72"
                >

                    {searchable && (

                        <Input
                            placeholder="Search..."
                            value={
                                search
                            }
                            onChange={
                                e =>
                                    setSearch(
                                        e.target
                                            .value
                                    )
                            }
                        />

                    )}

                    <div className="max-h-64 overflow-y-auto">

                        {filtered.map(
                            option => (

                                <button
                                    key={
                                        option.value
                                    }
                                    type="button"
                                    className="flex w-full items-center gap-3 rounded-md px-2 py-2 hover:bg-accent"
                                    onClick={() =>
                                        toggle(
                                            option
                                        )
                                    }
                                >

                                    <Checkbox
                                        checked={value.includes(
                                            option.value
                                        )}
                                        readOnly
                                    />

                                    <span className="flex-1 text-left">

                                        {
                                            option.label
                                        }

                                    </span>

                                    {value.includes(
                                        option.value
                                    ) && (

                                        <Check
                                            size={
                                                16
                                            }
                                        />

                                    )}

                                </button>

                            )
                        )}

                    </div>

                    {clearable &&
                        value.length >
                            0 && (

                            <button
                                type="button"
                                className="flex items-center justify-center gap-2 rounded-md border py-2 text-sm hover:bg-accent"
                                onClick={() =>
                                    onChange?.(
                                        []
                                    )
                                }
                            >

                                <X
                                    size={16}
                                />

                                Clear

                            </button>

                        )}

                </Stack>

            </Popover>

        </FormGroup>

    );

};

export default MultiSelect;