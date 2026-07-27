import { Fragment } from "react";

import {
    Button,
    Typography,
} from "../../../primitives";

import type {
    BreadcrumbItem,
    BreadcrumbProps,
} from "./Breadcrumb.types";

const Breadcrumb = ({
    items,
    separator = "/",
    onItemClick,
}: BreadcrumbProps) => {
    return (
        <nav
            aria-label="Breadcrumb"
            className="flex items-center flex-wrap gap-2"
        >
            {items.map((item, index) => {
                const isLast =
                    index === items.length - 1;

                return (
                    <Fragment key={item.id}>
                        {isLast ? (
                            <Typography
                                variant="body-sm"
                                weight="semibold"
                            >
                                {item.icon}

                                {item.label}
                            </Typography>
                        ) : (
                            <Button
                                variant="ghost"
                                size="sm"
                                onClick={() =>
                                    onItemClick?.(
                                        item
                                    )
                                }
                                className="h-auto p-0"
                            >
                                <span className="flex items-center gap-2">
                                    {item.icon}

                                    {item.label}
                                </span>
                            </Button>
                        )}

                        {!isLast && (
                            <Typography
                                variant="body-sm"
                                className="text-muted-foreground"
                            >
                                {separator}
                            </Typography>
                        )}
                    </Fragment>
                );
            })}
        </nav>
    );
};

export default Breadcrumb;