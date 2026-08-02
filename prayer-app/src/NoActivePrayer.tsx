import {
    Bell,
    CalendarDays,
    Clock,
} from "lucide-react";

import {
    Stack,
    Surface,
    Badge,
    Typography,
    Button,
} from "./../primitives";

const NoActivePrayer = () => {
    return (
        <Surface
            elevation="raised"
            padding="xl"
            radius="xl"
        >
            <Stack gap="lg">

                <Badge>
                    NO LIVE SESSION
                </Badge>

                <Typography variant="display">
                    No prayer session is live right now
                </Typography>

                <Typography
                    variant="subtitle"
                    color="muted"
                >
                    The next prayer session starts tomorrow morning.
                </Typography>

                {/* Next session */}

                <div className="rounded-xl border border-border p-5">

                    <Stack gap="md">

                        <Typography
                            variant="body-sm"
                            color="muted"
                        >
                            UP NEXT
                        </Typography>

                        <Typography variant="h3">
                            Morning Prayer
                        </Typography>

                        <div className="flex flex-wrap items-center gap-6 text-muted-foreground">

                            <div className="flex items-center gap-2">
                                <CalendarDays className="h-4 w-4" />

                                <span className="text-sm">
                                    Tomorrow
                                </span>
                            </div>

                            <div className="flex items-center gap-2">
                                <Clock className="h-4 w-4" />

                                <span className="text-sm">
                                    6:00 AM
                                </span>
                            </div>

                            <span className="text-sm">
                                30 min
                            </span>

                        </div>

                    </Stack>

                </div>

                <Button
                    variant="outline"
                    size="lg"
                    fullWidth
                    leftIcon={
                        <Bell className="h-4 w-4" />
                    }
                >
                    Remind Me
                </Button>

            </Stack>
        </Surface>
    );
};

export default NoActivePrayer;