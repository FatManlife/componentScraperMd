import api from "./axios";
import type { ProductParams, ProductResponse } from "../constants/types";

export const FetchProducts = async (
    params: ProductParams,
): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        if (params.name) queryParams.append("name", params.name);
        if (params.website && params.website.length > 0) {
            params.website.forEach((site) =>
                queryParams.append("website", site),
            );
        }
        if (params.page !== undefined)
            queryParams.append("offset", params.page.toString());
        if (params.min !== undefined)
            queryParams.append("min", params.min.toString());
        if (params.max !== undefined)
            queryParams.append("max", params.max.toString());
        if (params.order) queryParams.append("order", params.order);

        const response = await api.get<ProductResponse>(
            `/?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
