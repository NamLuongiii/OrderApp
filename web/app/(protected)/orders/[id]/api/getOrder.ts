import axios from "@/shared/axios";
import IOrderResponse from "@/app/(protected)/orders/interface/IOrderResponse";

export default async function getOrder(id: string) {
    const res = await axios.get<IOrderResponse>(`/order/${id}`)
    return res.data
}