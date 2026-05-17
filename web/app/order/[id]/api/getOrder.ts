import axios from "@/shared/axios";
import IOrder from "@/app/(protected)/orders/interface/IOrder";
import IOrderResponse from "@/app/(protected)/orders/interface/IOrderResponse";
import {IOrderDetail} from "@/app/(protected)/orders/[id]/interface/OrderDetail";
import Money from "@/shared/class/money";

type Order = NonNullable<IOrderResponse['orders']>[number]

export default async function getOrder(id: string): Promise<IOrderDetail> {
    const res = await axios.get<Order>(`/order/${id}`)
    const order = res.data
    return {
        id: order.id,
        name: order.name,
        phone: order.phone,
        address: order.address,
        email: order.email,
        note: order.note,
        status: order.status,
        createdAt: order.createdAt,
        updatedAt: order.updatedAt,
        total: order.total,
        formatedTotal: new Money(order.total).format(),

        items: order.items.map(item => ({
            itemID: item.itemID,
            name: item.name,
            price: item.price,
            quantity: item.quantity,
            formattedPrice: new Money(item.price).format(),
            formattedTotal: new Money(item.total).format(),
            productID: item.productID,
            total: item.total,
        }))
    }
}