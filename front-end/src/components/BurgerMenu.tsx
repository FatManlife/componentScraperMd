import { useState } from "react";
import { Link } from "react-router-dom";

interface MenuItem {
    path: string;
    label: string;
}

interface BurgerMenuProps {
    items: MenuItem[];
}

function BurgerMenu({ items }: BurgerMenuProps) {
    const [isOpen, setIsOpen] = useState(false);

    const toggleMenu = () => {
        setIsOpen(!isOpen);
    };

    const closeMenu = () => {
        setIsOpen(false);
    };

    return (
        <div className="flex items-center gap-4">
            {/* Burger Icon */}
            <button
                onClick={toggleMenu}
                className="flex flex-col justify-center items-center w-11 h-11 space-y-1.5 focus:outline-none transition-all duration-200 shrink-0"
                style={{ backgroundColor: isOpen ? "#F4F4F4" : "transparent" }}
                aria-label="Toggle menu"
            >
                <span
                    className={`block w-6 h-0.5 transition-all duration-300 ease-in-out ${
                        isOpen ? "rotate-45 translate-y-2" : ""
                    }`}
                    style={{ backgroundColor: "#000000" }}
                />
                <span
                    className={`block w-6 h-0.5 transition-all duration-300 ease-in-out ${
                        isOpen ? "opacity-0" : ""
                    }`}
                    style={{ backgroundColor: "#000000" }}
                />
                <span
                    className={`block w-6 h-0.5 transition-all duration-300 ease-in-out ${
                        isOpen ? "-rotate-45 -translate-y-2" : ""
                    }`}
                    style={{ backgroundColor: "#000000" }}
                />
            </button>

            {/* Horizontal Menu */}
            <div
                className="overflow-hidden transition-all duration-300 ease-in-out"
                style={{
                    maxWidth: isOpen ? "1000px" : "0",
                    opacity: isOpen ? 1 : 0,
                }}
            >
                <ul className="flex items-center gap-1 whitespace-nowrap">
                    {items.map((item) => (
                        <li key={item.path}>
                            <Link
                                to={item.path}
                                onClick={closeMenu}
                                className="block px-4 py-2 transition-all duration-200 font-medium"
                                style={{
                                    color: "#000000",
                                    backgroundColor: "transparent",
                                    borderBottom: "2px solid transparent",
                                }}
                                onMouseEnter={(e) => {
                                    e.currentTarget.style.backgroundColor =
                                        "#F4F4F4";
                                    e.currentTarget.style.borderBottomColor =
                                        "#000000";
                                }}
                                onMouseLeave={(e) => {
                                    e.currentTarget.style.backgroundColor =
                                        "transparent";
                                    e.currentTarget.style.borderBottomColor =
                                        "transparent";
                                }}
                            >
                                {item.label}
                            </Link>
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
}

export default BurgerMenu;
