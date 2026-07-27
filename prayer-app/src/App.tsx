import { Stack } from "./../primitives";


import BorderRadius from "./theme/BorderRadius";
import ColorPalette from "./theme/ColorPalette";
import Elevation from "./theme/Elevation";
import Spacing from "./theme/Spacing";
import ThemeSwitcher from "./theme/ThemeSwitcher";
import TypographyScale from "./theme/TypographyScale";

import "./index.css";

const App = () => {
    return (
        <div className="min-h-screen bg-background text-foreground">
            <div className="mx-auto max-w-7xl p-8">
                <Stack gap="xl">
                    <ThemeSwitcher />

                    <ColorPalette />

                    <TypographyScale />

                    <Spacing />

                    <BorderRadius />

                    <Elevation />
                </Stack>
            </div>
        </div>
    );
};

export default App;