// New file generated from PrayerAgenda.tsx
import { cn } from "../../../utils/cn";

import {
    prayerAgendaContentVariants,
    prayerAgendaVariants,
} from "./PrayerAgenda.styles";

import type {
    PrayerAgendaProps,
} from "./PrayerAgenda.types";

const PrayerAgenda = ({
    header,
    items,
    currentItem,
    controls,
    footer,
}: PrayerAgendaProps) => {

    return (

        <section
            className={cn(
                prayerAgendaVariants()
            )}
        >

            {header}

            {currentItem}

            <div
                className={cn(
                    prayerAgendaContentVariants()
                )}
            >

                {items}

            </div>

            {controls}

            {footer}

        </section>

    );

};

export default PrayerAgenda;