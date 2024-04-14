"use client";

import { ImSpinner8 } from "react-icons/im";

import Container from "@/components/ui/container";
import { Page } from "@/components/ui/page";
import { VStack } from "@/components/ui/vstack";
import { useCashbacks } from "@/lib/queries/useCashbacks";
import { cn } from "@/lib/utils";
import { useInView } from "framer-motion";
import React, { useEffect, useRef, useState } from "react";
import Link from "next/link";
import { HStack } from "@/components/ui/hstack";
import { Cashback } from "@/types";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

type Props = {};

function CashbacksPage({}: Props) {
    const ref = useRef(null);
    const isInView = useInView(ref);
    const [page, setPage] = useState(1);
    const { data, isLoading } = useCashbacks(page);
    const [cashbacks, setCashbacks] = useState<Cashback[]>([]);
    const [bankName, setBankName] = useState("");

    const searchByBank: React.FormEventHandler<HTMLFormElement> = (e) => {
        console.log(e);
    };

    useEffect(() => {
        if (!isLoading && isInView) {
            setPage((prev) => prev + 1);
        }
    }, [isInView]);

    useEffect(() => {
        if (data?.promos) {
            setCashbacks((prev) => [...prev, ...data.promos]);
        }
    }, [data]);

    return (
        <Page>
            <Container>
                <VStack
                    style={{ gap: 12 }}
                    className={cn("py-8")}
                    align="center"
                >
                    {cashbacks.length <= 0 ? undefined : (
                        <VStack>
                            <h1
                                className={cn(
                                    "text-primary text-3xl font-bold"
                                )}
                            >
                                Available cashbacks
                            </h1>
                            <HStack className="my-4">
                                <span className={cn("text-primary")}>
                                    Общее количество акций в банках:{" "}
                                    {data?.metadata.total_records}
                                </span>
                            </HStack>
                            <HStack>
                                <form onSubmit={searchByBank}>
                                    <Input placeholder="Bank name" />
                                    <Button type="submit">Search</Button>
                                </form>
                            </HStack>
                            <div
                                className={cn(
                                    "grid grid-cols-2 lg:grid-cols-4 gap-4"
                                )}
                            >
                                {cashbacks.map((cashback) => (
                                    <div
                                        key={cashback.ID}
                                        className={cn(
                                            "rounded-lg bg-primary p-4 flex flex-col gap-5 overflow-x-hidden text-balance"
                                        )}
                                    >
                                        {cashback.Title && (
                                            <h3
                                                className={cn(
                                                    "font-bold text-lg"
                                                )}
                                            >
                                                {cashback.Title}
                                            </h3>
                                        )}
                                        {cashback.BonusRate && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Cashback: {cashback.BonusRate}%
                                            </span>
                                        )}
                                        {cashback.BankName && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Bank: {cashback.BankName}
                                            </span>
                                        )}
                                        {cashback.CardType && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Bank`s card type:{" "}
                                                {cashback.CardType}
                                            </span>
                                        )}
                                        {cashback.CompanyName && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Company: {cashback.CompanyName}
                                            </span>
                                        )}
                                        {cashback.Location && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Location: {cashback.Location}
                                            </span>
                                        )}
                                        {cashback.CategoryName && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Category:{" "}
                                                {cashback.CategoryName}
                                            </span>
                                        )}
                                        {cashback.Requirements && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Requirements:{" "}
                                                {cashback.Requirements}
                                            </span>
                                        )}
                                        {cashback.Restrictions && (
                                            <span
                                                className={cn(
                                                    "font-medium text-md"
                                                )}
                                            >
                                                Restrictions:{" "}
                                                {cashback.Restrictions}
                                            </span>
                                        )}
                                        {cashback.SourceUrl && (
                                            <span className=" text-blue-600 hover:underline">
                                                <Link href={cashback.SourceUrl}>
                                                    Check src
                                                </Link>
                                            </span>
                                        )}
                                    </div>
                                ))}
                            </div>
                        </VStack>
                    )}
                    <div ref={ref}>
                        <ImSpinner8 className="my-4 h-10 w-full animate-spin fill-white" />
                    </div>
                </VStack>
            </Container>
        </Page>
    );
}

export default CashbacksPage;
