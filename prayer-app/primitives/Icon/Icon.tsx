import type { IconProps } from "./Icon.types";

import { iconSizes } from "./Icon.styles";

const Icon = ({
    icon: LucideIcon,
    size = "md",
    strokeWidth = 2,
    color,
    className,
    ...props
}: IconProps) => {

    return (

        <LucideIcon

            size={iconSizes[size]}

            strokeWidth={strokeWidth}

            color={color}

            className={className}

            {...props}

        />

    );

};

export default Icon;