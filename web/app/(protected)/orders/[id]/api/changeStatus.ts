import {useMutation} from "@tanstack/react-query";
import axios from "@/shared/axios";

export const useChangeStatus = () => {
    const {isPending, mutateAsync} = useMutation({
        mutationKey: ['change-status'],
        mutationFn: async ({orderId, status}: {orderId: string, status: string}) => {
            await  axios.put(`/order/${orderId}?status=${status}`)
            return true;
        }
    })

    return {isPending, changeStatus: mutateAsync}
}