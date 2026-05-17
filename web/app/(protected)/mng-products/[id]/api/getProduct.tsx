import getProductApi from "@/shared/axios/getProductApi";
import IProductInput from "@/app/(protected)/mng-products/[id]/interface/ProductForm";

export async function getProduct(id: string): Promise<IProductInput> {
    const res = await getProductApi(id)
    const prod = res.data
    return {
        name: prod.name,
        price: prod.price,
        salePrice: prod.salePrice,
        finalPrice: prod.finalPrice
    }
}