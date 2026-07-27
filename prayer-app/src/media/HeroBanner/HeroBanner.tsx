import { forwardRef } from "react";

import {
    BackgroundImage,
    GradientOverlay,
} from "..";

import {
    Container,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { heroBannerVariants } from "./HeroBanner.styles";

import type { HeroBannerProps } from "./HeroBanner.types";

const HeroBanner = forwardRef<
    HTMLElement,
    HeroBannerProps
>(
    (
        {
            backgroundImage,
            heading,
            description,
            badge,
            primaryAction,
            secondaryAction,
            overlay = true,
            alignment,
            fullHeight,
            children,
            className,
            ...props
        },
        ref
    ) => {
        return (
            <section
                ref={ref}
                className={cn(
                    heroBannerVariants({
                        alignment,
                        fullHeight,
                    }),
                    className
                )}
                {...props}
            >
                <BackgroundImage
                    src={backgroundImage}
                    className="absolute inset-0"
                />

                {overlay && (
                    <GradientOverlay />
                )}

                <Container className="relative z-10 flex h-full items-center py-24">

                    <Stack
                        gap="lg"
                        className="max-w-3xl"
                    >

                        {badge}

                        <Typography variant="display">
                            {heading}
                        </Typography>

                        {description && (
                            <Typography
                                variant="body-lg"
                                className="text-muted-foreground"
                            >
                                {description}
                            </Typography>
                        )}

                        {(primaryAction ||
                            secondaryAction) && (
                            <Flex gap="md">

                                {primaryAction}

                                {secondaryAction}

                            </Flex>
                        )}

                        {children}

                    </Stack>

                </Container>

            </section>
        );
    }
);

HeroBanner.displayName =
    "HeroBanner";

export default HeroBanner;