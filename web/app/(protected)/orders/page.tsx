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
                                    {order.items.map(lineItem => (
                                        <li key={lineItem.itemID}>
                                            {lineItem.name} x {lineItem.quantity}
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
