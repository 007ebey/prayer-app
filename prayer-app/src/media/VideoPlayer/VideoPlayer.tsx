import {
    forwardRef,
    useRef,
} from "react";

import {
    Pause,
    Play,
} from "lucide-react";

import {
    Button,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { videoPlayerVariants } from "./VideoPlayer.styles";

import type { VideoPlayerProps } from "./VideoPlayer.types";

const VideoPlayer = forwardRef<
    HTMLVideoElement,
    VideoPlayerProps
>(
    (
        {
            heading,
            description,
            artwork,
            poster,
            showControls = true,
            className,
            src,
            ...props
        },
        ref
    ) => {
        const videoRef =
            useRef<HTMLVideoElement>(null);

        return (
            <div
                className={cn(
                    videoPlayerVariants(),
                    className
                )}
            >
                {(artwork ||
                    heading ||
                    description) && (
                    <Flex
                        align="center"
                        gap="md"
                    >
                        {artwork}

                        <Stack
                            gap="xs"
                            className="flex-1"
                        >
                            {heading && (
                                <Typography variant="body">
                                    {heading}
                                </Typography>
                            )}

                            {description && (
                                <Typography
                                    variant="caption"
                                    className="text-muted-foreground"
                                >
                                    {description}
                                </Typography>
                            )}
                        </Stack>
                    </Flex>
                )}

                <video
                    ref={(node) => {
                        videoRef.current = node;

                        if (
                            typeof ref ===
                            "function"
                        ) {
                            ref(node);
                        } else if (
                            ref
                        ) {
                            ref.current = node;
                        }
                    }}
                    src={src}
                    poster={poster}
                    controls={showControls}
                    className="aspect-video w-full rounded-md bg-black"
                    {...props}
                />

                {!showControls && (
                    <Flex gap="sm">
                        <Button
                            size="icon"
                            variant="outline"
                            aria-label="Play video"
                            onClick={() => {
                                videoRef.current
                                    ?.play()
                                    .catch(() => {});
                            }}
                        >
                            <Play size={18} />
                        </Button>

                        <Button
                            size="icon"
                            variant="outline"
                            aria-label="Pause video"
                            onClick={() =>
                                videoRef.current?.pause()
                            }
                        >
                            <Pause size={18} />
                        </Button>
                    </Flex>
                )}
            </div>
        );
    }
);

VideoPlayer.displayName =
    "VideoPlayer";

export default VideoPlayer;