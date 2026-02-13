import api from "./axios";
import type {
    DefaultSpecs,
    ComponentFiltersResponse,
} from "../constants/types";

export const FetchProductFilters = async (): Promise<DefaultSpecs> => {
    try {
        const response = await api.get<DefaultSpecs>(`product/spec`);
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchComponentFilters = async <T>(
    category: string,
): Promise<ComponentFiltersResponse<T>> => {
    try {
        const response = await api.get<ComponentFiltersResponse<T>>(
            `${category}/spec`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
