import {
    ReactNode,
    VideoHTMLAttributes,
} from "react";

export interface VideoPlayerProps
    extends Omit<
        VideoHTMLAttributes<HTMLVideoElement>,
        "children"
    > {

    heading?: ReactNode;

    description?: ReactNode;

    poster?: string;

    artwork?: ReactNode;

    showControls?: boolean;
}