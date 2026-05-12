'use client'

import { useCartStore } from "@/components/cart/store/store";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {CheckoutFormData, checkoutSchema} from "@/app/order/new/schema/formSchema";
import {CreateOrderCommand} from "@/app/order/new/interface/CreateOrderCommand";
import {useCreateOrder} from "@/app/order/new/api/useCreateOrder";
import {useAppStore} from "@/shared/store/app";
import {useRouter} from "next/navigation";
import Image from "next/image";
import cuteKidImage from "@/public/cute-kid.jpg"
import Link from "next/link";
import ItemQuantityCalculator from "@/app/order/new/component/ItemQuantityCalculator";

export default function CheckoutPage() {
    const cartItems = useCartStore(state => state.items);
    const clearCart = useCartStore(state => state.clearCart);
    const showDialog = useAppStore(state => state.showAppDialog)
    const router = useRouter()
    const cartTotal = useCartStore(state => state.total)
    const updateItemQuantity = useCartStore(state => state.updateItemQuantity)

    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<CheckoutFormData>({
        resolver: zodResolver(checkoutSchema),
    });

    const { mutateAsync, isPending } = useCreateOrder()

    const onSubmit = async (data: CheckoutFormData) => {
        const orderPayload: CreateOrderCommand = {
            products: cartItems.map(item => ({
                product_id: item.id,
                quantity: item.quantity
            })),
            name: data.customerName,
            email: data.email,
            phone: data.phone,
            address: data.address,
            note: data.note,
        };

        try {
            const orderId = await  mutateAsync(orderPayload);

            showDialog({
                message: 'Đặt hàng thành công',
                onClose: () => {
                    router.push(`/router/${orderId}`)
                    clearCart(); // Xóa giỏ hàng sau khi mua
                }
            })

        } catch (error) {
            console.error("Lỗi khi đặt hàng:", error);
        }
    };

    return (
        <div className="p-8 space-y-12">
            <h1 className="text-2xl font-bold">Mua hàng</h1>

            {/* Hiển thị tóm tắt giỏ hàng */}
            <div className="pb-4 space-y-16">
                <table className='w-full border-collapse'>
                    <tbody>
                    {cartItems.map((item) => (
                        <tr key={item.id} className="flex justify-between py-2 border-b">
                            <td className="flex items-center gap-4">
                                <Image
                                    src={cuteKidImage}
                                    alt=""
                                    width={100}
                                    height={80}
                                />
                                <div>
                                    <div className='text-lg font-medium'>{item.name}</div>
                                    <ItemQuantityCalculator
                                        quantity={item.quantity}
                                        onQuantityChange={(newQuantity) => {
                                            updateItemQuantity(item.id, newQuantity)
                                        }}
                                    />
                                </div>
                            </td>
                            <td className="font-medium">{item.formattedTotal}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>

                {/* People Also Bought Products */}
                <div className='space-y-6'>
                    <h2 className='text-xl font-medium'>Người khác cũng mua</h2>
                    <div className='flex'>
                        <div className="flex items-center gap-4 p-2 border shadow-sm">
                            <Image
                                src={cuteKidImage}
                                alt=""
                                width={64}
                                height={64}
                                className="object-cover border aspect-square"
                            />
                            <div>
                                <div>Áo tank-top</div>
                                <div className='text-sm'>100.000đ</div>
                                <Button>Mua</Button>
                            </div>
                        </div>
                    </div>
                </div>

                <div className='flex justify-between items-baseline'>
                    <span>Giá tiền</span>
                    <span className='font-semibold text-2xl'>{cartTotal}</span>
                </div>
            </div>

            {/* Form xử lý chính */}
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-8">
                <h3 className='text-2lg font-medium'>Thông tin người nhận</h3>
                <div>
                    <Input {...register("customerName")} placeholder='Tên người mua' />
                    {errors.customerName && <p className="text-red-500 text-xs mt-1">{errors.customerName.message}</p>}
                </div>

                <div>
                    <Input {...register("email")} placeholder='Địa chỉ email' />
                    {errors.email && <p className="text-red-500 text-xs mt-1">{errors.email.message}</p>}
                </div>

                <div>
                    <Input {...register("phone")} placeholder='Số điện thoại' />
                    {errors.phone && <p className="text-red-500 text-xs mt-1">{errors.phone.message}</p>}
                </div>

                <div>
                    <Input {...register("address")} placeholder='Địa chỉ nhận hàng' />
                    {errors.address && <p className="text-red-500 text-xs mt-1">{errors.address.message}</p>}
                </div>

                <div>
                    <Input {...register("note")} placeholder='Ghi chú (tùy chọn)' />
                </div>

                <div className="flex gap-4 justify-between">
                    <Link href="/"><Button variant='outline'>Quay lại</Button></Link>
                    <Button
                        type="submit"
                        disabled={isSubmitting || cartItems.length === 0 || isPending}
                    >
                        {isSubmitting ? "Đang xử lý..." : "Xác nhận mua"}
                    </Button>
                </div>
            </form>

        </div>
    );
}