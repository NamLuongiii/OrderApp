import Link from "next/link";
import {getMeApi} from "@/shared/axios/getMeApi";

export default async function ProtectedLayout({children}: {children: React.ReactNode}) {

    const res = await getMeApi()
    const me = res.data

    return <div>
        <div className="flex gap-8 items-center bg-gray-100 p-4 mb-8 font-semibold">
            <Link href='/orders'>
                Đơn hàng
            </Link>
            <Link href='/mng-products'>Sản phẩm</Link>
            <div className='ml-auto text-gray-800'>{me.email}</div>
        </div>
        {children}
    </div>

}