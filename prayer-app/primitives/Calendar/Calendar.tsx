import {
    ChevronLeft,
    ChevronRight,
} from "lucide-react";
import {
    DayPicker,
} from "react-day-picker";
import "react-day-picker/dist/style.css";

import { cn } from "../../utils/cn";

import {
    calendarVariants,
} from "./Calendar.styles";

import type {
    CalendarProps,
} from "./Calendar.types";

const Calendar = ({
    className,
    classNames,
    components,
    ...props
}: CalendarProps) => {

    return (

        <DayPicker
            showOutsideDays
            className={cn(
                calendarVariants(),
                className,
            )}
            classNames={{
                months: "flex flex-col gap-4 sm:flex-row",

                month: "space-y-4",

                month_caption:
                    "relative flex items-center justify-center",

                caption_label:
                    "text-sm font-semibold",

                nav: "flex items-center gap-1",

                button_previous:
                    "inline-flex h-8 w-8 items-center justify-center rounded-md hover:bg-accent",

                button_next:
                    "inline-flex h-8 w-8 items-center justify-center rounded-md hover:bg-accent",

                weekdays: "flex",

                weekday:
                    "w-9 text-xs font-medium text-muted-foreground",

                week: "mt-2 flex w-full",

                day_button:
                    "h-9 w-9 rounded-md text-sm hover:bg-accent focus:outline-none focus:ring-2 focus:ring-ring",

                selected:
                    "bg-primary text-primary-foreground",

                today:
                    "border border-primary",

                outside:
                    "text-muted-foreground opacity-50",

                disabled:
                    "opacity-40",

                hidden:
                    "invisible",

                ...classNames,
            }}
            components={{
                Chevron: ({
                    orientation,
                }) =>
                    orientation === "left"
                        ? <ChevronLeft size={16} />
                        : <ChevronRight size={16} />,

                ...components,
            }}
            {...props}
        />

    );

};

export default Calendar;