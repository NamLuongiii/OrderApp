import {useMutation} from "@tanstack/react-query";
import {LoginCommand} from "@/app/login/interface/loginCommand";
import axios from "@/shared/axios";

export const useLogin = () => {
    return useMutation({
        mutationKey: ['login'],
        mutationFn: async (data: LoginCommand) => {
            const res = await axios.post('/login', data)
            return res.data
        }
    })
}