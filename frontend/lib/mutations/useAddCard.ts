import { UserBankCard } from "@/types";
import { useMutation } from "@tanstack/react-query";
import api from "../api";

const addUserBankCard = async (card: UserBankCard) => {
    const {} = await api.post(
        "/v1/card",
        {
            card_number: card.CardNumber,
            card_type: card.CardType,
            bank_name: card.BankName,
        },
        {
            headers: {
                "X-Auth": localStorage.getItem("token"),
            },
        }
    );
};

export const useAddCard = () => {
    return useMutation({
        mutationKey: ["add_card"],
        mutationFn: addUserBankCard,
    });
};
