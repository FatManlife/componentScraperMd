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
        navigate(`?${newParams.toString()}`, { replace: true });
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
                    className="w-full px-3 py-2 text-sm focus:outline-none transition-all"
                    style={{
                        backgroundColor: '#F4F4F4',
                        border: '1px solid #D9D9D9',
                        borderRadius: '2px',
                        color: '#000000'
                    }}
                    onFocus={(e) => {
                        e.target.style.backgroundColor = '#FFFFFF';
                        e.target.style.borderColor = '#8A8A8A';
                    }}
                    onBlur={(e) => {
                        e.target.style.backgroundColor = '#F4F4F4';
                        e.target.style.borderColor = '#D9D9D9';
                    }}
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
                                className="flex items-center gap-2 cursor-pointer p-2 transition-colors"
                                style={{ backgroundColor: 'transparent' }}
                                onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                                onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
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
                                    className="w-4 h-4"
                                />
                                <span className="text-sm" style={{ color: '#000000' }}>
                                    {displayLabel}
                                </span>
                            </label>
                        );
                    })
                ) : (
                    <div className="text-sm p-2" style={{ color: '#8A8A8A' }}>
                        No results found
                    </div>
                )}
            </div>
        </div>
    );
}

export default ScrollableCheckboxList;
