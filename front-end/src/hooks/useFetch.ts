import { useState } from "react";

export function useFetch<T>(fetchFn: () => Promise<T>) {
    const [data, setData] = useState<T | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const execute = async () => {
        setLoading(true);
        setError(null);
        try {
            const response = await fetchFn();
            setData(response);
        } catch (err) {
            setError("Failed to fetch data from server");
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    return { data, loading, error, execute };
}
