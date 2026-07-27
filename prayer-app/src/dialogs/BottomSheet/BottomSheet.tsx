// New file generated from BottomSheet.tsx
import {
    useEffect,
} from "react";

import {
    X,
} from "lucide-react";

import {
    Button,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    backdropVariants,
    bottomSheetVariants,
} from "./BottomSheet.styles";

import type {
    BottomSheetProps,
} from "./BottomSheet.types";

const BottomSheet = ({
    open,
    onClose,
    heading,
    description,
    footer,
    children,
    closeOnBackdrop = true,
    closeOnEscape = true,
    showHandle = true,
    size = "md",
    className,
    ...props
}: BottomSheetProps) => {

    useEffect(() => {

        if (!open || !closeOnEscape) {
            return;
        }

        const handler = (
            event: KeyboardEvent
        ) => {

            if (event.key === "Escape") {
                onClose();
            }

        };

        window.addEventListener(
            "keydown",
            handler
        );

        return () =>
            window.removeEventListener(
                "keydown",
                handler
            );

    }, [
        open,
        closeOnEscape,
        onClose,
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
                className={cn(
                    bottomSheetVariants({
                        size,
                    }),
                    className
                )}
                role="dialog"
                aria-modal="true"
                {...props}
            >
                <Stack
                    className="h-full"
                >

                    {showHandle && (
                        <Flex
                            justify="center"
                            className="py-3"
                        >
                            <div className="h-1.5 w-12 rounded-full bg-muted-foreground/30" />
                        </Flex>
                    )}

                    {(heading ||
                        description) && (
                        <>
                            <Flex
                                justify="between"
                                align="start"
                                className="px-6"
                            >
                                <Stack gap="xs">

                                    {heading && (
                                        <Typography
                                            variant="h3"
                                        >
                                            {heading}
                                        </Typography>
                                    )}

                                    {description && (
                                        <Typography
                                            variant="body-sm"
                                            className="text-muted-foreground"
                                        >
                                            {description}
                                        </Typography>
                                    )}

                                </Stack>

                                <Button
                                    variant="ghost"
                                    size="icon"
                                    onClick={
                                        onClose
                                    }
                                >
                                    <X
                                        size={18}
                                    />
                                </Button>

                            </Flex>

                            <Divider />
                        </>
                    )}

                    <div className="flex-1 overflow-auto px-6 py-4">
                        {children}
                    </div>

                    {footer && (
                        <>
                            <Divider />

                            <div className="p-6">
                                {footer}
                            </div>
                        </>
                    )}

                </Stack>
            </div>
        </>
    );
};

export default BottomSheet;