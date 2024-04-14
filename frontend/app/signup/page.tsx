"use client";

import { Button } from "@/components/ui/button";
import Container from "@/components/ui/container";
import {
    FormField,
    FormItem,
    FormLabel,
    FormControl,
    FormMessage,
    Form,
} from "@/components/ui/form";
import { HStack } from "@/components/ui/hstack";
import { Input } from "@/components/ui/input";
import { Page } from "@/components/ui/page";
import { VStack } from "@/components/ui/vstack";
import api from "@/lib/api";
import { cn } from "@/lib/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { useForm, useFormState } from "react-hook-form";
import { z } from "zod";

interface SignupPageProps {}

const phoneRegex = new RegExp(
    /^([+]?[\s0-9]+)?(\d{3}|[(]?[0-9]+[)])?([-]?[\s]?[0-9])+$/
);
const nameRegex = new RegExp(/^[a-zA-Z]+$/);

const formSchema = z.object({
    email: z
        .string()
        .email("Not a valid email")
        .min(6, { message: "Minimum length of email is 6" })
        .max(32, { message: "Maximum length of email is 32" }),
    password: z
        .string()
        .min(6, { message: "Weak password" })
        .max(26, { message: "Maximum password length os 26" }),
    code: z.string().length(6, { message: "Invalid code" }),
    phone: z.string().regex(phoneRegex, "Invalid Number"),
    name: z
        .string()
        .regex(nameRegex, "Invalid name")
        .min(3, { message: "Name must contain at least 3 " })
        .max(15, { message: "Maximum length of name is 15" }),
    surname: z
        .string()
        .regex(nameRegex, "Invalid name")
        .min(3, { message: "Surname must contain at least 3 " })
        .max(15, { message: "Maximum length of surname is 15" }),
    address: z.string().optional(),
});

enum SignupStage {
    enterEmail = "enterEmail",
    enterCode = "enterCode",
    enterFullData = "enterFullData",
}

const SignupPage: React.FC<SignupPageProps> = ({}) => {
    const router = useRouter();
    const [signupStage, setSignupStage] = useState<SignupStage>(
        SignupStage.enterEmail
    );

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            email: "",
            password: "",
            code: "",
            phone: "",
            name: "",
            surname: "",
            address: "",
        },
    });

    const { errors: emailErrors } = useFormState({
        control: form.control,
        name: "email",
    });
    const { errors: codeErrors } = useFormState({
        control: form.control,
        name: "code",
    });

    const processEmailStage = async () => {
        if (emailErrors.email) {
            return;
        }

        try {
            await api.post("/v1/signup/email", {
                email: form.getValues("email"),
            });
            setSignupStage(SignupStage.enterCode);
        } catch (e) {
            const res = e.response;
            if (res.status === 409) {
                router.push("/login");
            }
        }
    };

    const processCodeStage = async () => {
        if (codeErrors.code) {
            return;
        }

        try {
            const formValues = form.getValues();
            await api.post("/v1/signup/code", {
                email: formValues.email,
                code: formValues.code,
            });
            setSignupStage(SignupStage.enterFullData);
        } catch (e) {
            const res = e.response;
            if (res.status === 401) {
                form.setError("code", { message: "Invalid code" });
            }
        }
    };

    const processCodeFullDataStage = async () => {
        try {
            const formValues = form.getValues();
            const res = await api.post("/v1/signup", {
                name: formValues.name,
                password: formValues.password,
                email: formValues.email,
                // code: form.values.code,
            });
            router.replace("/login");
        } catch (e) {}
    };

    return (
        <Page>
            <Container>
                <VStack
                    className="absolute bg-primary p-12 rounded-2xl w-10/12"
                    style={{
                        top: "50%",
                        left: "50%",
                        transform: "translate(-50%, -50%)",
                    }}
                >
                    <Form {...form}>
                        <form
                            onSubmit={form.handleSubmit(
                                processCodeFullDataStage
                            )}
                            className="space-y-4 w-full"
                        >
                            {signupStage === SignupStage.enterEmail && (
                                <VStack style={{ gap: 12 }}>
                                    <FormField
                                        control={form.control}
                                        name="email"
                                        render={({ field }) => (
                                            <FormItem className={cn("w-full")}>
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
                                    <Button
                                        className={cn("w-full bg-secondary")}
                                        onClick={processEmailStage}
                                    >
                                        Get verification code
                                    </Button>
                                </VStack>
                            )}
                            {signupStage === SignupStage.enterCode && (
                                <VStack style={{ gap: 12 }}>
                                    <FormField
                                        control={form.control}
                                        name="code"
                                        render={({ field }) => (
                                            <FormItem className={cn("w-full")}>
                                                <FormLabel>Code</FormLabel>
                                                <FormControl>
                                                    <Input
                                                        placeholder="Write code from email"
                                                        autoComplete="off"
                                                        {...field}
                                                    />
                                                </FormControl>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />
                                    <Button
                                        className={cn("w-full bg-secondary")}
                                        onClick={processCodeStage}
                                    >
                                        Submit code
                                    </Button>
                                </VStack>
                            )}
                            {signupStage === SignupStage.enterFullData && (
                                <VStack style={{ gap: 12 }}>
                                    <HStack style={{ gap: 10 }}>
                                        <FormField
                                            control={form.control}
                                            name="name"
                                            render={({ field }) => (
                                                <FormItem>
                                                    <FormLabel>
                                                        Your name
                                                    </FormLabel>
                                                    <FormControl>
                                                        <Input
                                                            placeholder="Alex"
                                                            {...field}
                                                        />
                                                    </FormControl>
                                                    <FormMessage />
                                                </FormItem>
                                            )}
                                        />
                                        <FormField
                                            control={form.control}
                                            name="surname"
                                            render={({ field }) => (
                                                <FormItem>
                                                    <FormLabel>
                                                        Your surname
                                                    </FormLabel>
                                                    <FormControl>
                                                        <Input
                                                            placeholder="White"
                                                            {...field}
                                                        />
                                                    </FormControl>
                                                    <FormMessage />
                                                </FormItem>
                                            )}
                                        />
                                    </HStack>
                                    <FormField
                                        control={form.control}
                                        name="phone"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>
                                                    Your phone
                                                </FormLabel>
                                                <FormControl>
                                                    <Input
                                                        placeholder="+77771234567"
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
                                        onClick={processCodeFullDataStage}
                                    >
                                        Create Account
                                    </Button>
                                </VStack>
                            )}
                        </form>
                    </Form>
                </VStack>
            </Container>
        </Page>
    );
};

export default SignupPage;
