import type { ReactNode } from "react";
import Pagination from "./Pagination";
import ProductList from "./ProductList";
import Sidebar from "./Sidebar";
import type { Products, DefaultFilters } from "../constants/types";

type ProductListLayoutProps = {
    title: string;
    loading: boolean;
    error: string | null;
    data: Products[] | null;
    currentPage: number;
    totalCount: number | null;
    filters: DefaultFilters | null;
    children?: ReactNode;
};

function ProductListLayout({
    title,
    loading,
    error,
    data,
    currentPage,
    totalCount,
    filters,
    children,
}: ProductListLayoutProps) {
    return (
        <div className="bg-gray-100 py-8 px-4">
            <div className="max-w-6xl mx-auto">
                <h2 className="text-3xl font-bold text-gray-800 mb-6">
                    {title}
                </h2>

                {children}

                <div className="flex gap-6">
                    <Sidebar filters={filters} />
                    
                    <div className="flex-1 flex flex-col">
                        <div className="min-h-[800px]">
                            {loading && (
                                <div className="text-center text-gray-600 py-20">Loading...</div>
                            )}

                            {error && (
                                <div className="mt-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
                                    {error}
                                </div>
                            )}

                            {data && !loading && <ProductList products={data} />}
                        </div>

                        <div className="h-16 flex items-center justify-center mt-6">
                            {totalCount !== null && (
                                <Pagination
                                    currentPage={currentPage}
                                    totalCount={totalCount}
                                    itemsPerPage={24}
                                />
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default ProductListLayout;
