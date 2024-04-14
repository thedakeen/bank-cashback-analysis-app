import { useQuery } from "@tanstack/react-query";
import api from "../api";
import { Cashback } from "@/types";

const getCashbacks = async () => {
    const { data } = await api.get("/v1/cashbacks", {
        headers: {
            "X-Auth": localStorage.getItem("token"),
        },
    });
    return data.promos;
};

export const useCashbacks = () => {
    return useQuery<Cashback[]>({
        queryKey: ["cashbacks"],
        queryFn: getCashbacks,
    });
};
