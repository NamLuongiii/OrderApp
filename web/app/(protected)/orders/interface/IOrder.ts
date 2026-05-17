import IOrderResponse from "@/app/(protected)/orders/interface/IOrderResponse";
type RawOrderItem = NonNullable<IOrderResponse['orders']>[number];

export default interface IOrder extends RawOrderItem {
    formattedTotal: string;
    createdAt: string;
    updatedAt: string;
}