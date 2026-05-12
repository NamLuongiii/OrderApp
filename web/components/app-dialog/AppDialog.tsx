'use client'
import {Dialog} from "@/components/ui/dialog";
import {useAppStore} from "@/shared/store/app";
import {Button} from "@/components/ui/button";

export default function AppDialog() {
    const command = useAppStore(state => state.appDialogCommand)
    const open = command !== null
    const hideDialog = useAppStore(state => state.hideAppDialog)

    const handleClose = () => {
        hideDialog()
        command?.onClose?.()
    }
    return (
        <Dialog
            open={open}
            onOpenChange={open => !open && handleClose()}
            showCloseButton={false}
        >
            <div className='text-lg text-center'>{command?.message}</div>
            <Button
                className='w-full mt-4'
                size='lg'
                onClick={handleClose}
            >Đóng</Button>
        </Dialog>
    )
}