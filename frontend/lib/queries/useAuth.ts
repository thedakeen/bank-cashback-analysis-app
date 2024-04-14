import { useQuery } from "@tanstack/react-query";
import api from "../api";

const authUser = async (email: string, password: string) => {
    const res = await api.post("/v1/login", {
        Email: email,
        Password: password,
    });
    localStorage.setItem("token", res.data.token);
    return res.data;
};

export const useAuth = (enabled: boolean, email: string, password: string) => {
    return useQuery({
        queryKey: ["auth", email, password],
        queryFn: () => authUser(email, password),
        enabled,
    });
};
