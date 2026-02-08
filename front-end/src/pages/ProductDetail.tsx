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
                    <div className="text-xl text-gray-600">Loading...</div>
                </div>
            </div>
        );
    }

    if (error || !product) {
        return (
            <div className="max-w-6xl mx-auto px-4 py-8">
                <div className="flex flex-col items-center justify-center min-h-100">
                    <div className="text-xl text-red-600 mb-4">
                        {error || "Product not found"}
                    </div>
                    <button
                        onClick={() => navigate(-1)}
                        className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
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
                className="mb-6 flex items-center gap-2 text-blue-600 hover:text-blue-700"
            >
                <span>←</span>
                <span>Back to results</span>
            </button>

            <div className="bg-white rounded-lg shadow-lg overflow-hidden">
                <div className="grid md:grid-cols-2 gap-8 p-8">
                    {/* Product Image */}
                    <div className="aspect-square bg-gray-200 flex items-center justify-center rounded-lg overflow-hidden">
                        {product.product.image_url ? (
                            <img
                                src={product.product.image_url}
                                alt={product.product.name}
                                className="w-full h-full object-cover"
                            />
                        ) : (
                            <span className="text-gray-400">No Image</span>
                        )}
                    </div>

                    {/* Product Details */}
                    <div className="flex flex-col">
                        <h1 className="text-3xl font-bold text-gray-800 mb-4">
                            {product.product.name}
                        </h1>

                        <div className="mb-6">
                            <p className="text-gray-600 mb-2">
                                <span className="font-semibold">Brand:</span>{" "}
                                {product.product.brand || "Unknown"}
                            </p>
                            <p className="text-gray-600 mb-2">
                                <span className="font-semibold">Category:</span>{" "}
                                {product.product.category}
                            </p>
                        </div>

                        <div className="mb-6">
                            <div className="text-4xl font-bold text-blue-600 mb-2">
                                MDL {product.product.price}
                            </div>
                        </div>

                        {/* Component-specific specs */}
                        <div className="mb-6 p-4 bg-gray-50 rounded-lg">
                            <h2 className="text-xl font-semibold text-gray-800 mb-3">
                                Specifications
                            </h2>
                            <div className="space-y-2">
                                {Object.entries(product).map(([key, value]) => {
                                    if (key === "product") return null;
                                    return (
                                        <div
                                            key={key}
                                            className="flex justify-between py-2 border-b border-gray-200 last:border-0"
                                        >
                                            <span className="text-gray-600 capitalize">
                                                {key.replace(/_/g, " ")}:
                                            </span>
                                            <span className="font-semibold text-gray-800">
                                                {value?.toString() || "N/A"}
                                            </span>
                                        </div>
                                    );
                                })}
                            </div>
                        </div>

                        {/* Store Info */}
                        <div className="flex items-center gap-3 p-4 bg-gray-50 rounded-lg mb-6">
                            <img
                                src={product.product.website_image}
                                alt={product.product.website}
                                className="h-8 w-auto object-contain"
                            />
                            <div>
                                <p className="text-sm text-gray-600">
                                    Available at
                                </p>
                                <p className="font-semibold text-gray-800">
                                    {product.product.website}
                                </p>
                            </div>
                        </div>

                        {/* View on Store Button */}
                        <a
                            href={product.product.url}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="w-full bg-blue-600 text-white py-3 px-6 rounded-lg hover:bg-blue-700 transition-colors text-center font-semibold"
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
