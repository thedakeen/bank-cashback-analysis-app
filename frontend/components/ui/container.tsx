import { cn } from "@/lib/utils";
import React from "react";

type Props = {
    children?: React.ReactNode;
};

function Container({ children }: Props) {
    return <div className={cn("container-xl mx-auto px-4")}>{children}</div>;
}

export default Container;
