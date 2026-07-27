import {
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import type { PageTitleProps } from "./PageTitle.types";

const PageTitle = ({
    title,
    subtitle,
    actions,
    centered = false,
}: PageTitleProps) => {
    return (
        <Flex
            justify="between"
            align="center"
            className={
                centered
                    ? "flex-col gap-4 text-center"
                    : undefined
            }
        >
            <Stack gap="xs">
                {typeof title === "string" ? (
                    <Typography variant="h2">
                        {title}
                    </Typography>
                ) : (
                    title
                )}

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

            {!centered && actions}
        </Flex>
    );
};

export default PageTitle;