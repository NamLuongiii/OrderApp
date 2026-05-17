import getOrder from "@/app/order/[id]/api/getOrder";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import Image from "next/image";
import cuteKidImage from "@/public/cute-kid.jpg"
import {CancelOrderAction} from "@/app/order/[id]/component/CancelOrderAction";

export default async function Page({ params }: { params: { id: string } }) {
    const {id} = await params
    let order = null
    let error = null

    try {
        order = await getOrder(id)
    } catch {
        error = "Có lỗi xảy ra"
    }

    if (error || !order) return <div>{error}</div>

    return <div className='space-y-8 py-8'>
        <div>
            <h1 className='text-2xl'>Thông tin đơn hàng</h1>
            <div className='text-lg font-semibold'>#{order.id}</div>
        </div>
        <table className='w-full border-collapse'>
            <tbody>
            {order.items.map(lineItem => (
                <tr key={lineItem.itemID}>
                    <td>
                        <Image src={cuteKidImage} alt="" width={100} height={100} className='object-cover aspect-square' />
                    </td>
                    <td>{lineItem.name}</td>
                    <td className='text-right'>x{lineItem.quantity}</td>
                    <td className='text-right'>{lineItem.formattedTotal}</td>
                </tr>
            ))}
            </tbody>
        </table>

        <div className='flex justify-between items-baseline'>
            <span>Tổng tiền đơn hàng</span>
            <span className='text-2xl font-semibold'>{order.total}</span>
        </div>

        <div className='space-y-4'>
            <h2 className='text-xl'>Thông tin nhận hàng</h2>
            <Input value={order.name} disabled />
            <Input value={order.email} disabled />
            <Input value={order.phone} disabled />
            <Input value={order.address} disabled />
            <Input value={order.note} disabled />
        </div>

        <div className='text-xl font-semibold'>
            {order.status === 'PENDING' && 'Đợi chủ shop xác nhận đơn'}
            {order.status === 'CANCELED' && 'Đơn hàng đã bị huỷ'}
            {order.status === 'CONFIRMED' && 'Chủ shop đã xác nhận và đang chuẩn bị hàng'}
            {order.status === 'DELIVERED' && 'Đơn hàng đã được gửi cho đơn vị vận chuyển'}
            {order.status === 'COMPLETED' && 'Đơn hàng đã hoàn tất'}
        </div>

        {order.status !== 'CANCELED' && <div>
            <CancelOrderAction orderId={id} />
        </div>}
    </div>
}