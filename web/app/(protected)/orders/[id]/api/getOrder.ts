import axios from "@/shared/axios";
import Order from "@/app/(protected)/orders/interface/Order";

export default async function getOrder(id: string) {
    const res = await axios.get<Order>(`/order/${id}`)
    return res.data
}