import { Status } from "./status";


export default interface IOrderResponse {
    pagination: {
        Page: number;
        PageSize: number;
        PageNums: number
    };
    orders?: {
        id: string;
        items: {
            productID: string,
            quantity: number,
            price: number,
            total: number,
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
    }[]
}