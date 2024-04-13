import { cn } from "@/lib/utils";
import { Slot } from "@radix-ui/react-slot";
import { VariantProps, cva } from "class-variance-authority";
import React from "react";

const hstackVariants = cva("flex flex-row", {
    variants: {
        justify: {
            center: "justify-center",
            start: "justify-start",
            end: "justify-end",
            between: "justify-between",
            around: "justify-around",
            evenly: "justify-evenly",
            stretch: "justify-stretch",
        },
        align: {
            stretch: "items-stretch",
            start: "items-start",
            end: "items-end",
            center: "items-center",
            baseline: "items-baseline",
        },
    },
    defaultVariants: {
        justify: "start",
        align: "start",
    },
});

export interface HStackProps
    extends React.HTMLAttributes<HTMLDivElement>,
        VariantProps<typeof hstackVariants> {
    asChild?: boolean;
}

const HStack = React.forwardRef<HTMLDivElement, HStackProps>(
    ({ className, justify, align, asChild = false, ...props }, ref) => {
        const Comp = asChild ? Slot : "div";
        return (
            <Comp
                className={cn(hstackVariants({ justify, align, className }))}
                ref={ref}
                {...props}
            />
        );
    }
);
HStack.displayName = "HStack";

export { HStack, hstackVariants };
