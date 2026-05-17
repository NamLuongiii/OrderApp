import axios from "@/shared/axios/index";

export const getMeApi = async () =>
    await axios.get('/me');

export interface MeResponse {
    id: string;
    email: string;
    name: string;
    role: string;
}