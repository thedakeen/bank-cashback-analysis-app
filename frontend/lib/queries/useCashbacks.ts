import { useQuery } from "@tanstack/react-query";
import api from "../api";
import { Cashback } from "@/types";

type Data = {
    metadata: object;
    promos: Cashback[];
};

const getCashbacks = async (page: number) => {
    const { data } = await api.get(`/v1/cashbacks?page=${page}`, {
        headers: {
            "X-Auth": localStorage.getItem("token"),
        },
    });
    return data;
};

export const useCashbacks = (page: number = 1) => {
    return useQuery<Data>({
        queryKey: ["cashbacks", page],
        queryFn: () => getCashbacks(page),
    });
};
