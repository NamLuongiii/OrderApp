'use client'
import {Button} from "@/components/ui/button";
import {useCartStore} from "@/components/cart/store/store";
import IProduct from "@/app/home/interface/IProduct";
import Money from "@/shared/class/money";

type Props = {
    product: IProduct
}
export default function Product({product}: Props) {
    const addToCart = useCartStore((state) => state.addToCart);

    const handleAddToCart = (product: IProduct) => {
        addToCart({
            id: product.id,
            name: product.name,
            price: product.finalPrice,
            quantity: 1,
            formattedPrice: product.formatedFinalPrice,
            formattedTotal: product.formatedFinalPrice,
        })

    }

    return (
        <div key={product.id}>
            <div>{product.name}</div>
            {product.salePrice && <div className='text-sm line-through'>{product.formatedPrice}</div>}
            <div>{product.formatedFinalPrice}</div>
            <Button className='w-full' onClick={() => handleAddToCart(product)}>Add to cart</Button>
        </div>
    )
}
