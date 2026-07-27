// New file generated from EmptyStateCard.tsx
import {
    Card,
    Stack,
    Typography,
} from "../../../primitives";

import { cn } from "../../utils/cn";

import { emptyStateCardVariants } from "./EmptyStateCard.styles";

import type { EmptyStateCardProps } from "./EmptyStateCard.types";

const EmptyStateCard = ({
    icon,
    illustration,
    heading,
    description,
    actions,
    className,
    ...props
}: EmptyStateCardProps) => {

    return (
        <Card
            className={cn(
                emptyStateCardVariants(),
                className
            )}
            {...props}
        >
            <Stack
                gap="lg"
                align="center"
            >

                {illustration ?? (
                    icon && (
                        <div className="mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-muted text-muted-foreground">
                            {icon}
                        </div>
                    )
                )}

                <Stack gap="sm">

                    <Typography
                        variant="h4"
                    >
                        {heading}
                    </Typography>

                    {description && (
                        <Typography
                            variant="body"
                            className="text-muted-foreground"
                        >
                            {description}
                        </Typography>
                    )}

                </Stack>

                {actions}

            </Stack>
        </Card>
    );
};

export default EmptyStateCard;