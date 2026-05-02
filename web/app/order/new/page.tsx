'use client'

import { useCartStore } from "@/components/cart/store/store";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {CheckoutFormData, checkoutSchema} from "@/app/order/new/schema/formSchema";
import {CreateOrderCommand} from "@/app/order/new/interface/CreateOrderCommand";
import {useCreateOrder} from "@/app/order/new/api/useCreateOrder";
import Decimal from "decimal.js";

export default function CheckoutPage() {
    const cartItems = useCartStore(state => state.items);
    const clearCart = useCartStore(state => state.clearCart);

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
            await  mutateAsync(orderPayload);
            clearCart(); // Xóa giỏ hàng sau khi mua
        } catch (error) {
            console.error("Lỗi khi đặt hàng:", error);
        }
    };

    return (
        <div className="max-w-2xl mx-auto p-4">
            <h1 className="text-2xl font-bold mb-4">Mua hàng</h1>

            {/* Hiển thị tóm tắt giỏ hàng */}
            <div className="mb-6 border-b pb-4">
                {cartItems.map((item) => (
                    <div key={item.id} className="flex justify-between py-1">
                        <span>{item.name} x {item.quantity}</span>
                        <span className="font-medium">{item.price} VND</span>
                    </div>
                ))}

                <div>Giá tiền </div>
            </div>

            {/* Form xử lý chính */}
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
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

                <Button
                    type="submit"
                    className='w-full'
                    disabled={isSubmitting || cartItems.length === 0 || isPending}
                >
                    {isSubmitting ? "Đang xử lý..." : "Xác nhận mua"}
                </Button>
            </form>
        </div>
    );
}