import Cart from "@/app/common/component/cart/Cart";
import Link from "next/link";

export default function Header() {
    return <header className='text-center text-xl'>
        <Link href='/' className=' bg-yellow-400 p-2'>Order App</Link>
        <Cart />
    </header>
}