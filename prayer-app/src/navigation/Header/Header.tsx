import {
    Card,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import type { HeaderProps } from "./Header.types";

const Header = ({
    title,
    subtitle,
    logo,
    startContent,
    endContent,
    sticky = true,
    bordered = true,
}: HeaderProps) => {
    return (
        <header
            className={
                sticky
                    ? "sticky top-0 z-40"
                    : undefined
            }
        >
            <Card className="rounded-none">
                <Flex
                    justify="between"
                    align="center"
                    className="px-6 py-4"
                >
                    <Flex
                        align="center"
                        gap="md"
                    >
                        {logo}

                        {startContent}

                        {(title || subtitle) && (
                            <Stack gap="none">
                                {typeof title ===
                                "string" ? (
                                    <Typography variant="h4">
                                        {title}
                                    </Typography>
                                ) : (
                                    title
                                )}

                                {subtitle &&
                                    (typeof subtitle ===
                                    "string" ? (
                                        <Typography
                                            variant="body-sm"
                                            className="text-muted-foreground"
                                        >
                                            {subtitle}
                                        </Typography>
                                    ) : (
                                        subtitle
                                    ))}
                            </Stack>
                        )}
                    </Flex>

                    {endContent}
                </Flex>

                {bordered && <Divider />}
            </Card>
        </header>
    );
};

export default Header;