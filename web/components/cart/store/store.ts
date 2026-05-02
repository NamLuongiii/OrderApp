import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';
import { CartItem } from '../interface/cart-item';
import Money from "@/shared/class/money";

interface CartState {
    items: CartItem[];
    addToCart: (item: CartItem) => void;
    clearCart: () => void;
    getTotalItems: () => number;
}

export const useCartStore = create<CartState>()(
    persist(
        (set, get) => ({
            items: [],

            addToCart: (newItem: CartItem) => {
                const currentItems = get().items;
                const existingItem = currentItems.find((item) => item.id === newItem.id);

                if (existingItem) {
                    const newItems = currentItems;
                    const existingItemIdx = newItems.findIndex((item) => item.id === newItem.id);
                    const newQuantity = existingItem.quantity + newItem.quantity;

                    newItems[existingItemIdx] = {
                        ...existingItem,
                        quantity: newQuantity,
                        formattedTotal: new Money(existingItem.price).multiply(String(newQuantity)).format()
                    }
                    set({ items: newItems });
                } else {
                    set({ items: [...currentItems, newItem] });
                }
            },

            clearCart: () => set({ items: [] }),

            getTotalItems: () => {
                return get().items.reduce((total, item) => total + item.quantity, 0);
            },
        }),
        {
            name: 'cart-storage',
            storage: createJSONStorage(() => localStorage),
        }
    )
);