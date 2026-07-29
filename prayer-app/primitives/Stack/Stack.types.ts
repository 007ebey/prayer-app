import type {
    HTMLAttributes,
    ReactNode,
} from "react";

import type {
    FlexGap,
    FlexAlign,
    FlexJustify,
} from "../Flex";

export interface StackProps
    extends HTMLAttributes<HTMLDivElement> {

    children: ReactNode;

    gap?: FlexGap;

    align?: FlexAlign;

    justify?: FlexJustify;

    reverse?: boolean;
}