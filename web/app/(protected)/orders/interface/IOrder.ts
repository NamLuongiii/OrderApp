import IOrderResponse from "@/app/(protected)/orders/interface/IOrderResponse";

export default interface IOrder extends IOrderResponse {
    formattedTotal: string;
    createdAt: string;
    updatedAt: string;
}