import {useMutation} from "@tanstack/react-query";
import {createProductApi, CreateProductRequest} from "@/shared/axios/createProductApi";

export const useCreate = () => {
    const { isPending, mutateAsync} = useMutation({
        mutationKey: ['create-product'],
        mutationFn: async (command: CreateProductRequest) => {
            return createProductApi(command)
        }
    })

    return {isPending, create: mutateAsync}
}