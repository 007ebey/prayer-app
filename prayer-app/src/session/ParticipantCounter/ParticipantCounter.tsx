// New file generated from ParticipantCounter.tsx
import { forwardRef } from "react";

import { Users } from "lucide-react";

import { cn } from "../../../utils/cn";

import {
    participantCounterContentVariants,
    participantCounterIconVariants,
    participantCounterLabelVariants,
    participantCounterSecondaryVariants,
    participantCounterValueVariants,
    participantCounterVariants,
} from "./ParticipantCounter.styles";

import type {
    ParticipantCounterProps,
} from "./ParticipantCounter.types";

const ParticipantCounter = forwardRef<
    HTMLDivElement,
    ParticipantCounterProps
>(({
    icon,
    label,
    value,
    secondaryValue,
    className,
    ...props
}, ref) => {

    return (

        <div
            ref={ref}
            className={cn(
                participantCounterVariants(),
                className
            )}
            {...props}
        >

            <div
                className={participantCounterIconVariants()}
            >

                {icon ?? (
                    <Users className="size-5" />
                )}

            </div>

            <div
                className={participantCounterContentVariants()}
            >

                <span
                    className={participantCounterValueVariants()}
                >

                    {value}

                </span>

                {label && (

                    <span
                        className={participantCounterLabelVariants()}
                    >

                        {label}

                    </span>

                )}

                {secondaryValue && (

                    <span
                        className={participantCounterSecondaryVariants()}
                    >

                        {secondaryValue}

                    </span>

                )}

            </div>

        </div>

    );

});

ParticipantCounter.displayName =
    "ParticipantCounter";

export default ParticipantCounter;