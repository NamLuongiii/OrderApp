import {create} from "zustand";

type TAppDialogCommand = {
    message: string;
    onClose?: () => void;
}

type AppStore = {
    isLoggedIn: boolean;
    appDialogCommand: TAppDialogCommand | null;
    showAppDialog(command: TAppDialogCommand): void;
    hideAppDialog(): void;
}

export const useAppStore = create<AppStore>((set) => ({
    isLoggedIn: false,
    appDialogCommand: null,

    showAppDialog(command: TAppDialogCommand) {
        set({ appDialogCommand: command })
    },
    hideAppDialog() {
        set({ appDialogCommand: null })
    },
}))