import { useNavigate, useSearchParams } from "react-router-dom";
import { useState } from "react";

type ScrollableCheckboxListProps = {
    items: string[];
    paramName: string;
    label: string;
    formatLabel?: (item: string) => string;
};

function ScrollableCheckboxList({
    items,
    paramName,
    label,
    formatLabel,
}: ScrollableCheckboxListProps) {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const [searchQuery, setSearchQuery] = useState("");

    const selectedItems = searchParams.getAll(paramName);

    const filteredItems = items.filter((item) =>
        item.toLowerCase().includes(searchQuery.toLowerCase()),
    );

    const handleCheckboxChange = (value: string, checked: boolean) => {
        const updated = checked
            ? [...selectedItems, value]
            : selectedItems.filter((v) => v !== value);

        const newParams = new URLSearchParams(searchParams);
        newParams.delete(paramName);
        newParams.delete("page");
        updated.forEach((v) => newParams.append(paramName, v));
        navigate(`?${newParams.toString()}`);
    };

    return (
        <div>
            {/* Search Input */}
            <div className="mb-2">
                <input
                    type="text"
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                    placeholder={`Search ${label}...`}
                    className="w-full px-3 py-2 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
            </div>

            {/* Scrollable List */}
            <div
                className="space-y-2 max-h-50 overflow-y-auto scrollbar-hide"
                style={{
                    scrollbarWidth: "none",
                    msOverflowStyle: "none",
                }}
            >
                <style>{`
                    .scrollbar-hide::-webkit-scrollbar {
                        display: none;
                    }
                `}</style>
                {filteredItems.length > 0 ? (
                    filteredItems.map((item, index) => {
                        const displayLabel = formatLabel
                            ? formatLabel(item)
                            : item === null ||
                                item === "" ||
                                item === "null" ||
                                item === "0"
                              ? "Unk"
                              : item;

                        return (
                            <label
                                key={index}
                                className="flex items-center gap-2 cursor-pointer hover:bg-gray-50 p-2 rounded"
                            >
                                <input
                                    type="checkbox"
                                    value={item}
                                    checked={selectedItems.includes(item)}
                                    onChange={(e) =>
                                        handleCheckboxChange(
                                            item,
                                            e.target.checked,
                                        )
                                    }
                                    className="w-4 h-4 text-blue-600 rounded focus:ring-2 focus:ring-blue-500"
                                />
                                <span className="text-sm text-gray-700">
                                    {displayLabel}
                                </span>
                            </label>
                        );
                    })
                ) : (
                    <div className="text-sm text-gray-500 p-2">
                        No results found
                    </div>
                )}
            </div>
        </div>
    );
}

export default ScrollableCheckboxList;
