'use client'

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import Cookies from "js-cookie";
import { useRouter, useSearchParams } from "next/navigation";
import {useLogin} from "@/app/login/api/useLogin";

const loginSchema = z.object({
    email: z.string().email("Email không hợp lệ"),
    password: z.string().min(4, "Mật khẩu phải từ 4 ký tự trở lên"),
});

type LoginFormData = z.infer<typeof loginSchema>;

export default function LoginPage() {
    const router = useRouter();
    const searchParams = useSearchParams();

    const callbackUrl = searchParams.get('from') || '/orders';

    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
    } = useForm<LoginFormData>({
        resolver: zodResolver(loginSchema),
    });

    const { mutateAsync, isPending } = useLogin()

    const onSubmit = async (data: LoginFormData) => {
        try {
            const token = await mutateAsync(data)

            Cookies.set('auth-token', token, { expires: 12 / 24 });

            router.push(callbackUrl);
            router.refresh();
        } catch (error) {
            console.error("Lỗi đăng nhập:", error);
        }
    };

    return (
        <div className="max-w-md mx-auto mt-10 p-6 border rounded-lg shadow-sm">
            <h1 className="text-2xl font-bold mb-6 text-center">Đăng nhập</h1>

            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
                <div>
                    <Input
                        {...register("email")}
                        type='text'
                        placeholder='Email'
                    />
                    {errors.email && <p className="text-red-500 text-xs mt-1">{errors.email.message}</p>}
                </div>

                <div>
                    <Input
                        {...register("password")}
                        type='password'
                        placeholder='Mật khẩu'
                    />
                    {errors.password && <p className="text-red-500 text-xs mt-1">{errors.password.message}</p>}
                </div>

                <Button
                    type="submit"
                    className='w-full'
                    disabled={isSubmitting || isPending}
                >
                    {isSubmitting ? "Đang xử lý..." : "Đăng nhập"}
                </Button>
            </form>
        </div>
    )
}