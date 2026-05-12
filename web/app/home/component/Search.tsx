'use client'
import React, {useState} from "react";
import {useRouter} from "next/navigation";
import {SearchIcon} from "lucide-react";

export default function Search() {
    const router = useRouter();
    const [search, setSearch] = useState('')

    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            if (!search.trim()) return;
            e.preventDefault();
            router.push(`/order/${search.trim()}`)
        }
    }

    const handleSearch = () => {
        if (!search.trim()) return;
        router.push(`/order/${search.trim()}`)
    }

    return <div className='flex items-center gap-2 rounded-full p-1 bg-gray-100 max-w-md mx-auto border focus-within:border-cyan-600 transition-colors'>
        <input
            value={search}
            onChange={e => setSearch(e.target.value)}
            placeholder='Tìm bằng mã đơn'
            onKeyDown={handleKeyDown}
            className='flex-1 outline-none py-2 px-4 text-lg'
        />
        <button className='border p-2 rounded-full bg-cyan-600 text-white cursor-pointer hover:shadow-lg'
            onClick={handleSearch}
        >
            <SearchIcon />
        </button>
    </div>
}
