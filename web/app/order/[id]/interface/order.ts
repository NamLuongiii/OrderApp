import {Status} from "@/app/order/[id]/interface/status";

export default interface Order {
    id: string;
    items: {
        productID: string,
        quantity: number,
        price: string,
        total: string,
        itemID: string,
        name: string,
    }[];
    total: number;
    name: string;
    phone: string;
    address: string;
    email: string;
    note?: string;
    status: Status;
    createdAt: string;
    updatedAt: string;
}