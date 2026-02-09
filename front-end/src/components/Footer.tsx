function Footer() {
    return (
        <footer className="mt-auto" style={{ backgroundColor: '#000000', color: '#FFFFFF', borderTop: '1px solid #8A8A8A' }}>
            <div className="container mx-auto px-4 py-6">
                <div className="flex flex-col md:flex-row justify-between items-center">
                    <div className="mb-4 md:mb-0">
                        <p className="text-sm">
                            © {new Date().getFullYear()} Component Finder. All rights reserved.
                        </p>
                    </div>
                    <div className="flex flex-col md:flex-row gap-3 md:gap-6 text-sm">
                        <span style={{ color: '#FFFFFF' }}>
                            info@componentfinder.com
                        </span>
                        <span style={{ color: '#FFFFFF' }}>
                            +1 (555) 123-4567
                        </span>
                    </div>
                </div>
            </div>
        </footer>
    );
}

export default Footer;
