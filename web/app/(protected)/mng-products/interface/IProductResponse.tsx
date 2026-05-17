export default  interface IProductResponse {
    pagination: {
        Page: number;
        PageSize: number;
        PageNums: number
    },
    products?: {
        id: string;
        name: string;
        price: number;
        salePrice?: number;
        finalPrice: number;
    }[]
}
