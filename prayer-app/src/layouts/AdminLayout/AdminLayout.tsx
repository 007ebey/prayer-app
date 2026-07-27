import {
    Flex,
} from "../../../primitives";

import type {
    AdminLayoutProps,
} from "./AdminLayout.types";

const AdminLayout = ({
    header,
    mobileHeader,
    sidebar,
    footer,
    children,
}: AdminLayoutProps) => {
    return (
        <div className="min-h-screen bg-background text-foreground">

            {mobileHeader}

            {header}

            <Flex className="min-h-[calc(100vh-4rem)]">

                {sidebar}

                <main className="flex-1 overflow-auto p-6">

                    {children}

                </main>

            </Flex>

            {footer}

        </div>
    );
};

export default AdminLayout;