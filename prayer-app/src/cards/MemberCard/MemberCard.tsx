// New file generated from MemberCard.tsx
import {
    Calendar,
    MapPin,
} from "lucide-react";

import {
    Avatar,
    Badge,
    Card,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { memberCardVariants } from "./MemberCard.styles";

import type {
    MemberCardProps,
} from "./MemberCard.types";

const roleColors = {
    Member: "primary",
    Host: "warning",
    Admin: "danger",
    "Super Admin": "success",
} as const;

const MemberCard = ({
    name,
    avatar,
    status,
    role,
    location,
    bio,
    joinedDate,
    badges,
    actions,
    className,
    ...props
}: MemberCardProps) => {

    return (
        <Card
            className={cn(
                memberCardVariants(),
                className
            )}
            {...props}
        >
            <Stack
                gap="lg"
                className="p-6"
            >

                <Flex
                    justify="between"
                    align="start"
                >

                    <Flex
                        gap="md"
                        align="center"
                    >

                        <Avatar
                            size="lg"
                            {...(
                                typeof name === "string"
                                    ? { name }
                                    : {}
                            )}
                            {...(
                                avatar
                                    ? { src: avatar }
                                    : {}
                            )}
                            {...(
                                status
                                    ? { status }
                                    : {}
                            )}
                        />

                        <Stack gap="xs">

                            <Typography
                                variant="h4"
                            >
                                {name}
                            </Typography>

                            {role && (
                                <Badge
                                    variant="soft"
                                    color={
                                        typeof role ===
                                            "string" &&
                                            role in
                                            roleColors
                                            ? roleColors[
                                            role as keyof typeof roleColors
                                            ]
                                            : "neutral"
                                    }
                                    size="sm"
                                >
                                    {role}
                                </Badge>
                            )}

                        </Stack>

                    </Flex>

                    {actions}

                </Flex>

                {bio && (
                    <Typography
                        variant="body-sm"
                        className="text-muted-foreground"
                    >
                        {bio}
                    </Typography>
                )}

                <Stack gap="sm">

                    {location && (
                        <Flex
                            gap="sm"
                            align="center"
                        >
                            <MapPin
                                size={16}
                            />

                            <Typography
                                variant="body-sm"
                            >
                                {location}
                            </Typography>
                        </Flex>
                    )}

                    {joinedDate && (
                        <Flex
                            gap="sm"
                            align="center"
                        >
                            <Calendar
                                size={16}
                            />

                            <Typography
                                variant="body-sm"
                            >
                                Joined {joinedDate}
                            </Typography>
                        </Flex>
                    )}

                </Stack>

                {badges && (
                    <Flex
                        gap="sm"
                        wrap="wrap"
                    >
                        {badges}
                    </Flex>
                )}

            </Stack>
        </Card>
    );
};

export default MemberCard;