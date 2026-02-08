import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchCooler } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    CoolerParams,
    ProductOrder,
    ComponentFiltersResponse,
    CoolerSpecs,
} from "../constants/types";

function Cooler() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<CoolerSpecs> | null>(null);

    const params = useMemo<CoolerParams>(() => {
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
            type: searchParams.getAll("type"),
            fan_rpm: searchParams
                .getAll("fan_rpm")
                .map((r) => parseInt(r, 10))
                .filter((r) => !isNaN(r)),
            noise: searchParams
                .getAll("noise")
                .map((n) => parseFloat(n))
                .filter((n) => !isNaN(n)),
            compatibility: searchParams.getAll("compatibility"),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchCooler(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        // Fetch filters for Cooler category (only once on mount)
        FetchComponentFilters<CoolerSpecs>("cooler")
            .then(setFilters)
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="Cooler Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="cooler"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Cooler;
