import { Clock, Users } from "lucide-react";

import {
    Stack,
    Surface,
    Badge,
    Typography,
    Avatar,
    AvatarGroup,
    Button,
} from "./../primitives";

interface ActivePrayerProps {
    onJoin: () => void;
}

const ActivePrayer = ({ onJoin }: ActivePrayerProps) => {
    return (
        <Surface
            elevation="raised"
            padding="xl"
            radius="xl"
        >
            <Stack gap="lg">

                <Badge color="success">
                    LIVE NOW
                </Badge>

                <Typography variant="display">
                    Evening Prayer Meeting
                </Typography>

                <Typography
                    variant="subtitle"
                    color="muted"
                >
                    Join believers across the world
                    praying together.
                </Typography>

                <div className="flex flex-wrap items-center gap-6">

                    <div className="flex items-center gap-3">

                        <AvatarGroup>
                            <Avatar name="PE" />
                            <Avatar name="JS" />
                            <Avatar name="AM" />
                            <Avatar name="RK" />
                        </AvatarGroup>

                        <div className="flex items-center gap-2 text-muted-foreground">
                            <Users className="h-4 w-4" />

                            <span className="text-sm">
                                24 praying now
                            </span>
                        </div>

                    </div>

                    <div className="flex items-center gap-2 text-muted-foreground">
                        <Clock className="h-4 w-4" />

                        <span className="text-sm">
                            7:00 PM • 45 min
                        </span>
                    </div>

                </div>

                <Button
                    size="lg"
                    fullWidth
                    onClick={onJoin}
                >
                    Join Prayer
                </Button>

            </Stack>
        </Surface>
    );
};

export default ActivePrayer;