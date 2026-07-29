import type {
    HTMLAttributes
} from "react";

import type {
    LucideIcon
} from "lucide-react";

export type IconSize =
    | "xs"
    | "sm"
    | "md"
    | "lg"
    | "xl";

export interface IconProps
    extends HTMLAttributes<SVGSVGElement> {

    icon: LucideIcon;

    size?: IconSize;

    strokeWidth?: number;

    color?: string;
}