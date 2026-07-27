// New file generated from WelcomeCard.tsx
import {
    Card,
    Flex,
    Image,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { welcomeCardVariants } from "./WelcomeCard.styles";

import type {
    WelcomeCardProps,
} from "./WelcomeCard.types";

const WelcomeCard = ({
    greeting = "Welcome",
    name,
    message,
    backgroundImage,
    illustration,
    actions,
    footer,
    className,
    ...props
}: WelcomeCardProps) => {

    return (
        <Card
            className={cn(
                welcomeCardVariants(),
                className
            )}
            {...props}
        >

            {backgroundImage && (
                <>
                    <Image
                        src={backgroundImage}
                        alt=""
                        radius="none"
                        className="absolute inset-0 h-full w-full object-cover"
                    />

                    <div className="absolute inset-0 bg-gradient-to-r from-black/70 via-black/40 to-transparent" />
                </>
            )}

            <Flex
                justify="between"
                align="center"
                className="relative z-10 p-8"
            >

                <Stack
                    gap="md"
                    className="max-w-xl"
                >

                    <Stack gap="xs">

                        <Typography
                            variant="body"
                            className={
                                backgroundImage
                                    ? "text-white/80"
                                    : "text-muted-foreground"
                            }
                        >
                            {greeting}
                        </Typography>

                        <Typography
                            variant="display-sm"
                            weight="bold"
                            className={
                                backgroundImage
                                    ? "text-white"
                                    : undefined
                            }
                        >
                            {name}
                        </Typography>

                        {message && (
                            <Typography
                                variant="body"
                                className={
                                    backgroundImage
                                        ? "text-white/90"
                                        : "text-muted-foreground"
                                }
                            >
                                {message}
                            </Typography>
                        )}

                    </Stack>

                    {actions}

                    {footer}

                </Stack>

                {illustration && (
                    <div className="hidden lg:block">
                        {illustration}
                    </div>
                )}

            </Flex>

        </Card>
    );
};

export default WelcomeCard;