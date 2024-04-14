import { z } from "zod";

export const CashbacksType = z.enum([
    "Company",
    "Promo",
    "Location",
    "Category",
]);

export const Cashback = z.object({
    ID: z.string(),
    Title: z.string().optional(),
    SourceUrl: z.string().optional(),
    BankName: z.string().optional(),
    CompanyName: z.string().optional(),
    CategoryName: z.string().optional(),
    BonusRate: z.string().optional(),
    Requirements: z.string().optional(),
    Restrictions: z.string().optional(),
    Location: z.string().optional(),
    Type: CashbacksType,
    CardType: z.string(),
});

export const UserBankCard = z.object({
    CardNumber: z
        .string()
        .length(16)
        .regex(/^[0-9]+$/),
    CardType: z.string(),
    BankName: z.string(),
});

export type Cashback = z.infer<typeof Cashback>;
export type CashbacksType = z.infer<typeof CashbacksType>;
export type UserBankCard = z.infer<typeof UserBankCard>;
