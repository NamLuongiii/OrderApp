'use client'
import {Input} from "@/components/ui/input";
import React, {useState} from "react";
import {useRouter} from "next/navigation";

export default function Search() {
    const router = useRouter();
    const [search, setSearch] = useState('')
    const handleSearch = () => {
        router.push(`/orders`)
    }

    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === 'Enter') {
            e.preventDefault();
            handleSearch();
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
