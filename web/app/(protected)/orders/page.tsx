import {getOrders} from "@/app/(protected)/orders/api/getOrders";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import OrderInformationInMonth from "@/app/(protected)/orders/component/OrderInformationInMonth";

export default async function OrdersPage() {
  const orders = await getOrders()

  return (
    <div>
      <div className='space-y-8'>

          <OrderInformationInMonth />


          <div className="flex gap-8 items-center bg-gray-100 p-4">
              <Link href='/orders' className="font-semibold p-2 px-4 rounded-lg">
                  Đơn hàng
              </Link>
              <Link href='/products'>Sản phẩm</Link>
          </div>
        
        <table className="w-full border border-collapse shadow-md">
            <thead className='bg-gray-200'>
                <tr className="border-b *:text-left *:p-4">
                    <th>Sản phẩm</th>
                    <th>Giá trị</th>
                    <th>Trạng thái</th>
                    <th>Ngày tạo đơn</th>
                    <th>Chi tiết</th>
                </tr>
            </thead>
            <tbody>
                    {orders.map((order) => (
                        <tr key={order.id} className='border-b-2 *:p-4'>
                            <td>
                                <ol className="pl-4 list-disc">
                                    {order.line_items.map(lineItem => (
                                        <li key={lineItem.id}>
                                            {lineItem.product_name} x {lineItem.quantity}
                                        </li>
                                    ))}
                                </ol>
                            </td>
                            <td className='text-right'>{order.formattedTotal}</td>
                            <td>{order.status}</td>
                            <td>{order.createdAt}</td>
                            <td>
                                <Link href={`/orders/${order.id}`}>
                                    <Button>Chi tiết</Button>
                                </Link>
                            </td>
                        </tr>
                    ))}
            </tbody>
        </table>
      </div>
    </div>
  );
}
