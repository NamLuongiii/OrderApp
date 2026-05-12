import { Status } from "./status";


export default interface IOrderResponse {
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
    created_at: number;
    updated_at: number;
}