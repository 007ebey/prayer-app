import type {
    ImgHTMLAttributes,
    ReactNode,
} from "react";

export type ImageFit =
    | "cover"
    | "contain"
    | "fill"
    | "none"
    | "scale-down";

export type ImageRadius =
    | "none"
    | "sm"
    | "md"
    | "lg"
    | "xl"
    | "full";

export type ImageAspectRatio =
    | "auto"
    | "square"
    | "video"
    | "portrait";

export interface ImageProps
    extends Omit<
        ImgHTMLAttributes<HTMLImageElement>,
        "children"
    > {

    fit?: ImageFit;

    radius?: ImageRadius;

    aspectRatio?: ImageAspectRatio;

    fallback?: ReactNode;

    loadingIndicator?: ReactNode;
}