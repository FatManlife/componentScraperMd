import type { Products } from "../constants/types";

type ProductCardProps = {
    product: Products;
};

function ProductCard({ product }: ProductCardProps) {
    return (
        <div className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
            <div className="aspect-square bg-gray-200 flex items-center justify-center">
                {product.image_url ? (
                    <img
                        src={product.image_url}
                        alt={product.name}
                        className="w-full h-full object-cover"
                    />
                ) : (
                    <span className="text-gray-400">No Image</span>
                )}
            </div>
            <div className="p-4">
                <h3 className="font-semibold text-gray-800 mb-2 line-clamp-2 min-h-12">
                    {product.name}
                </h3>
                <p className="text-sm text-gray-600 mb-2">{product.brand}</p>
                <div className="flex items-center justify-between mb-3">
                    <span className="text-lg font-bold text-blue-600">
                        MDL{product.price}
                    </span>
                    <span className="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded">
                        {product.category}
                    </span>
                </div>
                <div className="flex items-center gap-2 pt-2 border-t border-gray-200">
                    <img
                        src={product.website_image}
                        alt={product.website}
                        className="h-5 w-auto object-contain"
                    />
                    <span className="text-xs text-gray-500 ml-auto">
                        {product.website}
                    </span>
                </div>
            </div>
        </div>
    );
}

export default ProductCard;
