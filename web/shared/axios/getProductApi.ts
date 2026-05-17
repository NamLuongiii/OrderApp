import axios from "@/shared/axios/index";

 const getProductApi = (id: string) =>
     axios.get<ProductDetailResponse>('product/' + id);
 export default getProductApi;

interface ProductDetailResponse {
    id: string;
    name: string;
    price: number;
    salePrice?: number;
    finalPrice: number;
}