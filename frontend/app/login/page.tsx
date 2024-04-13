"use client";

import { Button } from "@/components/ui/button";
import Container from "@/components/ui/container";
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
import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({
    email: z
        .string()
        .email("Not a valid email")
        .min(6, { message: "Minimum length of email is 6" })
        .max(24, { message: "Maximum length of email is 24" }),
    password: z
        .string()
        .min(6, { message: "Minimum password length is 6" })
        .max(26, { message: "maximum password length os 26" }),
});

function LoginPage() {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            email: "",
            password: "",
        },
    });

    function onSubmit(values: z.infer<typeof formSchema>) {
        console.log(values);
    }

    return (
        <Page>
            <Container>
                <VStack
                    // align="center"
                    className="absolute bg-primary p-12 rounded-2xl w-4/12"
                    style={{
                        top: "50%",
                        left: "50%",
                        transform: "translate(-50%, -50%)",
                    }}
                >
                    <Form {...form}>
                        <form
                            onSubmit={form.handleSubmit(onSubmit)}
                            className="space-y-4 w-full"
                        >
                            <FormField
                                control={form.control}
                                name="email"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Email</FormLabel>
                                        <FormControl>
                                            <Input
                                                placeholder="example@email.com"
                                                {...field}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="password"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Password</FormLabel>
                                        <FormControl>
                                            <Input
                                                placeholder="Write your password"
                                                type="password"
                                                autoComplete="off"
                                                {...field}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <Button
                                type="submit"
                                className="bg-secondary w-full"
                            >
                                Login
                            </Button>
                            <p className="mt-2 text-xs text-center text-gray-700">
                                {" "}
                                Don't have an account?{" "}
                                <span className=" text-blue-600 hover:underline">
                                    <Link href="/signup">Sign up</Link>
                                </span>
                            </p>
                        </form>
                    </Form>
                </VStack>
            </Container>
        </Page>
    );
}

export default LoginPage;
