'use client'
import {Button} from "@/components/ui/button";
import {useCartStore} from "@/components/cart/store/store";
import IProduct from "@/app/home/interface/IProduct";
import Image from "next/image"
import cuteKidImage from "@/public/cute-kid.jpg"

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
        <div key={product.id} className='bg-white border shadow-md hover:shadow-lg flex flex-col'>
            <Image src={cuteKidImage} alt="" />
            <div className='p-4 flex-1'>
                <div className='font-semibold'>{product.name}</div>
                {product.salePrice && <div className='text-sm line-through'>{product.formatedPrice}</div>}
                <div>{product.formatedFinalPrice}</div>
            </div>
            <Button
                className='w-full rounded-none'
                onClick={() => handleAddToCart(product)}>Thêm vào giỏ
            </Button>
        </div>
    )
}
