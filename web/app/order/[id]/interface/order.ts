import {Status} from "@/app/order/[id]/interface/status";

export default interface Order {
    id: string;
    line_items: {
        product_id: string,
        quantity: number,
        price: string,
        total: string,
        id: string,
        product_name: string,
    }[];
    total: string;
    name: string;
    phone: string;
    address: string;
    email: string;
    note?: string;
    status: Status;
    created_at: string;
    updated_at: string;
}