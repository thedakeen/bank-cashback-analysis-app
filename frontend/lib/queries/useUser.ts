import { useQuery } from "@tanstack/react-query";
import api from "../api";

const getUser = async () => {
    const { data } = await api.get("/v1/login", {
        headers: {
            "X-Auth": localStorage.getItem("token"),
        },
    });
    return data;
};

export const useUser = (enabled: boolean = false) => {
    return useQuery({
        queryKey: ["user"],
        queryFn: getUser,
        enabled,
    });
};
