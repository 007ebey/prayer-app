import type { DashboardLayoutProps } from "./DashboardLayout.types";

const DashboardLayout = ({
    children,
    header,
    sidebar,
    mobileHeader,
    bottomNavigation,
    footer,
}: DashboardLayoutProps) => {
    return (
        <div className="flex min-h-screen bg-background text-foreground">

            {/* Desktop Sidebar */}
            {sidebar && (
                <aside className="hidden lg:flex">
                    {sidebar}
                </aside>
            )}

            <div className="flex min-h-screen flex-1 flex-col">

                {/* Mobile Header */}
                {mobileHeader && (
                    <div className="lg:hidden">
                        {mobileHeader}
                    </div>
                )}

                {/* Desktop Header */}
                {header && (
                    <div className="hidden lg:block">
                        {header}
                    </div>
                )}

                <main className="flex-1 overflow-y-auto p-6">
                    {children}
                </main>

                {footer}

                {/* Mobile Bottom Navigation */}
                {bottomNavigation && (
                    <div className="lg:hidden">
                        {bottomNavigation}
                    </div>
                )}

            </div>

        </div>
    );
};

export default DashboardLayout;