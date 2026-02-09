import { Link } from "react-router-dom";
import type { Products } from "../constants/types";

type ProductCardProps = {
    product: Products;
};

function ProductCard({ product }: ProductCardProps) {
    return (
        <Link
            to={`/${product.category}/${product.id}`}
            className="block overflow-hidden transition-all duration-200"
            style={{ 
                backgroundColor: '#FFFFFF',
                border: '1px solid #D9D9D9',
                borderRadius: '2px'
            }}
            onMouseEnter={(e) => e.currentTarget.style.borderColor = '#8A8A8A'}
            onMouseLeave={(e) => e.currentTarget.style.borderColor = '#D9D9D9'}
        >
            <div className="aspect-square flex items-center justify-center" style={{ backgroundColor: '#F4F4F4' }}>
                {product.image_url ? (
                    <img
                        src={product.image_url}
                        alt={product.name}
                        className="w-full h-full object-cover"
                    />
                ) : (
                    <span style={{ color: '#8A8A8A' }}>No Image</span>
                )}
            </div>
            <div className="p-4">
                <h3 className="font-semibold mb-2 line-clamp-2 min-h-12" style={{ color: '#000000' }}>
                    {product.name}
                </h3>
                <p className="text-sm mb-2" style={{ color: '#8A8A8A' }}>
                    {product.brand && product.brand !== ""
                        ? product.brand
                        : "Unk"}
                </p>
                <div className="flex items-center justify-between mb-3">
                    <span className="text-lg font-bold" style={{ color: '#000000' }}>
                        MDL {product.price}
                    </span>
                    <span className="text-xs px-2 py-1" style={{ color: '#8A8A8A', backgroundColor: '#F4F4F4' }}>
                        {product.category}
                    </span>
                </div>
                <div className="flex items-center gap-2 pt-2" style={{ borderTop: '1px solid #D9D9D9' }}>
                    <img
                        src={product.website_image}
                        alt={product.website}
                        className="h-5 w-auto object-contain"
                    />
                    <span className="text-xs ml-auto" style={{ color: '#8A8A8A' }}>
                        {product.website}
                    </span>
                </div>
            </div>
        </Link>
    );
}

export default ProductCard;
