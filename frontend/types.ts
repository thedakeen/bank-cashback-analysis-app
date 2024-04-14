import { z } from "zod";

const CashbacksType = z.enum(["Company", "Promo", "Location", "Category"]);

const Cashback = z.object({
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

export type Cashback = z.infer<typeof Cashback>;
export type CashbacksType = z.infer<typeof CashbacksType>;
