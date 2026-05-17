import getOrder from "@/app/(protected)/orders/[id]/api/getOrder";
import {Input} from "@/components/ui/input";
import {StatusAction} from "@/app/(protected)/orders/[id]/component/StatusAction";

export default async function OrderDetailPage({ params }: { params: { id: string } }) {
  const {id} = await params;

  const order = await getOrder(id)

  return (
    <div className='space-y-8 py-8'>
        <h1 className="text-2xl font-bold mb-6">
          Chi tiết đơn hàng
        </h1>


      <table className='space-y-2 w-full border-collapse shadow-md'>
          <thead className='bg-gray-200'>
          <tr className='*:border *:p-2'>
              <th>Tên sản phẩm</th>
              <th>Số lượng</th>
              <th>Giá 1 đơn vị</th>
              <th>Giá tổng</th>
          </tr>
          </thead>
        <tbody>
        {order.items.map(lineItem => (
            <tr key={lineItem.itemID} className='*:border *:p-2'>
                <td>{lineItem.name}</td>
                <td className='text-right'>{lineItem.quantity}</td>
                <td className='text-right'>{lineItem.formattedPrice}</td>
                <td className='text-right'>{lineItem.formattedTotal}</td>
            </tr>
        ))}
        </tbody>
      </table>

        <div className='flex justify-between items-baseline'>
            <span>Tổng tiền đơn hàng</span>
            <span className='text-3xl'>{order.formatedTotal}</span></div>


      <div className='space-y-2'>
        <h2>Thông tin nhận hàng</h2>
        <Input value={order.name} disabled />
        <Input value={order.email} disabled />
        <Input value={order.phone} disabled />
        <Input value={order.address} disabled />
        <Input value={order.note} disabled />
      </div>

      <div>
        {order.status === 'PENDING' && <div>
            <div className='text-2xl'>Vui lòng xác nhận đơn hàng</div>
            <StatusAction orderId={order.id} status='CONFIRMED' label="Xác nhận"></StatusAction>
        </div>}
        {order.status === 'CANCELED' && <div className='text-2xl'>
            Đơn hàng đã bị huỷ
        </div>}
        {order.status === 'CONFIRMED' && <div>
            <div>Đơn hàng đã được xác nhận</div>
            <div className='text-2xl'>Vui lòng chuẩn bị hàng và giao cho đơn vị vận chuyển</div>
            <StatusAction orderId={order.id} status='DELIVERED' label="Tôi đã gửi hàng"></StatusAction>
        </div>}
        {order.status === 'DELIVERED' && <div>
            <div className='text-2xl'>Hàng đang được giao cho khách</div>
            <StatusAction orderId={order.id} status='COMPLETED' label="Đánh dấu đơn hàng đã hoàn thành"></StatusAction>
        </div>}
        {order.status === 'COMPLETED' && <div className='text-2xl'>
            Đơn hàng đã hoàn thành
        </div>}
      </div>
    </div>
  );
}
