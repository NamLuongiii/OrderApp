import axios from "@/shared/axios";
import IProductResponse from "@/app/home/interface/IProductResponse";
import IProduct from "@/app/home/interface/IProduct";
import Money from "@/shared/class/money";

export default async function getProducts(): Promise<IProduct[]> {
    const res = await axios.get<IProductResponse[]>('/product')
    return res.data.map(product => ({
        ...product,
        formatedPrice: new Money(product.price).format(),
        formatedFinalPrice: new Money(product.finalPrice).format(),
    }))
}