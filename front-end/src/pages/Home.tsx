import { useEffect, useMemo, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { FetchProducts } from "../api/products";
import { useFetch } from "../hooks/useFetch";
import SearchBar from "../components/SearchBar";
import type { Products, ProductParams, ProductOrder } from "../constants/types";

function Home() {
    const [searchParams] = useSearchParams();
    const navigate = useNavigate();

    const params = useMemo(() => {
        const productParams: ProductParams = {};

        const limit = searchParams.get("limit");
        if (limit) productParams.limit = parseInt(limit);

        const name = searchParams.get("name");
        if (name) productParams.name = name;

        const website = searchParams.getAll("website");
        if (website.length > 0) productParams.website = website;

        const after = searchParams.get("after");
        if (after) productParams.after = parseInt(after);

        const min = searchParams.get("min");
        if (min) productParams.min = parseFloat(min);

        const max = searchParams.get("max");
        if (max) productParams.max = parseFloat(max);

        const order = searchParams.get("order");
        if (order && (order === "products.id ASC" || order === "price_asc" || order === "price_desc")) {
            productParams.order = order as ProductOrder;
        }

        return productParams;
    }, [searchParams]);

    const { data, loading, error, execute } = useFetch<Products>(() =>
        FetchProducts(params)
    );

    useEffect(() => {
        execute();
    }, [params]);

    const handleSearch = (searchTerm: string) => {
        const newParams = new URLSearchParams(searchParams);
        if (searchTerm) {
            newParams.set("name", searchTerm);
        } else {
            newParams.delete("name");
        }
        navigate(`/?${newParams.toString()}`);
    };

    return (
        <div className="bg-gray-100 py-8 px-4">
            <div className="max-w-6xl mx-auto">
                <h2 className="text-3xl font-bold text-gray-800 mb-6">
                    Products
                </h2>

                <SearchBar 
                    onSearch={handleSearch} 
                    placeholder="Search products by name..."
                />

                {loading && (
                    <div className="text-center text-gray-600">Loading...</div>
                )}

                {error && (
                    <div className="mt-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
                        {error}
                    </div>
                )}

                {data && (
                    <div className="mt-4 p-3 bg-green-100 border border-green-400 text-green-700 rounded">
                        <pre className="text-sm overflow-auto">
                            {JSON.stringify(data, null, 2)}
                        </pre>
                    </div>
                )}
            </div>
        </div>
    );
}

export default Home;
