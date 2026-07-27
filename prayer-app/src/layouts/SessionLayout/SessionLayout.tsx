import type { SessionLayoutProps } from "./SessionLayout.types";

const SessionLayout = ({
    children,
    header,
    prayerPanel,
    participantsPanel,
    chatPanel,
    footer,
}: SessionLayoutProps) => {
    return (
        <div className="flex min-h-screen flex-col bg-background text-foreground">

            {header}

            <div className="flex flex-1 overflow-hidden">

                {prayerPanel && (
                    <aside className="hidden w-80 border-r lg:block">
                        {prayerPanel}
                    </aside>
                )}

                <main className="flex flex-1 flex-col overflow-hidden">
                    {children}
                </main>

                {(participantsPanel || chatPanel) && (
                    <aside className="hidden w-80 flex-col border-l xl:flex">

                        {participantsPanel && (
                            <div className="flex-1 overflow-auto border-b">
                                {participantsPanel}
                            </div>
                        )}

                        {chatPanel && (
                            <div className="flex-1 overflow-auto">
                                {chatPanel}
                            </div>
                        )}

                    </aside>
                )}

            </div>

            {footer}

        </div>
    );
};

export default SessionLayout;