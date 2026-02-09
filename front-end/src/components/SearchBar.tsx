import { useState } from "react";

type SearchBarProps = {
    onSearch: (searchTerm: string) => void;
    placeholder?: string;
};

function SearchBar({
    onSearch,
    placeholder = "Search products...",
}: SearchBarProps) {
    const [searchTerm, setSearchTerm] = useState("");

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        onSearch(searchTerm);
    };

    const handleClear = () => {
        setSearchTerm("");
        onSearch("");
    };

    return (
        <form onSubmit={handleSubmit} className="w-full max-w-3xl">
            <div className="flex gap-3">
                <input
                    type="text"
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    placeholder={placeholder}
                    className="flex-1 px-4 py-3 focus:outline-none transition-all duration-200"
                    style={{
                        backgroundColor: "#F4F4F4",
                        border: "1px solid #D9D9D9",
                        borderRadius: "2px",
                        color: "#000000",
                    }}
                    onFocus={(e) => {
                        e.target.style.backgroundColor = "#FFFFFF";
                        e.target.style.borderColor = "#8A8A8A";
                    }}
                    onBlur={(e) => {
                        e.target.style.backgroundColor = "#F4F4F4";
                        e.target.style.borderColor = "#D9D9D9";
                    }}
                />
                <button
                    type="submit"
                    className="px-8 py-3 font-medium transition-all duration-200"
                    style={{
                        backgroundColor: "#000000",
                        color: "#FFFFFF",
                        borderRadius: "2px",
                    }}
                    onMouseEnter={(e) =>
                        (e.currentTarget.style.backgroundColor = "#8A8A8A")
                    }
                    onMouseLeave={(e) =>
                        (e.currentTarget.style.backgroundColor = "#000000")
                    }
                >
                    Search
                </button>
                {searchTerm && (
                    <button
                        type="button"
                        onClick={handleClear}
                        className="px-5 py-3 font-medium transition-all duration-200"
                        style={{
                            backgroundColor: "#F4F4F4",
                            color: "#000000",
                            border: "1px solid #D9D9D9",
                            borderRadius: "2px",
                        }}
                        onMouseEnter={(e) =>
                            (e.currentTarget.style.backgroundColor = "#D9D9D9")
                        }
                        onMouseLeave={(e) =>
                            (e.currentTarget.style.backgroundColor = "#F4F4F4")
                        }
                    >
                        Clear
                    </button>
                )}
            </div>
        </form>
    );
}

export default SearchBar;
