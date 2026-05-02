'use client'
import {Input} from "@/components/ui/input";
import React, {useState} from "react";
import {useRouter} from "next/navigation";

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

    return <div>
        <Input
            value={search}
            onChange={e => setSearch(e.target.value)}
            placeholder='Tìm kiếm đơn hàng bằng ID'
            onKeyDown={handleKeyDown}
        />
    </div>
}
