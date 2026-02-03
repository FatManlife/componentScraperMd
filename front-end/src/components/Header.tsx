import { Link, useNavigate, useSearchParams } from "react-router-dom";
import SearchBar from "./SearchBar";

function Header() {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();

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
        <header className="bg-white shadow-md">
            <nav className="container mx-auto px-4 py-4">
                <div className="flex items-center justify-between mb-4">
                    <h1 className="text-2xl font-bold text-gray-800">
                        Component Finder
                    </h1>
                    <ul className="flex space-x-6">
                        <li>
                            <Link
                                to="/"
                                className="text-gray-600 hover:text-blue-500 transition"
                            >
                                Home
                            </Link>
                        </li>
                        <li>
                            <Link
                                to="/aio"
                                className="text-gray-600 hover:text-blue-500 transition"
                            >
                                AIO
                            </Link>
                        </li>
                        <li>
                            <Link
                                to="/test"
                                className="text-gray-600 hover:text-blue-500 transition"
                            >
                                Test
                            </Link>
                        </li>
                    </ul>
                </div>
                <SearchBar
                    onSearch={handleSearch}
                    placeholder="Search products by name..."
                />
            </nav>
        </header>
    );
}

export default Header;
