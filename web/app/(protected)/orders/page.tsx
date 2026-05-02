import {getOrders} from "@/app/(protected)/orders/api/getOrders";
import {Button} from "@/components/ui/button";
import Link from "next/link";

export default async function OrdersPage() {
  const orders = await getOrders()

  return (
    <div>
      <div >
        <h1 className="text-3xl font-bold">
          Đơn hàng
        </h1>
        
        <div className='flex flex-col gap-4'>
            Danh sách đơn hàng sẽ hiển thị ở đây

            {orders.map((order) => (
                <Link href={`/orders/${order.id}`} key={order.id}>
                  {order.total} {order.status}
                    <Button >View</Button>
                </Link>
            ))}
        </div>
      </div>
    </div>
  );
}
