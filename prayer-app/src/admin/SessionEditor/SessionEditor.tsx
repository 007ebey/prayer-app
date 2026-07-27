// New file generated from SessionEditor.tsx
import {
    Button,
    Card,
    Divider,
    Input,
    Stack,
    Switch,
    Textarea,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { sessionEditorVariants } from "./SessionEditor.styles";

import type {
    SessionEditorProps,
    SessionEditorValues,
} from "./SessionEditor.types";

const SessionEditor = ({
    heading,
    description,
    values,
    loading = false,
    onValuesChange,
    onSubmit,
    onCancel,
    className,
    ...props
}: SessionEditorProps) => {

    const update = <
        K extends keyof SessionEditorValues
    >(
        key: K,
        value: SessionEditorValues[K]
    ) => {
        onValuesChange?.({
            ...values,
            [key]: value,
        });
    };

    return (
        <Card
            className={cn(
                sessionEditorVariants(),
                className
            )}
            {...props}
        >
            <form onSubmit={onSubmit}>
                <Stack gap="lg">

                    {(heading || description) && (
                        <>
                            <Stack gap="xs">

                                {heading && (
                                    <Typography variant="h3">
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

                            <Divider />
                        </>
                    )}

                    <Input
                        label="Session Title"
                        value={values.title}
                        onChange={(event) =>
                            update(
                                "title",
                                event.target.value
                            )
                        }
                    />

                    <Textarea
                        label="Description"
                        value={values.description}
                        onChange={(event) =>
                            update(
                                "description",
                                event.target.value
                            )
                        }
                    />

                    <Input
                        label="Host ID"
                        value={values.hostId}
                        onChange={(event) =>
                            update(
                                "hostId",
                                event.target.value
                            )
                        }
                    />

                    <Input
                        label="Background Image"
                        value={
                            values.backgroundImage ??
                            ""
                        }
                        onChange={(event) =>
                            update(
                                "backgroundImage",
                                event.target.value
                            )
                        }
                    />

                    <Input
                        type="datetime-local"
                        label="Start"
                        value={values.startDate}
                        onChange={(event) =>
                            update(
                                "startDate",
                                event.target.value
                            )
                        }
                    />

                    <Input
                        type="datetime-local"
                        label="End"
                        value={values.endDate}
                        onChange={(event) =>
                            update(
                                "endDate",
                                event.target.value
                            )
                        }
                    />

                    <Switch
                        label="Public Session"
                        checked={values.isPublic}
                        onCheckedChange={(checked) =>
                            update(
                                "isPublic",
                                checked
                            )
                        }
                    />

                    <Switch
                        label="Active"
                        checked={values.isActive}
                        onCheckedChange={(checked) =>
                            update(
                                "isActive",
                                checked
                            )
                        }
                    />

                    <Divider />

                    <div className="flex justify-end gap-3">

                        {onCancel && (
                            <Button
                                variant="outline"
                                type="button"
                                onClick={onCancel}
                            >
                                Cancel
                            </Button>
                        )}

                        <Button
                            type="submit"
                            loading={loading}
                        >
                            Save Session
                        </Button>

                    </div>

                </Stack>
            </form>
        </Card>
    );
};

export default SessionEditor;