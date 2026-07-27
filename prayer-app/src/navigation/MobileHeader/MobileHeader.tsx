import {
    Card,
    Divider,
    Flex,
    Typography,
} from "../../../primitives";

import type { MobileHeaderProps } from "./MobileHeader.types";

const MobileHeader = ({
    title,
    logo,
    menuButton,
    endContent,
    sticky = true,
    bordered = true,
}: MobileHeaderProps) => {
    return (
        <header
            className={
                sticky
                    ? "sticky top-0 z-50 md:hidden"
                    : "md:hidden"
            }
        >
            <Card className="rounded-none">
                <Flex
                    justify="between"
                    align="center"
                    className="h-16 px-4"
                >
                    <Flex
                        align="center"
                        gap="md"
                    >
                        {menuButton}

                        {logo}

                        {typeof title === "string" ? (
                            <Typography variant="title">
                                {title}
                            </Typography>
                        ) : (
                            title
                        )}
                    </Flex>

                    {endContent}
                </Flex>

                {bordered && <Divider />}
            </Card>
        </header>
    );
};

export default MobileHeader;