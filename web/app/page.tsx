import getProducts from "@/app/home/api/getProducts";
import {Button} from "@/components/ui/button";
import Product from "@/app/home/component/Product";
import { Input } from "@/components/ui/input";

export default async function Home() {
    const products = await getProducts()
    console.log(products)
  return (
    <div className='grid grid-cols-3 gap-4 my-8'>
        <form className='col-span-3'>
            <Input placeholder='Find order by ID' />
        </form>
        {products.map((product) => (
            <Product key={product.id} product={product} />
        ))}
    </div>
  );
}
