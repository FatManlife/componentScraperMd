import type { AioParams, ProductResponse } from "../constants/types";
import api from "./axios";
import { appendDefaultParams } from "./utils";

export const FetchAio = async (params: AioParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        console.log("FetchAio params:", params);

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle AIO-specific params
        if (params.diagonal && params.diagonal.length > 0) {
            params.diagonal.forEach((d) =>
                queryParams.append("diagonal", d.toString()),
            );
        }
        if (params.cpu && params.cpu.length > 0) {
            params.cpu.forEach((c) => queryParams.append("cpu", c));
        }
        if (params.ram && params.ram.length > 0) {
            params.ram.forEach((r) => queryParams.append("ram", r.toString()));
        }
        if (params.storage && params.storage.length > 0) {
            params.storage.forEach((s) =>
                queryParams.append("storage", s.toString()),
            );
        }
        if (params.gpu && params.gpu.length > 0) {
            params.gpu.forEach((g) => queryParams.append("gpu", g));
        }

        console.log("Query string:", queryParams.toString());

        const response = await api.get<ProductResponse>(
            `/aio?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
