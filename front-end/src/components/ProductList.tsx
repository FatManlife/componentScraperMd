import ProductCard from "./ProductCard";
import type { Products } from "../constants/types";

type ProductListProps = {
    products: Products[];
};

function ProductList({ products }: ProductListProps) {
    if (!products || products.length === 0) {
        return (
            <div className="text-center py-8" style={{ color: '#8A8A8A' }}>
                No products found.
            </div>
        );
    }

    return (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {products.map((product) => (
                <ProductCard key={product.id} product={product} />
            ))}
        </div>
    );
}

export default ProductList;
