import { cn } from "@/lib/utils";
import { Slot } from "@radix-ui/react-slot";
import { VariantProps, cva } from "class-variance-authority";
import React from "react";

const vstackVariants = cva("flex flex-col", {
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

export interface VStackProps
    extends React.HTMLAttributes<HTMLDivElement>,
        VariantProps<typeof vstackVariants> {
    asChild?: boolean;
}

const VStack = React.forwardRef<HTMLDivElement, VStackProps>(
    ({ className, justify, align, asChild = false, ...props }, ref) => {
        const Comp = asChild ? Slot : "div";
        return (
            <Comp
                className={cn(vstackVariants({ justify, align, className }))}
                ref={ref}
                {...props}
            />
        );
    }
);
VStack.displayName = "VStack";

export { VStack, vstackVariants };
