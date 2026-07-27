// New file generated from Profile.tsx
import { cn } from "../../utils/cn";

import {
    profileGridVariants,
    profileVariants,
} from "./Profile.styles";

import type {
    ProfileProps,
} from "./Profile.types";

const Profile = ({
    header,
    userCard,
    statistics,
    personalInformation,
    accountSettings,
    prayerHistory,
    achievements,
    activity,
    footer,
}: ProfileProps) => {

    return (

        <main className={cn(profileVariants())}>

            {header}

            <div className={cn(profileGridVariants())}>

                <aside className="space-y-6">

                    {userCard}

                    {statistics}

                    {achievements}

                </aside>

                <section className="space-y-6">

                    {personalInformation}

                    {accountSettings}

                    {prayerHistory}

                    {activity}

                </section>

            </div>

            {footer}

        </main>

    );

};

export default Profile;