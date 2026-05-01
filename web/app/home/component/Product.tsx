'use client'
import {Button} from "@/components/ui/button";
import {useCartStore} from "@/app/common/component/cart/store/store";

type Props = {
    product: Product
}
export default function Product({product}: Props) {
    const addToCart = useCartStore((state) => state.addToCart);

    return (
        <div key={product.id}>
            <div>{product.name}</div>
            <div>{product.price} vnđ</div>
            <Button className='w-full' onClick={() => addToCart({
                id: product.id,
                name: product.name,
                price: product.price,
                quantity: 1
            })}>Add to cart</Button>
        </div>
    )
}
