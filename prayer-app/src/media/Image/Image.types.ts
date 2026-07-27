import {
    ImgHTMLAttributes,
    ReactNode,
} from "react";

export type ImageFit =
    | "cover"
    | "contain"
    | "fill"
    | "none"
    | "scale-down";

export interface ImageProps
    extends Omit<
        ImgHTMLAttributes<HTMLImageElement>,
        "children"
    > {

    fallback?: ReactNode;

    fit?: ImageFit;

    rounded?: boolean;
}