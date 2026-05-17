'use client'

import { useForm, useWatch } from 'react-hook-form';
import { useEffect } from 'react';
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {useRouter} from "next/navigation";
import {CreateProductRequest} from "@/shared/axios/createProductApi";
import {useCreate} from "@/app/(protected)/mng-products/new/hook/useCreate";
import {useAppStore} from "@/shared/store/app";

// Khai báo kiểu dữ liệu cho Form Inputs
interface IProductInput {
    name: string;
    price: number;
    salePrice?: number;
    finalPrice: number;
}

export default function Form() {
    const {
        register,
        handleSubmit,
        setValue,
        control,
        formState: { errors }
    } = useForm<IProductInput>({
        defaultValues: {
            name: '',
            price: 0,
            salePrice: undefined,
            finalPrice: 0
        }
    });

    const router = useRouter()
    const back = () => router.back()

    const price = useWatch({ control, name: 'price' });
    const salePrice = useWatch({ control, name: 'salePrice' });
    const { isPending, create } = useCreate()
    const showDialog = useAppStore(state => state.showAppDialog)

    useEffect(() => {
        const p = Number(price) || 0;
        const s = Number(salePrice);

        // Nếu có giá sale thì finalPrice = salePrice, ngược lại = price
        const final = (!isNaN(s) && s > 0) ? s : p;
        setValue('finalPrice', final);
    }, [price, salePrice, setValue]);

    const onSubmit = (data: IProductInput) => {
        const command: CreateProductRequest = {
            name: data.name,
            price: Number(data.price),
            salePrice: data.salePrice ? Number(data.salePrice) : undefined
        }

        create(command).then((id: string) => {
            showDialog({
                message: "Tạo sản phẩm thành công",
                onClose: () => {
                    router.push('/mng-products/' + id)
                }
            })
        })
    };

    return (
        <form onSubmit={handleSubmit(onSubmit)} className='space-y-4'>
            {/* Product Name */}
            <div>
                <Input
                    type="text"
                    label='Tên sản phẩm'
                    {...register('name', { required: 'Name is required' })}
                />
                {errors.name && <p className="text-red-500 text-xs mt-1">{errors.name.message}</p>}
            </div>

            {/* Original Price */}
            <div>
                <Input
                    type="number"
                    label="Giá sản phẩm"
                    {...register('price', {
                        required: 'Price is required',
                        min: { value: 0, message: 'Price must be greater than or equal to 0' }
                    })}
                />
                {errors.price && <p className="text-red-500 text-xs mt-1">{errors.price.message}</p>}
            </div>

            {/* Sale Price (Optional) */}
            <div>
                <Input
                    type="number"
                    label="Giá khuyến mãi"
                    {...register('salePrice', {
                        min: { value: 0, message: 'Sale price must be greater than 0' }
                    })}
                />
                {errors.salePrice && <p className="text-red-500 text-xs mt-1">{errors.salePrice.message}</p>}
            </div>

            {/* Final Price (Read-only / Disabled) */}
            <div>
                <Input
                    type="number"
                    label="Giá cuối cùng"
                    disabled
                    {...register('finalPrice')}
                />
            </div>

            <div className="flex gap-4 justify-end">
                <Button variant='outline' onClick={back}>Quay lại</Button>
                <Button disabled={isPending} type="submit">Tạo sản phẩm</Button>
            </div>
        </form>
    );
}