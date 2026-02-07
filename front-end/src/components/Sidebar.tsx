import { useNavigate, useSearchParams } from "react-router-dom";
import { useState } from "react";
import RangeSlider from "./RangeSlider";
import type { DefaultSpecs, ProductOrder } from "../constants/types";
import type { ReactNode } from "react";

type SidebarProps = {
    filters: DefaultSpecs | null;
    specificFilters?: ReactNode;
};

function Sidebar({ filters, specificFilters }: SidebarProps) {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const [openSections, setOpenSections] = useState<{
        websites: boolean;
        price: boolean;
        sort: boolean;
    }>({
        websites: false,
        price: false,
        sort: false,
    });

    if (!filters) {
        return null;
    }

    const selectedWebsites = searchParams.getAll("website");
    const selectedOrder = (searchParams.get("order") as ProductOrder) || "";

    const toggleSection = (section: "websites" | "price" | "sort") => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    const handleWebsiteChange = (website: string, checked: boolean) => {
        const newWebsites = checked
            ? [...selectedWebsites, website]
            : selectedWebsites.filter((w) => w !== website);

        const newParams = new URLSearchParams(searchParams);
        newParams.delete("website");
        newParams.delete("page");
        newWebsites.forEach((w) => newParams.append("website", w));
        navigate(`?${newParams.toString()}`);
    };

    const handleOrderChange = (order: ProductOrder) => {
        const newParams = new URLSearchParams(searchParams);
        if (order) {
            newParams.set("order", order);
        } else {
            newParams.delete("order");
        }
        newParams.delete("page");
        navigate(`?${newParams.toString()}`);
    };

    const handleResetFilters = () => {
        // Keep only the search name if it exists
        const newParams = new URLSearchParams();
        const name = searchParams.get("name");
        if (name) {
            newParams.set("name", name);
        }
        navigate(`?${newParams.toString()}`);
    };

    return (
        <aside className="w-64 bg-white rounded-lg shadow-md p-4 sticky top-4 self-start max-h-[calc(100vh-2rem)] overflow-y-auto">
            <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-bold text-gray-800">Filters</h3>
                <button
                    onClick={handleResetFilters}
                    className="text-sm text-blue-600 hover:text-blue-800 font-medium px-2 py-1 rounded hover:bg-blue-50 transition-colors"
                >
                    Reset
                </button>
            </div>

            {/* Websites Section */}
            <div className="mb-6">
                <button
                    onClick={() => toggleSection("websites")}
                    className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                >
                    <h4 className="text-sm font-semibold text-gray-700">
                        Websites
                    </h4>
                    <span className="text-gray-500">
                        {openSections.websites ? "−" : "+"}
                    </span>
                </button>
                {openSections.websites && (
                    <div className="space-y-2">
                        {filters.websites.map((website, index) => (
                            <label
                                key={index}
                                className="flex items-center gap-2 cursor-pointer hover:bg-gray-50 p-2 rounded"
                            >
                                <input
                                    type="checkbox"
                                    value={website}
                                    checked={selectedWebsites.includes(website)}
                                    onChange={(e) =>
                                        handleWebsiteChange(
                                            website,
                                            e.target.checked,
                                        )
                                    }
                                    className="w-4 h-4 text-blue-600 rounded focus:ring-2 focus:ring-blue-500"
                                />
                                <span className="text-sm text-gray-700">
                                    {website}
                                </span>
                            </label>
                        ))}
                    </div>
                )}
            </div>

            {/* Price Range Slider */}
            <div className="mb-6">
                <button
                    onClick={() => toggleSection("price")}
                    className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                >
                    <h4 className="text-sm font-semibold text-gray-700">
                        Price Range
                    </h4>
                    <span className="text-gray-500">
                        {openSections.price ? "−" : "+"}
                    </span>
                </button>
                {openSections.price && (
                    <RangeSlider
                        values={filters.prices}
                        label=""
                        text="Price"
                    />
                )}
            </div>

            {/* Order Section */}
            <div className="mb-6">
                <button
                    onClick={() => toggleSection("sort")}
                    className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                >
                    <h4 className="text-sm font-semibold text-gray-700">
                        Sort By
                    </h4>
                    <span className="text-gray-500">
                        {openSections.sort ? "−" : "+"}
                    </span>
                </button>
                {openSections.sort && (
                    <div className="space-y-2">
                        <label className="flex items-center gap-2 cursor-pointer hover:bg-gray-50 p-2 rounded">
                            <input
                                type="radio"
                                name="order"
                                value="products.id ASC"
                                checked={selectedOrder === "products.id ASC"}
                                onChange={(e) =>
                                    handleOrderChange(
                                        e.target.value as ProductOrder,
                                    )
                                }
                                className="w-4 h-4 text-blue-600 focus:ring-2 focus:ring-blue-500"
                            />
                            <span className="text-sm text-gray-700">
                                Default
                            </span>
                        </label>
                        <label className="flex items-center gap-2 cursor-pointer hover:bg-gray-50 p-2 rounded">
                            <input
                                type="radio"
                                name="order"
                                value="price_asc"
                                checked={selectedOrder === "price_asc"}
                                onChange={(e) =>
                                    handleOrderChange(
                                        e.target.value as ProductOrder,
                                    )
                                }
                                className="w-4 h-4 text-blue-600 focus:ring-2 focus:ring-blue-500"
                            />
                            <span className="text-sm text-gray-700">
                                Price: Low to High
                            </span>
                        </label>
                        <label className="flex items-center gap-2 cursor-pointer hover:bg-gray-50 p-2 rounded">
                            <input
                                type="radio"
                                name="order"
                                value="price_desc"
                                checked={selectedOrder === "price_desc"}
                                onChange={(e) =>
                                    handleOrderChange(
                                        e.target.value as ProductOrder,
                                    )
                                }
                                className="w-4 h-4 text-blue-600 focus:ring-2 focus:ring-blue-500"
                            />
                            <span className="text-sm text-gray-700">
                                Price: High to Low
                            </span>
                        </label>
                    </div>
                )}
            </div>

            {/* Specific Filters */}
            {specificFilters}
        </aside>
    );
}

export default Sidebar;
