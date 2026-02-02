function Footer() {
    return (
        <footer className="bg-gray-800 text-white mt-auto">
            <div className="container mx-auto px-4 py-6">
                <div className="flex flex-col md:flex-row justify-between items-center">
                    <div className="mb-4 md:mb-0">
                        <p className="text-sm">
                            © {new Date().getFullYear()} Component Finder. All rights reserved.
                        </p>
                    </div>
                    <div className="flex space-x-6">
                        <a
                            href="/about"
                            className="text-sm hover:text-blue-400 transition"
                        >
                            About
                        </a>
                        <a
                            href="/contact"
                            className="text-sm hover:text-blue-400 transition"
                        >
                            Contact
                        </a>
                        <a
                            href="/privacy"
                            className="text-sm hover:text-blue-400 transition"
                        >
                            Privacy
                        </a>
                    </div>
                </div>
            </div>
        </footer>
    );
}

export default Footer;
