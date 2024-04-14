"use client";

import Container from "@/components/ui/container";
import { Page } from "@/components/ui/page";
import { VStack } from "@/components/ui/vstack";
import { useCashbacks } from "@/lib/queries/useCashbacks";
import { cn } from "@/lib/utils";
import { useRouter } from "next/navigation";
import React from "react";
import Link from "next/link";

type Props = {};

function CashbacksPage({}: Props) {
    const router = useRouter();
    const { data: cashbacks, isLoading, isSuccess } = useCashbacks();

    return (
        <Page>
            <Container>
                <VStack
                    style={{ gap: 12 }}
                    className={cn("py-8")}
                    align="center"
                >
                    <h1 className={cn("text-primary text-3xl font-bold")}>
                        Available cashbacks
                    </h1>
                    <div
                        className={cn("grid grid-cols-2 lg:grid-cols-4 gap-4")}
                    >
                        {isSuccess &&
                            cashbacks.slice(0, 20).map((cashback) => (
                                <div
                                    key={cashback.ID}
                                    className={cn(
                                        "rounded-lg bg-primary p-4 flex flex-col gap-5 overflow-x-hidden text-balance"
                                    )}
                                >
                                    {cashback.Title && (
                                        <h3 className={cn("font-bold text-lg")}>
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
                                            Category: {cashback.CategoryName}
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
            </Container>
        </Page>
    );
}

export default CashbacksPage;
