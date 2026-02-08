import api from "./axios";
import type { ProductParams, ProductResponse } from "../constants/types";
import { appendDefaultParams } from "./utils";

export const FetchProducts = async (
    params: ProductParams,
): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        appendDefaultParams(queryParams, params);

        const response = await api.get<ProductResponse>(
            `product?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchProduct = async <T>(
    category: string,
    id: number,
): Promise<T> => {
    try {
        const response = await api.get<T>(`/${category}/${id}`);
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
