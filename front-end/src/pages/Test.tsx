import { useEffect, useState } from "react";
import { FetchProductFilters } from "../api/filters";
import Sidebar from "../components/Sidebar";
import type { DefaultSpecs } from "../constants/types";

function Test() {
    const [filters, setFilters] = useState<DefaultSpecs | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            setError(null);
            try {
                const data = await FetchProductFilters();
                setFilters(data);
            } catch (err) {
                setError("Failed to fetch filters");
                console.error(err);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    return (
        <div className="bg-gray-100 min-h-screen py-8 px-4">
            <div className="max-w-6xl mx-auto">
                <h1 className="text-3xl font-bold text-gray-800 mb-6">
                    Test Default Filters
                </h1>

                {loading && (
                    <div className="text-center text-gray-600">
                        Loading filters...
                    </div>
                )}

                {error && (
                    <div className="p-3 bg-red-100 border border-red-400 text-red-700 rounded">
                        {error}
                    </div>
                )}

                <div className="flex gap-6">
                    <Sidebar filters={filters} />

                    {filters && (
                        <div className="flex-1 bg-white rounded-lg shadow-md p-6">
                            <h2 className="text-xl font-bold text-gray-800 mb-4">
                                Raw Filter Data
                            </h2>
                            <pre className="text-sm overflow-auto bg-gray-50 p-4 rounded">
                                {JSON.stringify(filters, null, 2)}
                            </pre>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}

export default Test;
