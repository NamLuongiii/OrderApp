import getOrder from "@/app/order/[id]/api/getOrder";
import {Input} from "@/components/ui/input";
import Order from "@/app/order/[id]/interface/order";

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

    return <div className='space-y-4'>
        <h1>Thông tin đơn hàng</h1>
        <div className='space-y-2'>
            {order.items.map(lineItem => (
                <div key={lineItem.itemID}>
                    <div>
                        {lineItem.name}
                        {lineItem.quantity}
                        {lineItem.price}
                        {lineItem.total}
                    </div>
                </div>
            ))}

            <div>Tổng tiền đơn hàng {order.total}</div>
        </div>

        <div className='space-y-2'>
            <h2>Thông tin nhận hàng</h2>
            <Input value={order.name} disabled />
            <Input value={order.email} disabled />
            <Input value={order.phone} disabled />
            <Input value={order.address} disabled />
            <Input value={order.note} disabled />
        </div>

        <div>
            {order.status === 'PROCESSING' && 'Đợi chủ shop xác nhận đơn'}
            {order.status === 'CANCELED' && 'Đơn hàng đã bị huỷ'}
            {order.status === 'CONFIRMED' && 'Chủ shop đã xác nhận và đang chuẩn bị hàng'}
            {order.status === 'DELIVERED' && 'Đơn hàng đã được gửi cho đơn vị vận chuyển'}
            {order.status === 'COMPLETED' && 'Đơn hàng đã hoàn tất'}
        </div>
    </div>
}