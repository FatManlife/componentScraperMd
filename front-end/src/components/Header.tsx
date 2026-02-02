import { Link } from "react-router-dom";

function Header() {
    return (
        <header className="bg-white shadow-md">
            <nav className="container mx-auto px-4 py-4">
                <div className="flex items-center justify-between">
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
                    </ul>
                </div>
            </nav>
        </header>
    );
}

export default Header;
