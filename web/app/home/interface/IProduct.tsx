export default interface IProduct {
    id: string;
    name: string;
    price: string;
    salePrice?: string;
    finalPrice: string;
    formatedPrice: string;
    formatedFinalPrice: string;
}