import { cn } from "@/lib/utils";
import { Slot } from "@radix-ui/react-slot";
import { VariantProps, cva } from "class-variance-authority";
import React from "react";

const pageVariants = cva("bg-accent h-svh", {
    variants: {},
    defaultVariants: {},
});

export interface PageProps
    extends React.HTMLAttributes<HTMLDivElement>,
        VariantProps<typeof pageVariants> {
    asChild?: boolean;
}

const Page = React.forwardRef<HTMLDivElement, PageProps>(
    ({ className, asChild = false, ...props }, ref) => {
        const Comp = asChild ? Slot : "div";
        return (
            <Comp
                className={cn(pageVariants({ className }))}
                ref={ref}
                {...props}
            />
        );
    }
);
Page.displayName = "Page";

export { Page, pageVariants as hstackVariants };
