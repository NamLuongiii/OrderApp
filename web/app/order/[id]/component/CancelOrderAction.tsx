'use client'

import {Button} from "@/components/ui/button";
import {useState} from "react";
import {useChangeStatus} from "@/app/order/[id]/api/changeStatus";
import {useAppStore} from "@/shared/store/app";
import {Dialog} from "@/components/ui/dialog";

export const CancelOrderAction = ({orderId}: {orderId: string}) => {
    const [open, setOpen] = useState(false)
    const showDialog = useAppStore(state => state.showAppDialog)
    const { isPending, changeStatus } = useChangeStatus()
    const cancel = async () => {
        changeStatus({ orderId, status: 'CANCELED' }).then(() => {
            showDialog({
                message: 'Huỷ đơn hàng thành công',
                onClose: () => {
                    setOpen(false)
                    window.location.reload()
                }
            })
        }).catch(error => {
            alert(error)
        })
    }

    return <div>
        <Button onClick={() => setOpen(true)}>Huỷ đơn hàng</Button>
        <Dialog open={open} onOpenChange={setOpen}>
            <div className='space-y-4'>
                <h1 className='text-2xl'>Huỷ đơn hàng?</h1>
                <div className='flex justify-between items-center gap-4'>
                    <Button variant='outline' onClick={() => setOpen(false)}>Quay lại</Button>
                    <Button disabled={isPending} onClick={cancel}>Xác nhận huỷ</Button>
                </div>
            </div>
        </Dialog>
    </div>
}