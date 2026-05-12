import {Button} from "@/components/ui/button";
import {MinusIcon, PlusIcon} from "lucide-react";

type Props = {
    quantity: number;
    onQuantityChange: (quantity: number) => void;
}

const MAX_QUANTITY = 100

export default function ItemQuantityCalculator({quantity, onQuantityChange}: Props) {
    return <div className='flex items-center'>
        <Button size='icon'
                disabled={quantity <= 1}
                onClick={() => {
            onQuantityChange(quantity - 1)
        }}>
            <MinusIcon/>
        </Button>
        <Button size='icon'>{quantity}</Button>
        <Button size='icon'
                disabled={quantity >= MAX_QUANTITY}
                onClick={() => {
            onQuantityChange(quantity + 1)
        }}>
            <PlusIcon/>
        </Button>
    </div>
}