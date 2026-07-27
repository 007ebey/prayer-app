import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import {
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { sectionVariants } from "./Section.styles";

import type { SectionProps } from "./Section.types";

const Section = forwardRef<
    HTMLElement,
    SectionProps
>(
    (
        {
            heading: title,
            description: subtitle,
            actions,
            children,
            bordered,
            padded,
            className,
            ...props
        },
        ref
    ) => {
        return (
            <section
                ref={ref}
                className={cn(
                    sectionVariants({
                        bordered,
                        padded,
                    }),
                    className
                )}
                {...props}
            >
                {(title ||
                    subtitle ||
                    actions) && (
                    <Flex
                        justify="between"
                        align="start"
                        className="mb-6"
                    >
                        <Stack gap="xs">
                            {title &&
                                (typeof title ===
                                "string" ? (
                                    <Typography variant="h3">
                                        {title}
                                    </Typography>
                                ) : (
                                    title
                                ))}

                            {subtitle &&
                                (typeof subtitle ===
                                "string" ? (
                                    <Typography
                                        variant="body"
                                        className="text-muted-foreground"
                                    >
                                        {subtitle}
                                    </Typography>
                                ) : (
                                    subtitle
                                ))}
                        </Stack>

                        {actions}
                    </Flex>
                )}

                {children}
            </section>
        );
    }
);

Section.displayName = "Section";

export default Section;