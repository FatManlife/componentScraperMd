import type { AioParams, ProductResponse } from "../constants/types";
import api from "./axios";

export const FetchAio = async (params: AioParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        const defaultParams = params.defaultParams;

        if (defaultParams.name) queryParams.append("name", defaultParams.name);
        if (defaultParams.website && defaultParams.website.length > 0) {
            defaultParams.website.forEach((site) =>
                queryParams.append("website", site),
            );
        }
        if (defaultParams.page !== undefined)
            queryParams.append("offset", defaultParams.page.toString());
        if (defaultParams.min !== undefined)
            queryParams.append("min", defaultParams.min.toString());
        if (defaultParams.max !== undefined)
            queryParams.append("max", defaultParams.max.toString());
        if (defaultParams.order)
            queryParams.append("order", defaultParams.order);

        // Handle AIO-specific params
        if (params.diagonal && params.diagonal.length > 0) {
            params.diagonal.forEach((d) => queryParams.append("diagonal", d));
        }
        if (params.cpu && params.cpu.length > 0) {
            params.cpu.forEach((c) => queryParams.append("cpu", c));
        }
        if (params.ram && params.ram.length > 0) {
            params.ram.forEach((r) => queryParams.append("ram", r));
        }
        if (params.storage && params.storage.length > 0) {
            params.storage.forEach((s) => queryParams.append("storage", s));
        }
        if (params.gpu && params.gpu.length > 0) {
            params.gpu.forEach((g) => queryParams.append("gpu", g));
        }

        const response = await api.get<ProductResponse>(
            `/aio?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
