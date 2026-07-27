// New file generated from AuditLog.tsx
import {
    Badge,
    Card,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../../utils/cn";

import { auditLogVariants } from "./AuditLog.styles";

import type {
    AuditLogItem,
    AuditLogProps,
} from "./AuditLog.types";


const actionColor = (
    action: AuditLogItem["action"]
) => {
    switch (action) {
        case "create":
        case "approve":
            return "success";

        case "delete":
        case "reject":
            return "danger";

        case "update":
        case "assign":
            return "warning";

        case "login":
        case "logout":
            return "info";

        case "remove":
            return "secondary";

        default:
            return "neutral";
    }
};

const AuditLog = ({
    heading,
    description,
    logs,
    loading = false,
    emptyState,
}: AuditLogProps) => {
    return (
        <Card className={cn(auditLogVariants())}>
            {(heading || description) && (
                <Stack
                    gap="xs"
                    className="p-6"
                >
                    {heading && (
                        <Typography variant="h4">
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
            )}

            <Divider />

            {loading ? (
                <div className="p-6">
                    <Typography>
                        Loading audit logs...
                    </Typography>
                </div>
            ) : logs.length === 0 ? (
                emptyState ?? (
                    <div className="p-6">
                        <Typography>
                            No audit records found.
                        </Typography>
                    </div>
                )
            ) : (
                <Stack gap="none">
                    {logs.map((log) => (
                        <div
                            key={log.id}
                            className="border-b p-6 last:border-b-0"
                        >
                            <Flex
                                justify="between"
                                align="start"
                            >
                                <Stack gap="xs">
                                    <Flex
                                        gap="sm"
                                        align="center"
                                    >
                                        <Typography
                                            variant="body"
                                            weight="medium"
                                        >
                                            {log.user}
                                        </Typography>

                                        <Badge
                                            variant="soft"
                                            color={actionColor(log.action)}
                                        >
                                            {log.action}
                                        </Badge>
                                    </Flex>

                                    <Typography
                                        variant="body-sm"
                                    >
                                        {log.target}
                                    </Typography>

                                    {log.description && (
                                        <Typography
                                            variant="caption"
                                            className="text-muted-foreground"
                                        >
                                            {log.description}
                                        </Typography>
                                    )}
                                </Stack>

                                <Stack
                                    gap="xs"
                                    align="end"
                                >
                                    <Typography
                                        variant="caption"
                                    >
                                        {log.timestamp}
                                    </Typography>

                                    {log.ipAddress && (
                                        <Typography
                                            variant="caption"
                                            className="text-muted-foreground"
                                        >
                                            {log.ipAddress}
                                        </Typography>
                                    )}
                                </Stack>
                            </Flex>
                        </div>
                    ))}
                </Stack>
            )}
        </Card>
    );
};

export default AuditLog;