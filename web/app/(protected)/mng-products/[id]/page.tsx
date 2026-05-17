import Form from "@/app/(protected)/mng-products/[id]/component/Form";
import {getProduct} from "@/app/(protected)/mng-products/[id]/api/getProduct";

export default async function MngProductPage({ params }: { params: { id: string } }) {
    const {id} = await params
    const productForm = await getProduct(id)
    return <div className='space-y-8 py-8'>
        <div className='text-2xl'>Cập nhật sản phẩm</div>
        <Form id={id} product={productForm} />
    </div>
}