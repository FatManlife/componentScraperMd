import axios from "axios";

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL || "http://localhost:8080",
    timeout: Number(import.meta.env.VITE_API_TIMEOUT) || 10000,
    headers: {
        "Content-Type": "application/json",
    },
});

// Request interceptor
api.interceptors.request.use(
    (config) => {
        // You can add auth tokens here if needed
        // const token = localStorage.getItem('token');
        // if (token) {
        //   config.headers.Authorization = `Bearer ${token}`;
        // }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    },
);

// Response interceptor
api.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        // Handle errors globally
        if (error.response) {
            // Server responded with error
            if (import.meta.env.VITE_ENV === "development") {
                console.error("API Error:", error.response.data);
            }
        } else if (error.request) {
            // No response received
            if (import.meta.env.VITE_ENV === "development") {
                console.error("Network Error:", error.message);
            }
        } else {
            if (import.meta.env.VITE_ENV === "development") {
                console.error("Error:", error.message);
            }
        }
        return Promise.reject(error);
    },
);

export default api;
