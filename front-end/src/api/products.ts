import api from "./axios";
import type { Products, ProductParams } from "../constants/types";

export const FetchProducts = async (
    params: ProductParams = { limit: 20 },
): Promise<Products> => {
    try {
        const queryParams = new URLSearchParams();

        // Set default limit if not provided
        const limit = params.limit ?? 20;
        queryParams.append("limit", limit.toString());

        if (params.name) queryParams.append("name", params.name);
        if (params.website && params.website.length > 0) {
            params.website.forEach((site) =>
                queryParams.append("website", site),
            );
        }
        if (params.after !== undefined)
            queryParams.append("after", params.after.toString());
        if (params.min !== undefined)
            queryParams.append("min", params.min.toString());
        if (params.max !== undefined)
            queryParams.append("max", params.max.toString());
        if (params.order) queryParams.append("order", params.order);

        const response = await api.get<Products>(`/?${queryParams.toString()}`);
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
