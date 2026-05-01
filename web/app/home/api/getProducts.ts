import axios from "@/app/common/axios";

export default async function getProducts() {
    const res = await axios.get<Product[]>('/product')
    return res.data
}