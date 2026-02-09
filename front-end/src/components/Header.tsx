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
        navigate(`/?${newParams.toString()}`, { replace: true });
    };

    const categories = [
        { path: "/aio", label: "AIO" },
        { path: "/case", label: "Case" },
        { path: "/cooler", label: "Cooler" },
        { path: "/cpu", label: "CPU" },
        { path: "/fan", label: "Fan" },
        { path: "/gpu", label: "GPU" },
        { path: "/hdd", label: "HDD" },
        { path: "/laptop", label: "Laptop" },
        { path: "/motherboard", label: "Motherboard" },
        { path: "/pc", label: "PC" },
        { path: "/pc-mini", label: "PC Mini" },
        { path: "/psu", label: "PSU" },
        { path: "/ram", label: "RAM" },
        { path: "/ssd", label: "SSD" },
    ];

    return (
        <header className="sticky top-0 z-30 shadow-sm" style={{ backgroundColor: '#FFFFFF', borderBottom: '1px solid #D9D9D9' }}>
            <nav className="max-w-7xl mx-auto px-6">
                <div className="flex items-center gap-6 overflow-hidden" style={{ height: '72px' }}>
                    <h1 className="text-2xl sm:text-3xl font-bold tracking-tight shrink-0" style={{ color: '#000000' }}>
                        <Link 
                            to="/" 
                            className="transition-opacity duration-200 hover:opacity-70"
                            style={{ color: '#000000' }}
                        >
                            Component Finder
                        </Link>
                    </h1>
                    <div className="flex items-center gap-1 overflow-x-auto flex-1 [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
                        {categories.map((category) => (
                            <Link
                                key={category.path}
                                to={category.path}
                                className="px-3 py-1.5 transition-all duration-200 font-medium whitespace-nowrap text-sm"
                                style={{
                                    color: '#000000',
                                    backgroundColor: 'transparent',
                                    borderBottom: '2px solid transparent'
                                }}
                                onMouseEnter={(e) => {
                                    e.currentTarget.style.backgroundColor = '#F4F4F4';
                                    e.currentTarget.style.borderBottomColor = '#000000';
                                }}
                                onMouseLeave={(e) => {
                                    e.currentTarget.style.backgroundColor = 'transparent';
                                    e.currentTarget.style.borderBottomColor = 'transparent';
                                }}
                            >
                                {category.label}
                            </Link>
                        ))}
                    </div>
                </div>
                <div className="pb-5">
                    <SearchBar
                        onSearch={handleSearch}
                        placeholder="Search products by name..."
                    />
                </div>
            </nav>
        </header>
    );
}

export default Header;
