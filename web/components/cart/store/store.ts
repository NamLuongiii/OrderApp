import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';
import { CartItem } from '../interface/cart-item';
import Money from "@/shared/class/money";

interface CartState {
    items: CartItem[];
    addToCart: (item: CartItem) => void;
    clearCart: () => void;
    getTotalItems: () => number;
    newAddedItems: CartItem[];
    popupNewAddedItem: (item: CartItem) => void;
    closeNewAddedItemPopup: () => void;
    total: string;
    calculateTotal: () => void;
    updateItemQuantity: (id: string, quantity: number) => void;
}

export const useCartStore = create<CartState>()(
        (set, get) => ({
            items: [],
            newAddedItems: [],
            total: '-',

            addToCart: (newItem: CartItem) => {
                const popupNewAddedItem = get().popupNewAddedItem;
                popupNewAddedItem(newItem);

                const currentItems = get().items;
                const existingItem = currentItems.find((item) => item.id === newItem.id);

                if (existingItem) {
                    const newItems = [...currentItems];
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

                const calculateTotal = get().calculateTotal;
                calculateTotal();
            },

            clearCart: () => set({ items: [] }),

            getTotalItems: () => {
                return get().items.reduce((total, item) => total + item.quantity, 0);
            },
            popupNewAddedItem: (item: CartItem) => {
                const currentItems = get().newAddedItems;
                set({ newAddedItems: [...currentItems, item] })
            },
            closeNewAddedItemPopup: () => set({ newAddedItems: [] }),
            calculateTotal: () => {
                const cartItems = get().items;
                const initialCartTotal = new Money('0');

                const cartTotalAccumulator = cartItems.reduce((acc: Money, item: CartItem) => {
                    const itemTotal = new Money(item.price).multiply(item.quantity.toString());
                    return acc.plus(itemTotal.toString())
                }, initialCartTotal);

                set({ total: cartTotalAccumulator.format() })
            },
            updateItemQuantity: (id: string, quantity: number) => {
                const currentItems = get().items;
                const existingItem = currentItems.find((item) => item.id === id);
                if (existingItem) {
                    const newItems = [...currentItems];
                    const existingItemIdx = newItems.findIndex((item) => item.id === id);
                    const newQuantity = quantity;
                    newItems[existingItemIdx] = {
                        ...existingItem,
                        quantity: newQuantity,
                        formattedTotal: new Money(existingItem.price).multiply(String(newQuantity)).format()
                    }
                    set({ items: newItems });
                }

                const calculateTotal = get().calculateTotal;
                calculateTotal();
            }
        }),
);