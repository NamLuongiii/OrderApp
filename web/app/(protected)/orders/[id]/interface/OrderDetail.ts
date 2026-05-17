import {Status} from "@/app/(protected)/orders/interface/status";

export interface IOrderDetail {
    id: string;
    name: string;
    phone: string;
    address: string;
    email: string;
    note?: string;
    status: Status;
    createdAt: string;
    updatedAt: string;
    formatedTotal: string;
    total: number;

    items: {
        productID: string,
        quantity: number,
        price: number,
        total: number,
        itemID: string,
        name: string,
        formattedPrice: string,
        formattedTotal: string,
    }[];
}
