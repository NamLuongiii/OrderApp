import axios from "@/shared/axios";
import Money from "@/shared/class/money";
import IProduct from "@/app/(protected)/mng-products/interface/IProduct";
import IProductResponse from "@/app/(protected)/mng-products/interface/IProductResponse";

export default async function getProducts(): Promise<IProduct[]> {
    const res = await axios.get<IProductResponse>('/products')
    return res.data.products?.map(product => ({
        ...product,
        formatedPrice: new Money(product.price).format(),
        formatedSalePrice: product.salePrice ? new Money(product.salePrice).format() : "",
        formatedFinalPrice: new Money(product.finalPrice).format(),
    })) ?? []
}