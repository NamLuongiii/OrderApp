
import axios from 'axios';
import Cookies from 'js-cookie';

const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json',
    },
});

api.interceptors.request.use(
    async (config) => {
        let token;

        if (typeof window !== 'undefined') {
            token = Cookies.get('auth-token');
        }
        else {
            try {
                const { cookies } = await import('next/headers');
                const cookieStore = await cookies();
                token = cookieStore.get('auth-token')?.value;
            } catch (e) {
                console.error("Không thể đọc cookie trên Server:", e);
            }
        }

        if (token) {
            config.headers['Authorization'] = token
        }

        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response && error.response.status === 401) {
            if (typeof window !== 'undefined') {
                Cookies.remove('auth-token');
                window.location.href = '/login';
            }
        }
        console.error("Error Response:", error.response);
        return Promise.reject(error);
    }
);

export default api;