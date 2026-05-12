import axios from "@/shared/axios";
import IOrderResponse from "@/app/(protected)/orders/interface/IOrderResponse";
import IOrder from "@/app/(protected)/orders/interface/IOrder";
import Money from "@/shared/class/money";
import Time from "@/shared/class/time";

export async function getOrders(): Promise<IOrder[]> {
    const res =await  axios.get<IOrderResponse[]>('/orders')
    return res.data.map(order => ({
        ...order,
        formattedTotal: new Money(order.total).format(),
        createdAt: new Time(order.created_at).format(),
        updatedAt: new Time(order.updated_at).format(),
    }))
}