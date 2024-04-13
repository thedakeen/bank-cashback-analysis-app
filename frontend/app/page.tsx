import { CiCreditCard1 } from "react-icons/ci";

import { cn } from "@/lib/utils";

import Container from "@/components/ui/container";
import { HStack } from "@/components/ui/hstack";
import { Page } from "@/components/ui/page";
import { Button } from "@/components/ui/button";
import Link from "next/link";

interface HomePageProps {}

const HomePage = ({}: HomePageProps) => {
    return (
        <Page>
            <Container>
                <HStack justify="between" align="center" className="p-16">
                    <HStack align="center" style={{ gap: 10 }}>
                        <CiCreditCard1
                            // @ts-ignore
                            className={cn("fill-secondary h-20 w-20")}
                        />
                        <h2 className={cn("text-md text-primary")}>
                            Cashbacks Master
                        </h2>
                    </HStack>
                    <Button asChild className="bg-secondary px-8">
                        <Link href="/login">Login</Link>
                    </Button>
                </HStack>
            </Container>
        </Page>
    );
};

export default HomePage;
