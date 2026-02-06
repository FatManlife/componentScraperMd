import { useNavigate, useSearchParams } from "react-router-dom";
import RangeSlider from "./RangeSlider";
import type { DefaultFilters, ProductOrder } from "../constants/types";

type SidebarProps = {
    filters: DefaultFilters | null;
};

function Sidebar({ filters }: SidebarProps) {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();

    if (!filters) {
        return null;
    }

    const selectedWebsites = searchParams.getAll("website");
    const selectedOrder = (searchParams.get("order") as ProductOrder) || "";

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

    return (
        <aside className="w-64 bg-white rounded-lg shadow-md p-4 sticky top-4 self-start max-h-[calc(100vh-2rem)] overflow-y-auto">
            <h3 className="text-lg font-bold text-gray-800 mb-4">Filters</h3>

            {/* Websites Section */}
            <div className="mb-6">
                <h4 className="text-sm font-semibold text-gray-700 mb-2">
                    Websites
                </h4>
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
            </div>

            {/* Price Range Slider */}
            <RangeSlider
                values={filters.prices}
                label="Price Range"
                text="Price"
            />

            {/* Order Section */}
            <div>
                <h4 className="text-sm font-semibold text-gray-700 mb-2">
                    Sort By
                </h4>
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
                        <span className="text-sm text-gray-700">Default</span>
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
            </div>
        </aside>
    );
}

export default Sidebar;
