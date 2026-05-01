import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';
import { CartItem } from '../interface/cart-item';

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
                    set({
                        items: currentItems.map((item) =>
                            item.id === newItem.id
                                ? { ...item, quantity: item.quantity + newItem.quantity }
                                : item
                        ),
                    });
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