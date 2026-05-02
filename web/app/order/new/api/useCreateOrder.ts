import {useMutation} from "@tanstack/react-query";
import {CreateOrderCommand} from "@/app/order/new/interface/CreateOrderCommand";
import axios from "@/shared/axios";

export const useCreateOrder = () => {
    return useMutation({
        mutationKey: ['create-order'],
        mutationFn: async (data: CreateOrderCommand) => {
            const res = await axios.post('/order', data)
            return res.data
        }
    })
}