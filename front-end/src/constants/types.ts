export type Product = {
    id: number;
    name: string;
    image_url: string;
    brand: string;
    price: number;
    url: string;
    category: string;
    website_id: number;
};

export type Products = {
    id: number;
    name: string;
    image_url: string;
    brand: string;
    price: number;
    category: string;
    website_id: number;
};

export type ProductResponse = Products[];

export type ProductOrder = "products.id ASC" | "price_asc" | "price_desc";

export type ProductParams = {
    name?: string;
    limit?: number;
    website?: string[];
    after?: number;
    min?: number;
    max?: number;
    order?: ProductOrder;
};
