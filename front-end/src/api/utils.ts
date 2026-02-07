import type { ProductParams } from "../constants/types";

export const appendDefaultParams = (
    queryParams: URLSearchParams,
    params: ProductParams,
): void => {
    if (params.name) queryParams.append("name", params.name);
    if (params.website && params.website.length > 0) {
        params.website.forEach((site) => queryParams.append("website", site));
    }
    if (params.page !== undefined)
        queryParams.append("offset", params.page.toString());
    if (params.min !== undefined)
        queryParams.append("min", params.min.toString());
    if (params.max !== undefined)
        queryParams.append("max", params.max.toString());
    if (params.order) queryParams.append("order", params.order);
};
