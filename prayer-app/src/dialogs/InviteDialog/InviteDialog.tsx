// New file generated from InviteDialog.tsx
import {
    useEffect,
    useState,
} from "react";

import {
    Mail,
} from "lucide-react";

import {
    Button,
    Dialog,
    Flex,
    Input,
    Stack,
    Textarea,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import {
    inviteDialogVariants,
} from "./InviteDialog.styles";

import type {
    InviteDialogProps,
} from "./InviteDialog.types";

const InviteDialog = ({
    open,
    onClose,
    onInvite,
    heading = "Invite Members",
    description = "Invite people by entering one or more email addresses.",
    defaultMessage = "",
    loading = false,
    className,
    ...props
}: InviteDialogProps) => {

    const [
        emails,
        setEmails,
    ] = useState("");

    const [
        message,
        setMessage,
    ] = useState(defaultMessage);

    useEffect(() => {

        if (open) {

            setEmails("");
            setMessage(defaultMessage);

        }

    }, [
        open,
        defaultMessage,
    ]);

    const handleInvite = () => {

        onInvite({
            emails,
            message,
        });

    };

    return (
        <Dialog
            open={open}
            onClose={onClose}
            heading={heading}
            description={description}
            className={cn(
                inviteDialogVariants(),
                className
            )}
            footer={
                <Flex
                    justify="end"
                    gap="md"
                >
                    <Button
                        variant="outline"
                        onClick={onClose}
                        disabled={loading}
                    >
                        Cancel
                    </Button>

                    <Button
                        loading={loading}
                        onClick={handleInvite}
                    >
                        Send Invite
                    </Button>
                </Flex>
            }
            {...props}
        >
            <Stack gap="lg">

                <Input
                    label="Email Addresses"
                    placeholder="john@example.com, jane@example.com"
                    value={emails}
                    onChange={(event) =>
                        setEmails(
                            event.target.value
                        )
                    }
                    leftIcon={<Mail size={18} />}
                    helperText="Separate multiple email addresses with commas."
                />

                <Textarea
                    label="Personal Message"
                    placeholder="I'd love for you to join our prayer session."
                    rows={5}
                    value={message}
                    onChange={(event) =>
                        setMessage(
                            event.target.value
                        )
                    }
                />

            </Stack>
        </Dialog>
    );
};

export default InviteDialog;