export interface CreateOrderCommand {
    items: {
        productId: string;
        quantity: number;
    }[];
    customer: {
        name: string;
        phone: string;
        address: string;
        email: string;
        note?: string;
    }
}