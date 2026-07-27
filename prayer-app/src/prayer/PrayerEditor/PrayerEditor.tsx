// New file generated from PrayerEditor.tsx
import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    prayerEditorContentVariants,
    prayerEditorFooterVariants,
    prayerEditorHeaderVariants,
    prayerEditorVariants,
} from "./PrayerEditor.styles";

import type {
    PrayerEditorProps,
} from "./PrayerEditor.types";

const PrayerEditor = forwardRef<
    HTMLFormElement,
    PrayerEditorProps
>(({
    header,
    toolbar,
    titleField,
    categoryField,
    scriptureField,
    contentField,
    mediaField,
    attachmentsField,
    tagsField,
    footer,
    className,
    children,
    ...props
}, ref) => {

    return (

        <form
            ref={ref}
            className={cn(
                prayerEditorVariants(),
                className
            )}
            {...props}
        >

            {(header || toolbar) && (

                <header
                    className={prayerEditorHeaderVariants()}
                >

                    {header}

                    {toolbar}

                </header>

            )}

            <div
                className={prayerEditorContentVariants()}
            >

                {titleField}

                {categoryField}

                {scriptureField}

                {contentField}

                {mediaField}

                {attachmentsField}

                {tagsField}

                {children}

            </div>

            {footer && (

                <footer
                    className={prayerEditorFooterVariants()}
                >

                    {footer}

                </footer>

            )}

        </form>

    );

});

PrayerEditor.displayName = "PrayerEditor";

export default PrayerEditor;