"use client"

import * as React from "react"
import { Dialog as DialogPrimitive } from "@base-ui/react/dialog"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { XIcon } from "lucide-react"

interface DialogProps extends DialogPrimitive.Root.Props {
  trigger?: React.ReactNode
  title?: string
  description?: string
  children: React.ReactNode
  showCloseButton?: boolean
}

export function Dialog({ trigger, title, description, children, showCloseButton = true, ...props }: DialogProps) {
  return (
    <DialogPrimitive.Root {...props}>
      {trigger && <DialogPrimitive.Trigger>{trigger}</DialogPrimitive.Trigger>}
      <DialogPrimitive.Portal>
        <DialogPrimitive.Backdrop className="fixed inset-0 z-50 bg-black/50" />
        <DialogPrimitive.Popup className="fixed top-1/2 left-1/2 z-50 w-full max-w-lg -translate-x-1/2 -translate-y-1/2 rounded-lg bg-white p-6 shadow-lg dark:bg-zinc-900">
          {title && <DialogPrimitive.Title className="text-lg font-semibold mb-2">{title}</DialogPrimitive.Title>}
          {description && <DialogPrimitive.Description className="text-sm text-zinc-600 dark:text-zinc-400 mb-4">{description}</DialogPrimitive.Description>}
          <div>{children}</div>
          {showCloseButton && (
            <DialogPrimitive.Close render={<Button variant="ghost" className="absolute top-2 right-2" size="icon-sm" />}>
              <XIcon />
            </DialogPrimitive.Close>
          )}
        </DialogPrimitive.Popup>
      </DialogPrimitive.Portal>
    </DialogPrimitive.Root>
  )
}
