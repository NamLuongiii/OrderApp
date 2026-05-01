'use client'
import {ShoppingCart} from "lucide-react";
import {Dialog} from "@/components/ui/dialog";
import {useEffect, useState} from "react";
import {Button} from "@/components/ui/button";
import {CartItem} from "@/app/common/component/cart/interface/cart-item";
import getCartItem from "@/app/common/component/cart/api/getCartItem";
import {useRouter} from "next/navigation";
import {useCartStore} from "@/app/common/component/cart/store/store";

export default function Cart() {
    const router = useRouter();
    const [open, setOpen] = useState(false);
    const cartItems = useCartStore(state => state.items)
    const total = useCartStore(state => state.getTotalItems)()


    const handleBuy = () => {
        setOpen(false);
        router.push('/order/new')
    }

    return <div className='float-right'>
        <Button variant="outline" size="icon" onClick={() => setOpen(true)}>
            <ShoppingCart />{total}
        </Button>
        <Dialog open={open} onOpenChange={setOpen}>
            <div>
                <h1>Cart</h1>

                {cartItems.map((item) => (
                    <div key={item.id}>
                        {item.name} - {item.price} - {item.quantity}
                    </div>
                ))}

                <Button size='lg' className='w-full' onClick={handleBuy}>Buy</Button>
            </div>
        </Dialog>
    </div>
}