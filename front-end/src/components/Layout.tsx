import Header from "./Header";
import Footer from "./Footer";

type LayoutProps = {
    children: React.ReactNode;
};

function Layout({ children }: LayoutProps) {
    return (
        <div className="flex flex-col min-h-screen" style={{ backgroundColor: '#F4F4F4' }}>
            <Header />
            <main className="grow">{children}</main>
            <Footer />
        </div>
    );
}

export default Layout;
