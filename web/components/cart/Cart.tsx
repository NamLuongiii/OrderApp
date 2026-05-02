'use client'
import {ShoppingCart} from "lucide-react";
import {Dialog} from "@/components/ui/dialog";
import {useEffect, useState} from "react";
import {Button} from "@/components/ui/button";
import {useRouter} from "next/navigation";
import {useCartStore} from "@/components/cart/store/store";

export default function Cart() {
    const router = useRouter();
    const [open, setOpen] = useState(false);
    const [mounted, setMounted] = useState(false);
    const cartItems = useCartStore(state => state.items)
    const total = useCartStore(state => state.getTotalItems)()
    const clearCart = useCartStore(state => state.clearCart)

    useEffect(() => {
        // eslint-disable-next-line react-hooks/set-state-in-effect
        setMounted(true);
    }, []);

    const handleBuy = () => {
        setOpen(false);
        router.push('/order/new')
    }

    return <div className='float-right'>
        <Button variant="outline" size="icon" onClick={() => setOpen(true)} suppressHydrationWarning>
            <ShoppingCart />{mounted ? total : 0}
        </Button>
        <Dialog open={open} onOpenChange={setOpen}>
            <div className='space-y-2'>
                <h1>Giỏ hàng</h1>

                {cartItems.map((item) => (
                    <div key={item.id}>
                        <div>{item.name}</div>
                        <div className='flex justify-between'>
                            <span>{item.formattedPrice}</span>
                            <span>x {item.quantity}</span>
                            <span>{item.formattedTotal}</span>
                        </div>
                    </div>
                ))}

                <div className='flex justify-between'>
                    <Button variant='outline' onClick={clearCart}>Xoá</Button>
                    <Button size='lg' onClick={handleBuy}>Mua</Button>
                </div>
            </div>
        </Dialog>
    </div>
}