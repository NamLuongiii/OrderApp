import axios from "@/shared/axios";
import Order from "@/app/(protected)/orders/interface/Order";

export async function getOrders() {
    const res =await  axios.get<Order[]>('/orders')
    return res.data
}