export type Product = {
    id: number;
    name: string;
    image_url: string;
    brand: string;
    price: number;
    url: string;
    category: string;
    website: string;
    website_image: string;
};

export type Products = {
    id: number;
    name: string;
    image_url: string;
    brand: string;
    price: number;
    category: string;
    website: string;
    website_image: string;
};

export type ProductResponse = {
    count: number;
    products: Product[];
};

export type ProductOrder = "products.id ASC" | "price_asc" | "price_desc";

export type ProductParams = {
    name?: string;
    website?: string[];
    page?: number;
    min?: number;
    max?: number;
    order?: ProductOrder;
};

export type AioResponse = {
    product: Product;
    diagonal: string;
    cpu: string;
    ram: string;
    storage: string;
    gpu: string;
};

export type AioParams = {
    defaultParams: ProductParams;
    diagonal?: string[];
    cpu?: string[];
    ram?: string[];
    storage?: string[];
    gpu?: string[];
};

export type DefaultFilters = {
    websites: string[];
    prices: number[];
    order: string[];
};
