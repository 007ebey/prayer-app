// New file generated from ImageViewer.tsx
import {
    useEffect,
} from "react";

import {
    ChevronLeft,
    ChevronRight,
    X,
} from "lucide-react";

import {
    Button,
    Flex,
    Image,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    backdropVariants,
    imageVariants,
    viewerVariants,
} from "./ImageViewer.styles";

import type {
    ImageViewerProps,
} from "./ImageViewer.types";

const ImageViewer = ({
    open,
    image,
    alt,
    heading,
    description,
    onClose,
    onPrevious,
    onNext,
    footer,
    showNavigation = true,
    closeOnBackdrop = true,
    closeOnEscape = true,
    className,
    ...props
}: ImageViewerProps) => {

    useEffect(() => {

        if (!open) {
            return;
        }

        const handler = (
            event: KeyboardEvent
        ) => {

            switch (event.key) {

                case "Escape":

                    if (closeOnEscape) {
                        onClose();
                    }

                    break;

                case "ArrowLeft":

                    onPrevious?.();

                    break;

                case "ArrowRight":

                    onNext?.();

                    break;

            }

        };

        window.addEventListener(
            "keydown",
            handler
        );

        return () =>
            window.removeEventListener(
                "keydown",
                handler);

    }, [
        open,
        closeOnEscape,
        onClose,
        onPrevious,
        onNext,
    ]);

    if (!open) {
        return null;
    }

    return (
        <>
            <div
                className={backdropVariants()}
                onClick={
                    closeOnBackdrop
                        ? onClose
                        : undefined
                }
            />

            <div
                role="dialog"
                aria-modal="true"
                className={cn(
                    viewerVariants(),
                    className
                )}
                {...props}
            >

                <Button
                    variant="ghost"
                    size="icon"
                    className="absolute right-6 top-6 text-white"
                    onClick={onClose}
                >
                    <X />
                </Button>

                {showNavigation &&
                    onPrevious && (
                        <Button
                            variant="ghost"
                            size="icon"
                            className="absolute left-6 text-white"
                            onClick={onPrevious}
                        >
                            <ChevronLeft />
                        </Button>
                    )}

                <Stack
                    gap="md"
                    align="center"
                >

                    <Image
                        src={image}
                        alt={alt ?? ""}
                        fit="contain"
                        className={imageVariants()}
                    />

                    {(heading ||
                        description) && (
                            <Stack
                                gap="xs"
                                align="center"
                            >

                                {heading && (
                                    <Typography
                                        variant="h4"
                                        className="text-white"
                                    >
                                        {heading}
                                    </Typography>
                                )}

                                {description && (
                                    <Typography
                                        variant="body-sm"
                                        className="text-white/80"
                                    >
                                        {description}
                                    </Typography>
                                )}

                            </Stack>
                        )}

                    {footer}

                </Stack>

                {showNavigation &&
                    onNext && (
                        <Button
                            variant="ghost"
                            size="icon"
                            className="absolute right-6 text-white"
                            onClick={onNext}
                        >
                            <ChevronRight />
                        </Button>
                    )}

            </div>
        </>
    );
};

export default ImageViewer;