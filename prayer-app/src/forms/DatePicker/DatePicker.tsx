import { useState } from "react";

import { CalendarIcon } from "lucide-react";
import { format } from "date-fns";
import type { Matcher } from "react-day-picker";

import {
    Calendar,
    Input,
    Popover,
} from "../../../primitives";

import type {
    DatePickerProps,
} from "./DatePicker.types";

const DatePicker = ({
    value,
    onChange,
    placeholder = "Select date",
    minDate,
    maxDate,
}: DatePickerProps) => {

    const [open, setOpen] =
        useState(false);

    const disabled: Matcher[] = [];

    if (minDate !== undefined) {

        disabled.push({
            before: minDate,
        });

    }

    if (maxDate !== undefined) {

        disabled.push({
            after: maxDate,
        });

    }

    return (

        <Popover
            open={open}
            onOpenChange={setOpen}
            trigger={

                <Input
                    readOnly
                    value={
                        value
                            ? format(
                                value,
                                "PPP"
                            )
                            : ""
                    }
                    placeholder={
                        placeholder
                    }
                    rightIcon={
                        <CalendarIcon
                            size={18}
                        />
                    }
                />

            }
        >

            <Calendar
                mode="single"
                selected={value}
                onSelect={(date) => {

                    onChange?.(
                        date
                    );

                    setOpen(
                        false
                    );

                }}
                {...(
                    disabled !== undefined
                        ? {
                            disabled,
                        }
                        : {}
                )}
            />

        </Popover>

    );

};

export default DatePicker;