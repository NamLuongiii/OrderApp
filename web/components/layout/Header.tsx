import Cart from "@/components/cart/Cart";
import Link from "next/link";

export default function Header() {
    return <header className='text-center text-xl mb-6'>
        <Link href='/' className=' bg-yellow-400 p-2 inline-block'>Order App</Link>
        <Cart />
    </header>
}