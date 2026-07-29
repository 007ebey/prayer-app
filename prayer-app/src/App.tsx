import { Church } from "lucide-react";
import { Stack, Logo } from "./../primitives";


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
                        <ThemeSwitcher/>
                     </div>
                     {/* Hero */}
                </Stack>
            </div>
        </div>
    );
};

export default App;