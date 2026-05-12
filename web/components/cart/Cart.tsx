'use client'
import {ShoppingCart, XIcon} from "lucide-react";
import {Dialog} from "@/components/ui/dialog";
import {useEffect, useState} from "react";
import {Button} from "@/components/ui/button";
import {useRouter} from "next/navigation";
import {useCartStore} from "@/components/cart/store/store";
import Image from "next/image";
import cuteKidImage from "@/public/cute-kid.jpg"

export default function Cart() {
    const router = useRouter();
    const [open, setOpen] = useState(false);
    const [mounted, setMounted] = useState(false);
    const cartItems = useCartStore(state => state.items)
    const clearCart = useCartStore(state => state.clearCart)
    const newAddedItems = useCartStore(state => state.newAddedItems)
    const closeNewAddedItemPopup = useCartStore(state => state.closeNewAddedItemPopup)
    const openNewAddedItemPopup = newAddedItems.length > 0
    const total = useCartStore(state => state.total)

    useEffect(() => {
        // eslint-disable-next-line react-hooks/set-state-in-effect
        setMounted(true);
    }, []);

    const handleBuy = () => {
        setOpen(false);
        router.push('/order/new')
    }

    if (!mounted) return null;

    return <div className='relative float-right'>
        <Button
            size="icon"
            onClick={() => setOpen(true)}
            suppressHydrationWarning>
            <ShoppingCart />
        </Button>
        <Dialog open={open} onOpenChange={setOpen}>
            {cartItems.length == 0 ? (
                <div>
                    <div className='text-center font-semibold'>Chưa có sản phẩm nào</div>
                </div>
            ): (
                <div>
                    <div className='space-y-2'>
                        <h1 className='font-medium text-lg'>Giỏ hàng</h1>
                        <div className='space-y-6'>
                                {cartItems.map((item) => (
                                    <div
                                        key={item.id}
                                        className='flex items-center gap-4 py-2 border-b hover:shadow-md transition-shadow'>
                                        <Image
                                            src={cuteKidImage}
                                            alt=''
                                            width={80}
                                            height={80}
                                            className='object-cover'
                                        />
                                        <div className='flex-1'>
                                            <div className='text-lg font-semibold'>{item.name}</div>
                                            <div className='flex justify-between'>
                                                <span>{item.formattedPrice}</span>
                                                <span>x {item.quantity}</span>
                                            </div>
                                        </div>
                                    </div>
                                ))}

                                <div className='flex justify-between items-baseline'>
                                    <span>Tổng</span>
                                    <span className='text-2xl font-medium'>{total}</span>
                                </div>

                                <div className='flex justify-between'>
                                    <Button variant='outline' onClick={clearCart}>Xoá</Button>
                                    <Button size='lg' onClick={handleBuy}>Đặt hàng</Button>
                                </div>
                            </div>
                    </div>
                </div>
            )}
        </Dialog>

        {openNewAddedItemPopup && (
            <div className='absolute top-10 right-0 bg-white p-4 shadow-xl border z-10 w-64 text-sm text-left'>
                <div>Đã thêm vào giỏ</div>
                {newAddedItems.map((item, index) => (
                    <div key={index} className='flex items-center gap-2 py-2 border-b'>
                        <Image
                            className='object-cover'
                            src={cuteKidImage}
                            alt=""
                            width={50}
                            height={50}
                        />
                        <div>
                            <div className='font-semibold'>{item.name}</div>
                            <div className='flex justify-between'>
                                <span>{item.formattedPrice}</span>
                            </div>
                        </div>
                    </div>
                ))}
                <Button
                    className='absolute top-2 right-2'
                    variant='outline'
                    size='icon'
                    onClick={closeNewAddedItemPopup}>
                    <XIcon />
                </Button>
            </div>
        )}
    </div>
}