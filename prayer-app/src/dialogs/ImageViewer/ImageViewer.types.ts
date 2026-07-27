import {
    HTMLAttributes,
    ReactNode,
} from "react";

export interface ImageViewerProps
    extends Omit<
        HTMLAttributes<HTMLDivElement>,
        "title"
    > {

    open: boolean;

    image: string;

    alt?: string;

    heading?: ReactNode;

    description?: ReactNode;

    onClose: () => void;

    onPrevious?: () => void;

    onNext?: () => void;

    footer?: ReactNode;

    showNavigation?: boolean;

    closeOnBackdrop?: boolean;

    closeOnEscape?: boolean;
}