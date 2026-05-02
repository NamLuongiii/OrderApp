import getProducts from "@/app/home/api/getProducts";
import Product from "@/app/home/component/Product";
import Search from "@/app/home/component/Search";

export default async function Home() {
    const products = await getProducts()
  return (
    <div className='space-y-4 my-8'>
        <Search />

        <div className='grid grid-cols-3 gap-4 my-8'>
            {products.map((product) => (
                <Product key={product.id} product={product} />
            ))}
        </div>
    </div>
  );
}
