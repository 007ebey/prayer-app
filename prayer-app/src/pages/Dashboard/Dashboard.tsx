// New file generated from Dashboard.tsx
import { cn } from "../../utils/cn";

import {
    dashboardVariants,
} from "./Dashboard.styles";

import type {
    DashboardProps,
} from "./Dashboard.types";

const Dashboard = ({
    welcome,
    statistics,
    activeSession,
    upcomingSessions,
    prayerRequests,
    onlineMembers,
    announcements,
    quickActions,
}: DashboardProps) => {

    return (

        <div className="space-y-6">

            {welcome}

            {statistics}

            <div
                className={cn(
                    dashboardVariants()
                )}
            >

                <div className="space-y-6">

                    {activeSession}

                    {prayerRequests}

                    {announcements}

                </div>

                <div className="space-y-6">

                    {upcomingSessions}

                    {onlineMembers}

                    {quickActions}

                </div>

            </div>

        </div>

    );

};

export default Dashboard;