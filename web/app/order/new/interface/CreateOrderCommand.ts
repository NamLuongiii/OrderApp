export interface CreateOrderCommand {
    products: {
        product_id: string;
        quantity: number;
    }[];
    name: string;
    phone: string;
    address: string;
    email: string;
    note?: string;
}