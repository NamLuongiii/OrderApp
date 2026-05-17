import getProducts from "@/app/(protected)/mng-products/api/getProducts";
import {Button} from "@/components/ui/button";
import Link from "next/link";

export default async function ProductsPage() {
    const products = await getProducts()

  return (
    <div className="py-8 space-y-8">
        <div>
            <Link href="/mng-products/new"><Button>Tạo sản phẩm</Button></Link>
        </div>
        <table className="w-full border-collapse shadow-md">
            <thead>
            <tr className="*:border *:p-2 text-left bg-gray-200">
                <th>Tên sản phẩm</th>
                <th className='text-right'>Giá</th>
                <th className='text-right'>Giảm giá</th>
                <th>Thêm</th>
            </tr>
            </thead>
            <tbody>
            {products.map((product) => (
                <tr key={product.id} className="*:border *:p-4 hover:bg-gray-100 transition-colors">
                    <td>{product.name}</td>
                    <td className='text-right'>{product.formatedPrice}</td>
                    <td className='text-right'>{product.formatedSalePrice}</td>
                    <td>
                        <Link href={`/mng-products/${product.id}`}><Button>Sửa</Button></Link>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    </div>
  );
}
