import { cn } from "../../../utils/cn";

import type { BlankLayoutProps } from "./BlankLayout.types";

const BlankLayout = ({
    children,
    centered = false,
    padded = false,
    background = true,
}: BlankLayoutProps) => {
    return (
        <main
            className={cn(
                "min-h-screen",

                background &&
                    "bg-background text-foreground",

                padded && "p-6",

                centered &&
                    "flex items-center justify-center"
            )}
        >
            {children}
        </main>
    );
};

export default BlankLayout;