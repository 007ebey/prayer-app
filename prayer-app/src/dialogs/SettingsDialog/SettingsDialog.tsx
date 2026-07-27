// New file generated from SettingsDialog.tsx
import {
    useEffect,
    useState,
} from "react";

import {
    Bell,
    Moon,
    Volume2,
    Zap,
} from "lucide-react";

import {
    Button,
    Dialog,
    Divider,
    Flex,
    Stack,
    Switch,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import {
    settingsDialogVariants,
} from "./SettingsDialog.styles";

import type {
    SettingsDialogProps,
    SettingsValues,
} from "./SettingsDialog.types";

const SettingsDialog = ({
    open,
    values,
    loading = false,
    onClose,
    onSave,
    className,
    ...props
}: SettingsDialogProps) => {

    const [
        settings,
        setSettings,
    ] = useState(values);

    useEffect(() => {

        if (open) {
            setSettings(values);
        }

    }, [
        open,
        values,
    ]);

    const updateSetting = (
        key: keyof SettingsValues,
        value: boolean
    ) => {

        setSettings(previous => ({
            ...previous,
            [key]: value,
        }));

    };

    return (
        <Dialog
            open={open}
            onClose={onClose}
            heading="Settings"
            description="Customize your prayer experience."
            size="md"
            className={cn(
                settingsDialogVariants(),
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
                        onClick={() =>
                            onSave(settings)
                        }
                    >
                        Save Changes
                    </Button>
                </Flex>
            }
            {...props}
        >
            <Stack gap="lg">

                <Flex
                    justify="between"
                    align="center"
                >
                    <Flex gap="sm" align="center">
                        <Bell size={18} />
                        <Typography>
                            Notifications
                        </Typography>
                    </Flex>

                    <Switch
                        checked={
                            settings.notifications
                        }
                        onCheckedChange={(
                            checked
                        ) =>
                            updateSetting(
                                "notifications",
                                checked
                            )
                        }
                    />
                </Flex>

                <Divider />

                <Flex
                    justify="between"
                    align="center"
                >
                    <Flex gap="sm" align="center">
                        <Bell size={18} />
                        <Typography>
                            Prayer Reminders
                        </Typography>
                    </Flex>

                    <Switch
                        checked={
                            settings.prayerReminders
                        }
                        onCheckedChange={(
                            checked
                        ) =>
                            updateSetting(
                                "prayerReminders",
                                checked
                            )
                        }
                    />
                </Flex>

                <Divider />

                <Flex
                    justify="between"
                    align="center"
                >
                    <Flex gap="sm" align="center">
                        <Volume2 size={18} />
                        <Typography>
                            Session Sounds
                        </Typography>
                    </Flex>

                    <Switch
                        checked={
                            settings.sessionSounds
                        }
                        onCheckedChange={(
                            checked
                        ) =>
                            updateSetting(
                                "sessionSounds",
                                checked
                            )
                        }
                    />
                </Flex>

                <Divider />

                <Flex
                    justify="between"
                    align="center"
                >
                    <Flex gap="sm" align="center">
                        <Zap size={18} />
                        <Typography>
                            Auto Join Live Sessions
                        </Typography>
                    </Flex>

                    <Switch
                        checked={
                            settings.autoJoin
                        }
                        onCheckedChange={(
                            checked
                        ) =>
                            updateSetting(
                                "autoJoin",
                                checked
                            )
                        }
                    />
                </Flex>

                <Divider />

                <Flex
                    justify="between"
                    align="center"
                >
                    <Flex gap="sm" align="center">
                        <Moon size={18} />
                        <Typography>
                            Dark Mode
                        </Typography>
                    </Flex>

                    <Switch
                        checked={
                            settings.darkMode
                        }
                        onCheckedChange={(
                            checked
                        ) =>
                            updateSetting(
                                "darkMode",
                                checked
                            )
                        }
                    />
                </Flex>

            </Stack>
        </Dialog>
    );
};

export default SettingsDialog;