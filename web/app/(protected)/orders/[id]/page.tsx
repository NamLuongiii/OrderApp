import getOrder from "@/app/(protected)/orders/[id]/api/getOrder";
import {Input} from "@/components/ui/input";

export default async function OrderDetailPage({ params }: { params: { id: string } }) {
  const {id} = await params;

  const order = await getOrder(id)

  return (
    <div>
        <h1 className="text-3xl font-bold mb-6">
          Chi tiết đơn hàng
        </h1>

      <div className='space-y-2'>
        {order.line_items.map(lineItem => (
            <div key={lineItem.id}>
              <div>
                {lineItem.product_name}
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
  );
}
