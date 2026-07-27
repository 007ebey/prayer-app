import * as PopoverPrimitive from "@radix-ui/react-popover";

import type {
    PopoverProps,
} from "./Popover.types";

const Popover = ({
    trigger,
    children,
    open,
    defaultOpen,
    onOpenChange,
    side = "bottom",
    align = "start",
}: PopoverProps) => {

    return (

        <PopoverPrimitive.Root
            {...(
                open !== undefined
                    ? { open }
                    : {}
            )}
            {...(
                defaultOpen !== undefined
                    ? { defaultOpen }
                    : {}
            )}
            {...(
                onOpenChange !== undefined
                    ? { onOpenChange }
                    : {}
            )}
        >

            <PopoverPrimitive.Trigger
                asChild
            >
                {trigger}
            </PopoverPrimitive.Trigger>

            <PopoverPrimitive.Portal>

                <PopoverPrimitive.Content
                    side={side}
                    align={align}
                    sideOffset={8}
                    className="z-50 rounded-xl border bg-background p-2 shadow-lg"
                >
                    {children}

                    <PopoverPrimitive.Arrow />

                </PopoverPrimitive.Content>

            </PopoverPrimitive.Portal>

        </PopoverPrimitive.Root>

    );

};

export default Popover;