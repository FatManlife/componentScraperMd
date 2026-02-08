import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchCase } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    CaseParams,
    ProductOrder,
    ComponentFiltersResponse,
    CaseSpecs,
} from "../constants/types";

function Case() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<CaseSpecs> | null>(null);

    const params = useMemo<CaseParams>(() => {
        return {
            defaultParams: {
                name: searchParams.get("name") ?? undefined,
                website: searchParams.getAll("website"),
                page: searchParams.get("page")
                    ? parseInt(searchParams.get("page")!)
                    : 1,
                min: searchParams.get("min")
                    ? parseFloat(searchParams.get("min")!)
                    : undefined,
                max: searchParams.get("max")
                    ? parseFloat(searchParams.get("max")!)
                    : undefined,
                order: ["products.id ASC", "price_asc", "price_desc"].includes(
                    searchParams.get("order") ?? "",
                )
                    ? (searchParams.get("order") as ProductOrder)
                    : undefined,
            },
            format: searchParams.getAll("format"),
            motherboard_form_factor: searchParams.getAll(
                "motherboard_form_factor",
            ),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchCase(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        // Fetch filters for Case category (only once on mount)
        FetchComponentFilters<CaseSpecs>("case")
            .then(setFilters)
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="Case Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="case"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Case;
