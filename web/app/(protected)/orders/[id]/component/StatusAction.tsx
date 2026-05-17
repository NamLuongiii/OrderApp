'use client'

import {Button} from "@/components/ui/button";
import {useChangeStatus} from "@/app/(protected)/orders/[id]/api/changeStatus";
import {useAppStore} from "@/shared/store/app";
import {Status} from "@/app/(protected)/orders/interface/status";

type Props = {
    orderId: string;
    status: Status;
    label: string;
}
export const StatusAction = ({ orderId, status, label }: Props) => {
    const { isPending, changeStatus } = useChangeStatus()
    const showAppDialog = useAppStore(state => state.showAppDialog)

    const handleClick = () => {
        changeStatus({orderId, status}).
        then(() => {
            showAppDialog({
                message: "Cập nhật thành công",
                onClose: () => window.location.reload(),
            })
        }).
        catch(console.error)
    }

    return <Button disabled={isPending} onClick={handleClick}>{label}</Button>
}