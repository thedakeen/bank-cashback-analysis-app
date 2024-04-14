"use client";

import Autoplay from "embla-carousel-autoplay";

import { cn } from "@/lib/utils";

import Container from "@/components/ui/container";
import { HStack } from "@/components/ui/hstack";
import { Page } from "@/components/ui/page";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { VStack } from "@/components/ui/vstack";

import KaspiLogo from "@/public/kaspi-logo.png";
import HalykLogo from "@/public/halyk-bank-logo.png";

import {
    Carousel,
    CarouselContent,
    CarouselItem,
    CarouselNext,
    CarouselPrevious,
} from "@/components/ui/carousel";
import Image from "next/image";

interface HomePageProps {}

const HomePage = ({}: HomePageProps) => {
    return (
        <Page>
            <Container>
                <HStack justify="between" align="center" className={cn("py-4")}>
                    <h2 className={cn("text-lg text-primary font-bold")}>CT</h2>
                    <Button asChild className="bg-secondary px-8">
                        <Link href="/login">Login</Link>
                    </Button>
                </HStack>
            </Container>
            <Container>
                <VStack className={cn("text-center gap-4 py-6")}>
                    <h1 className={cn("text-2xl text-primary font-bold")}>
                        Cashback Tracker - надежный помощник в мире кешбеков!
                    </h1>
                    <span className={cn("text-md text-primary font-medium")}>
                        Наш сайт предоставляет полную информацию о доступных
                        кешбеках, включая процентные ставки, условия получения и
                        особенности программы банков.
                    </span>
                    <span className={cn("text-md text-primary font-medium")}>
                        Вы сможете легко найти наиболее выгодные предложения и
                        выбрать оптимальный вариант.
                    </span>
                    <Carousel
                        opts={{
                            loop: true,
                            align: "center",
                        }}
                        plugins={[
                            Autoplay({
                                delay: 2000,
                            }),
                        ]}
                    >
                        <CarouselContent>
                            <CarouselItem className="basis-1/3">
                                <Image src={KaspiLogo} alt="" />
                            </CarouselItem>
                            <CarouselItem className="basis-1/3">
                                <Image src={HalykLogo} alt="" />
                            </CarouselItem>
                            <CarouselItem className="basis-1/3">
                                <Image src={KaspiLogo} alt="" />
                            </CarouselItem>
                            <CarouselItem className="basis-1/3">
                                <Image src={HalykLogo} alt="" />
                            </CarouselItem>
                            <CarouselItem className="basis-1/3">
                                <Image src={KaspiLogo} alt="" />
                            </CarouselItem>
                            <CarouselItem className="basis-1/3">
                                <Image src={HalykLogo} alt="" />
                            </CarouselItem>
                            <CarouselItem className="basis-1/3">
                                <Image src={KaspiLogo} alt="" />
                            </CarouselItem>
                            <CarouselItem className="basis-1/3">
                                <Image src={HalykLogo} alt="" />
                            </CarouselItem>
                        </CarouselContent>
                    </Carousel>
                </VStack>
            </Container>
        </Page>
    );
};

export default HomePage;
