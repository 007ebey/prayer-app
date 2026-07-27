import {
    Copy,
    Link2,
    Mail,
    MessageCircle,
    Send,
    X,
} from "lucide-react";

import {
    Button,
    Dialog,
    Flex,
    Input,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    shareDialogVariants,
    shareGridVariants,
} from "./ShareDialog.styles";

import type {
    ShareDialogProps,
    ShareOption,
} from "./ShareDialog.types";

const options = [
    {
        key: "copy",
        label: "Copy Link",
        icon: Copy,
    },
    {
        key: "email",
        label: "Email",
        icon: Mail,
    },
    {
        key: "whatsapp",
        label: "WhatsApp",
        icon: MessageCircle,
    },
    {
        key: "telegram",
        label: "Telegram",
        icon: Send,
    },
    {
        key: "x",
        label: "X",
        icon: X,
    },
] satisfies {
    key: ShareOption;
    label: string;
    icon: typeof Copy;
}[];

const ShareDialog = ({
    open,
    onClose,
    url,
    title = "Share",
    description = "Share this with others.",
    enabledOptions = options.map(
        option => option.key
    ),
    onShare,
    className,
    ...props
}: ShareDialogProps) => {

    const visibleOptions =
        options.filter(option =>
            enabledOptions.includes(
                option.key
            )
        );

    return (
        <Dialog
            open={open}
            onClose={onClose}
            heading={title}
            description={description}
            size="md"
            className={cn(
                shareDialogVariants(),
                className
            )}
            {...props}
        >
            <Stack gap="lg">

                <Input
                    readOnly
                    value={url}
                    leftIcon={
                        <Link2
                            size={18}
                        />
                    }
                />

                <div
                    className={shareGridVariants()}
                >
                    {visibleOptions.map(
                        option => {

                            const Icon =
                                option.icon;

                            return (
                                <Button
                                    key={
                                        option.key
                                    }
                                    variant="outline"
                                    className="h-auto flex-col py-4"
                                    onClick={() =>
                                        onShare?.(
                                            option.key
                                        )
                                    }
                                >
                                    <Icon
                                        size={22}
                                    />

                                    <Typography
                                        variant="caption"
                                    >
                                        {
                                            option.label
                                        }
                                    </Typography>
                                </Button>
                            );

                        }
                    )}
                </div>

                <Flex
                    justify="end"
                >
                    <Button
                        variant="ghost"
                        onClick={
                            onClose
                        }
                    >
                        Close
                    </Button>
                </Flex>

            </Stack>
        </Dialog>
    );
};

export default ShareDialog;