import axios from "@/shared/axios/index";

export const createProductApi = async (command: CreateProductRequest) => {
    const res = await axios.post<string>('/product', command);
    return res.data;
}

export interface CreateProductRequest {
    name: string;
    price: number;
    salePrice?: number;
}

