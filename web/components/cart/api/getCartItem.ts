import {CartItem} from "@/components/cart/interface/cart-item";

function getCartItem(): CartItem[] {
    return  [
        {
            id: '1',
            name: 'a',
            price: '10000',
            quantity: 1
        }
    ]
}

export default getCartItem