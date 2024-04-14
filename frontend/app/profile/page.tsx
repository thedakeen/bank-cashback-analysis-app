"use client";

import { Button } from "@/components/ui/button";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Page } from "@/components/ui/page";
import { VStack } from "@/components/ui/vstack";
import { useAddCard } from "@/lib/mutations/useAddCard";
import { UserBankCard } from "@/types";
import { zodResolver } from "@hookform/resolvers/zod";
import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

type Props = {};

function ProfilePage({}: Props) {
    const { mutateAsync: addCardAsync } = useAddCard();
    const form = useForm<z.infer<typeof UserBankCard>>({
        resolver: zodResolver(UserBankCard),
    });

    const onSubmit = async (values: z.infer<typeof UserBankCard>) => {
        // refetch();
        addCardAsync(values);
    };

    return (
        <Page>
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(onSubmit)}
                    className="space-y-4 w-full"
                >
                    <VStack style={{ gap: 12 }}>
                        <FormField
                            control={form.control}
                            name="BankName"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Bank Name</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Kaspi" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="CardNumber"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Card Number</FormLabel>
                                    <FormControl>
                                        <Input
                                            placeholder="1111222233334444"
                                            {...field}
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="CardType"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Card Type</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Gold" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <Button type="submit" className="bg-secondary w-full">
                            Add card
                        </Button>
                    </VStack>
                </form>
            </Form>
        </Page>
    );
}

export default ProfilePage;
