import api from "./axios";
import type { DefaultFilters } from "../constants/types";

export const FetchFilters = async (
    category?: string,
): Promise<DefaultFilters> => {
    try {
        const queryParams = new URLSearchParams();
        if (category) queryParams.append("category", category);

        const response = await api.get<DefaultFilters>(
            `/filter?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
