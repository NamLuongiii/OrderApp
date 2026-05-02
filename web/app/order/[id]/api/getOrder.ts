import axios from "@/shared/axios";
import Order from "@/app/order/[id]/interface/order";

export default async function getOrder(id: string): Promise<Order> {
    const res = await axios.get(`/order/${id}`)
    return res.data
}