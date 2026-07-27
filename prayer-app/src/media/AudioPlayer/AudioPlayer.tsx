import {
    forwardRef,
    useRef,
} from "react";

import {
    Button,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { audioPlayerVariants } from "./AudioPlayer.styles";

import type { AudioPlayerProps } from "./AudioPlayer.types";

const AudioPlayer = forwardRef<
    HTMLAudioElement,
    AudioPlayerProps
>(
    (
        {
            heading,
            description,
            artwork,
            showControls = true,
            className,
            src,
            ...props
        },
        ref
    ) => {
        const audioRef =
            useRef<HTMLAudioElement>(null);

        return (
            <div
                className={cn(
                    audioPlayerVariants(),
                    className
                )}
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

                    <audio
                        ref={(node) => {
                            audioRef.current =
                                node;

                            if (
                                typeof ref ===
                                "function"
                            ) {
                                ref(node);
                            } else if (
                                ref
                            ) {
                                ref.current =
                                    node;
                            }
                        }}
                        src={src}
                        controls={
                            showControls
                        }
                        {...props}
                        className="w-full"
                    />
                </Stack>

                {!showControls && (
                    <Flex gap="sm">
                        <Button
                            size="icon"
                            variant="outline"
                            onClick={() =>
                                audioRef.current?.play()
                            }
                        >
                            ▶
                        </Button>

                        <Button
                            size="icon"
                            variant="outline"
                            onClick={() =>
                                audioRef.current?.pause()
                            }
                        >
                            ❚❚
                        </Button>
                    </Flex>
                )}
            </div>
        );
    }
);

AudioPlayer.displayName =
    "AudioPlayer";

export default AudioPlayer;