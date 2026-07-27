import {
    AudioHTMLAttributes,
    ReactNode,
} from "react";

export interface AudioPlayerProps
    extends Omit<
        AudioHTMLAttributes<HTMLAudioElement>,
        "children"
    > {

    heading?: ReactNode;

    description?: ReactNode;

    artwork?: ReactNode;

    showControls?: boolean;

    autoPlay?: boolean;
}