import type { ProductParams } from "../constants/types";
import { sanitizeString, sanitizeStringArray, sanitizeNumber } from "./validation";

export const appendDefaultParams = (
    queryParams: URLSearchParams,
    params: ProductParams,
): void => {
    if (params.name) {
        queryParams.append("name", sanitizeString(params.name));
    }
    if (params.website && params.website.length > 0) {
        const sanitizedWebsites = sanitizeStringArray(params.website);
        sanitizedWebsites.forEach((site) => queryParams.append("website", site));
    }
    if (params.page !== undefined) {
        queryParams.append("offset", sanitizeNumber(params.page).toString());
    }
    if (params.min !== undefined) {
        queryParams.append("min", sanitizeNumber(params.min).toString());
    }
    if (params.max !== undefined) {
        queryParams.append("max", sanitizeNumber(params.max).toString());
    }
    if (params.order) {
        queryParams.append("order", sanitizeString(params.order));
    }
};
