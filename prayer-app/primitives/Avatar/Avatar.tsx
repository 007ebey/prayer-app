import { forwardRef } from "react";

import type { AvatarProps } from "./Avatar.types";

import {
    avatarShapes,
    avatarSizes,
    avatarStatus
} from "./Avatar.styles";

import { cn } from "../../utils/cn";

const Avatar = forwardRef<HTMLDivElement, AvatarProps>(
(
    {
        src,
        alt,
        name,
        size = "md",
        shape = "circle",
        status,
        bordered = false,
        className,
        ...props
    },
    ref
) => {

    const initials =
        name
            ?.split(" ")
            .map(word => word.charAt(0))
            .join("")
            .substring(0, 2)
            .toUpperCase();

    return (

        <div
            ref={ref}
            className={cn(
                "relative inline-flex",
                className
            )}
            {...props}
        >

            {
                src ? (

                    <img

                        src={src}

                        alt={alt ?? name}

                        className={cn(

                            avatarSizes[size],

                            avatarShapes[shape],

                            "object-cover bg-slate-200",

                            bordered &&
                                "ring-2 ring-white shadow-md"

                        )}

                    />

                ) : (

                    <div

                        className={cn(

                            avatarSizes[size],

                            avatarShapes[shape],

                            "bg-primary-500 text-white",

                            "flex items-center justify-center",

                            "font-semibold",

                            bordered &&
                                "ring-2 ring-white shadow-md"

                        )}

                    >

                        {initials ?? "?"}

                    </div>

                )
            }

            {
                status && (

                    <span

                        className={cn(

                            "absolute",

                            "bottom-0 right-0",

                            "h-3.5 w-3.5",

                            "rounded-full",

                            "border-2 border-white",

                            avatarStatus[status]

                        )}

                    />

                )
            }

        </div>

    );

});

Avatar.displayName = "Avatar";

export default Avatar;