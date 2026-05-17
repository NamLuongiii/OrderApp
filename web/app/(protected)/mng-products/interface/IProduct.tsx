export default interface IProduct {
    id: string;
    name: string;
    price: number;
    salePrice?: number;
    finalPrice: number;
    formatedPrice: string;
    formatedSalePrice?: string;
    formatedFinalPrice: string;
}