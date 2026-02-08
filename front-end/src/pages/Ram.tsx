import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchRam } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    RamParams,
    ProductOrder,
    ComponentFiltersResponse,
    RamSpecs,
} from "../constants/types";

function Ram() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<RamSpecs> | null>(null);

    const params = useMemo<RamParams>(() => {
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
            min_capacity: searchParams.get("min_capacity")
                ? parseInt(searchParams.get("min_capacity")!, 10)
                : undefined,
            max_capacity: searchParams.get("max_capacity")
                ? parseInt(searchParams.get("max_capacity")!, 10)
                : undefined,
            min_speed: searchParams.get("min_speed")
                ? parseInt(searchParams.get("min_speed")!, 10)
                : undefined,
            max_speed: searchParams.get("max_speed")
                ? parseInt(searchParams.get("max_speed")!, 10)
                : undefined,
            type: searchParams.getAll("type"),
            compatibility: searchParams.getAll("compatibility"),
            configuration: searchParams
                .getAll("configuration")
                .map((c) => parseInt(c, 10))
                .filter((c) => !isNaN(c)),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchRam(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        FetchComponentFilters<RamSpecs>("ram")
            .then((data) => {
                setFilters(data);
            })
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="RAM Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="ram"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Ram;
