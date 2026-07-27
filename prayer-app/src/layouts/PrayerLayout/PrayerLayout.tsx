import type { PrayerLayoutProps } from "./PrayerLayout.types";

const PrayerLayout = ({
    header,
    prayerList,
    prayerWall,
    participants,
    footer,
    children,
}: PrayerLayoutProps) => {
    return (
        <div className="flex min-h-screen flex-col bg-background text-foreground">

            {header}

            <div className="flex flex-1 overflow-hidden">

                {/* Prayer List */}
                {prayerList && (
                    <aside className="hidden w-72 border-r lg:block">
                        {prayerList}
                    </aside>
                )}

                {/* Prayer Wall */}
                <main className="flex flex-1 flex-col overflow-y-auto">

                    {prayerWall}

                    {children}

                </main>

                {/* Participants */}
                {participants && (
                    <aside className="hidden w-80 border-l xl:block">
                        {participants}
                    </aside>
                )}

            </div>

            {footer}

        </div>
    );
};

export default PrayerLayout;