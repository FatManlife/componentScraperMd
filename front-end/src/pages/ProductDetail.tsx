import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { FetchProduct } from "../api/products";
import type {
    AioResponse,
    CaseResponse,
    CoolerResponse,
    CpuResponse,
    FanResponse,
    GpuResponse,
    HddResponse,
    LaptopResponse,
    MotherboardResponse,
    PcResponse,
    PcMiniResponse,
    PsuResponse,
    RamResponse,
    SsdResponse,
} from "../constants/types";

type ComponentResponse =
    | AioResponse
    | CaseResponse
    | CoolerResponse
    | CpuResponse
    | FanResponse
    | GpuResponse
    | HddResponse
    | LaptopResponse
    | MotherboardResponse
    | PcResponse
    | PcMiniResponse
    | PsuResponse
    | RamResponse
    | SsdResponse;

function ProductDetail() {
    const { category, id } = useParams<{ category: string; id: string }>();
    const navigate = useNavigate();
    const [product, setProduct] = useState<ComponentResponse | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (!category || !id) {
            setError("Invalid product URL");
            setLoading(false);
            return;
        }

        const fetchProductDetails = async () => {
            try {
                setLoading(true);
                setError(null);
                // Normalize category: convert underscores to dashes for API endpoint
                const normalizedCategory = category.replace(/_/g, "-");
                const data = await FetchProduct<ComponentResponse>(
                    normalizedCategory,
                    parseInt(id, 10),
                );
                setProduct(data);
            } catch (err) {
                console.error(err);
                setError("Failed to load product details");
            } finally {
                setLoading(false);
            }
        };

        fetchProductDetails();
    }, [category, id]);

    if (loading) {
        return (
            <div className="max-w-6xl mx-auto px-4 py-8">
                <div className="flex items-center justify-center min-h-100">
                    <div className="text-xl" style={{ color: '#8A8A8A' }}>Loading...</div>
                </div>
            </div>
        );
    }

    if (error || !product) {
        return (
            <div className="max-w-6xl mx-auto px-4 py-8">
                <div className="flex flex-col items-center justify-center min-h-100">
                    <div className="text-xl mb-4" style={{ color: '#000000' }}>
                        {error || "Product not found"}
                    </div>
                    <button
                        onClick={() => navigate(-1)}
                        className="px-4 py-2 transition-colors"
                        style={{ backgroundColor: '#000000', color: '#FFFFFF', borderRadius: '2px' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#8A8A8A'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = '#000000'}
                    >
                        Go Back
                    </button>
                </div>
            </div>
        );
    }

    // Handle both Product (capital P) and product (lowercase p) from backend
    return (
        <div className="max-w-6xl mx-auto px-4 py-8">
            <button
                onClick={() => navigate(-1)}
                className="mb-6 flex items-center gap-2 transition-opacity"
                style={{ color: '#000000' }}
                onMouseEnter={(e) => e.currentTarget.style.opacity = '0.7'}
                onMouseLeave={(e) => e.currentTarget.style.opacity = '1'}
            >
                <span>←</span>
                <span>Back to results</span>
            </button>

            <div className="overflow-hidden" style={{ backgroundColor: '#FFFFFF', border: '1px solid #D9D9D9', borderRadius: '2px' }}>
                <div className="grid md:grid-cols-2 gap-8 p-8">
                    {/* Product Image */}
                    <div className="aspect-square flex items-center justify-center overflow-hidden" style={{ backgroundColor: '#F4F4F4', borderRadius: '2px' }}>
                        {product.product.image_url ? (
                            <img
                                src={product.product.image_url}
                                alt={product.product.name}
                                className="w-full h-full object-cover"
                            />
                        ) : (
                            <span style={{ color: '#8A8A8A' }}>No Image</span>
                        )}
                    </div>

                    {/* Product Details */}
                    <div className="flex flex-col">
                        <h1 className="text-3xl font-bold mb-4" style={{ color: '#000000' }}>
                            {product.product.name}
                        </h1>

                        <div className="mb-6">
                            <p className="mb-2" style={{ color: '#8A8A8A' }}>
                                <span className="font-semibold" style={{ color: '#000000' }}>Brand:</span>{" "}
                                {product.product.brand || "Unknown"}
                            </p>
                            <p className="mb-2" style={{ color: '#8A8A8A' }}>
                                <span className="font-semibold" style={{ color: '#000000' }}>Category:</span>{" "}
                                {product.product.category}
                            </p>
                        </div>

                        <div className="mb-6">
                            <div className="text-4xl font-bold mb-2" style={{ color: '#000000' }}>
                                MDL {product.product.price}
                            </div>
                        </div>

                        {/* Component-specific specs */}
                        <div className="mb-6 p-4" style={{ backgroundColor: '#F4F4F4', borderRadius: '2px' }}>
                            <h2 className="text-xl font-semibold mb-3" style={{ color: '#000000' }}>
                                Specifications
                            </h2>
                            <div className="space-y-2">
                                {Object.entries(product).map(([key, value]) => {
                                    if (key === "product") return null;
                                    return (
                                        <div
                                            key={key}
                                            className="flex justify-between py-2 last:border-0"
                                            style={{ borderBottom: '1px solid #D9D9D9' }}
                                        >
                                            <span className="capitalize" style={{ color: '#8A8A8A' }}>
                                                {key.replace(/_/g, " ")}:
                                            </span>
                                            <span className="font-semibold" style={{ color: '#000000' }}>
                                                {value?.toString() || "N/A"}
                                            </span>
                                        </div>
                                    );
                                })}
                            </div>
                        </div>

                        {/* Store Info */}
                        <div className="flex items-center gap-3 p-4 mb-6" style={{ backgroundColor: '#F4F4F4', borderRadius: '2px' }}>
                            <img
                                src={product.product.website_image}
                                alt={product.product.website}
                                className="h-8 w-auto object-contain"
                            />
                            <div>
                                <p className="text-sm" style={{ color: '#8A8A8A' }}>
                                    Available at
                                </p>
                                <p className="font-semibold" style={{ color: '#000000' }}>
                                    {product.product.website}
                                </p>
                            </div>
                        </div>

                        {/* View on Store Button */}
                        <a
                            href={product.product.url}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="w-full py-3 px-6 text-center font-semibold transition-colors"
                            style={{ backgroundColor: '#000000', color: '#FFFFFF', borderRadius: '2px', display: 'block' }}
                            onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#8A8A8A'}
                            onMouseLeave={(e) => e.currentTarget.style.backgroundColor = '#000000'}
                        >
                            View on {product.product.website}
                        </a>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default ProductDetail;
