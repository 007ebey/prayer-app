import {
    Children,
    cloneElement,
    forwardRef,
    isValidElement,
} from "react";

import { cn } from "../../utils/cn";

import { avatarGroupVariants } from "./AvatarGroup.styles";

import type { AvatarGroupProps } from "./AvatarGroup.types";

const AvatarGroup = forwardRef<
    HTMLDivElement,
    AvatarGroupProps
>(
    (
        {
            children,
            max,
            overlap,
            reverse,
            bordered = true,
            className,
            ...props
        },
        ref
    ) => {

        const avatars = Children.toArray(children);

        const visible =
            max
                ? avatars.slice(0, max)
                : avatars;

        const remaining =
            max
                ? avatars.length - visible.length
                : 0;

        return (

            <div
                ref={ref}
                className={cn(
                    avatarGroupVariants({
                        overlap,
                        reverse,
                    }),
                    className
                )}
                {...props}
            >

                {visible.map((child, index) => {

                    if (!isValidElement(child)) {
                        return child;
                    }

                    return (
                        <div key={index}>
                            {child}
                        </div>
                    );

                })}

                {remaining > 0 && (

                    <div
                        className="
                            flex
                            h-10
                            w-10
                            items-center
                            justify-center
                            rounded-full
                            border
                            border-background
                            bg-muted
                            text-xs
                            font-semibold
                        "
                    >
                        +{remaining}
                    </div>

                )}

            </div>

        );

    }
);

AvatarGroup.displayName =
    "AvatarGroup";

export default AvatarGroup;