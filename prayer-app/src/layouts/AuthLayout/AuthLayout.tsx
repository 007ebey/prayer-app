import {
    Card,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import type { AuthLayoutProps } from "./AuthLayout.types";

const AuthLayout = ({
    children,
    logo,
    title,
    subtitle,
    illustration,
    footer,
    reverse = false,
}: AuthLayoutProps) => {
    return (
        <div className="min-h-screen bg-background text-foreground">
            <Flex
                className={`min-h-screen ${
                    reverse ? "flex-row-reverse" : ""
                }`}
            >
                {illustration && (
                    <div className="hidden flex-1 items-center justify-center bg-muted p-12 lg:flex">
                        {illustration}
                    </div>
                )}

                <div className="flex flex-1 items-center justify-center p-6">
                    <Card className="w-full max-w-md p-8">
                        <Stack gap="xl">
                            {(logo || title || subtitle) && (
                                <Stack
                                    gap="sm"
                                    align="center"
                                >
                                    {logo}

                                    {title &&
                                        (typeof title === "string" ? (
                                            <Typography
                                                variant="h2"
                                                align="center"
                                            >
                                                {title}
                                            </Typography>
                                        ) : (
                                            title
                                        ))}

                                    {subtitle &&
                                        (typeof subtitle === "string" ? (
                                            <Typography
                                                variant="body"
                                                align="center"
                                                className="text-muted-foreground"
                                            >
                                                {subtitle}
                                            </Typography>
                                        ) : (
                                            subtitle
                                        ))}
                                </Stack>
                            )}

                            {children}

                            {footer}
                        </Stack>
                    </Card>
                </div>
            </Flex>
        </div>
    );
};

export default AuthLayout;