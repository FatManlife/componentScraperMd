import { useState } from "react";

type SearchBarProps = {
    onSearch: (searchTerm: string) => void;
    placeholder?: string;
};

function SearchBar({ onSearch, placeholder = "Search products..." }: SearchBarProps) {
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
        <form onSubmit={handleSubmit} className="w-full max-w-2xl mx-auto mb-6">
            <div className="flex gap-2">
                <input
                    type="text"
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    placeholder={placeholder}
                    className="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
                <button
                    type="submit"
                    className="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
                >
                    Search
                </button>
                {searchTerm && (
                    <button
                        type="button"
                        onClick={handleClear}
                        className="px-4 py-2 bg-gray-300 text-gray-700 rounded-lg hover:bg-gray-400 transition"
                    >
                        Clear
                    </button>
                )}
            </div>
        </form>
    );
}

export default SearchBar;
