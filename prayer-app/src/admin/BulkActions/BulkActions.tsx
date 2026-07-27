// New file generated from BulkActions.tsx
import {
    Button,
    Divider,
    Flex,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { bulkActionsVariants } from "./BulkActions.styles";

import type {
    BulkAction,
    BulkActionsProps,
} from "./BulkActions.types";

const BulkActions = ({
    selectedCount,
    actions,
    onClearSelection,
    className,
    ...props
}: BulkActionsProps) => {
    if (selectedCount === 0) {
        return null;
    }

    return (
        <div
            className={cn(
                bulkActionsVariants(),
                className
            )}
            {...props}
        >
            <Flex
                align="center"
                gap="md"
            >
                <Typography
                    variant="body"
                    weight="medium"
                >
                    {selectedCount} selected
                </Typography>

                {onClearSelection && (
                    <Button
                        variant="ghost"
                        size="sm"
                        onClick={onClearSelection}
                    >
                        Clear
                    </Button>
                )}
            </Flex>

            <Flex
                align="center"
                gap="sm"
            >
                {actions.map(
                    (action: BulkAction) => (
                        <Button
                            key={action.id}
                            variant="outline"
                            size="sm"
                            disabled={
                                action.disabled
                            }
                            leftIcon={
                                action.icon
                            }
                            onClick={
                                action.onClick
                            }
                        >
                            {action.label}
                        </Button>
                    )
                )}
            </Flex>
        </div>
    );
};

export default BulkActions;