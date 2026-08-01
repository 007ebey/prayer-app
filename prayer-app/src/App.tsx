import { Church } from "lucide-react";
import { Stack, Logo, Surface, Badge, Typography, Avatar, AvatarGroup, Button } from "./../primitives";


// import BorderRadius from "./theme/BorderRadius";
// import ColorPalette from "./theme/ColorPalette";
// import Elevation from "./theme/Elevation";
// import Spacing from "./theme/Spacing";
// import ThemeSwitcher from "./theme/ThemeSwitcher";
// import TypographyScale from "./theme/TypographyScale";

import "./index.css";
import ThemeSwitcher from "./theme/ThemeSwitcher";

const App = () => {
    return (
        <div className="min-h-screen bg-background">
            <div className="mx-auto max-w-7xl p-8">
                <Stack gap="2xl">
                    {/* Header */}
                    <div className="flex items-center justify-between">
                        <Logo
                            icon={<Church />}
                            name="Prayer App"
                        />
                        <ThemeSwitcher />
                    </div>
                    {/* Hero */}
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

                            <AvatarGroup>

                                <Avatar name="PE" />

                                <Avatar name="JS" />

                                <Avatar name="AM" />

                                <Avatar name="RK" />

                            </AvatarGroup>

                            <Button
                                size="lg"
                            >
                                Join Prayer
                            </Button>
                        </Stack>
                    </Surface>
                </Stack>
            </div>
        </div>
    );
};

export default App;