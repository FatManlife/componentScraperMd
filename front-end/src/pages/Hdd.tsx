import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchHdd } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    HddParams,
    ProductOrder,
    ComponentFiltersResponse,
    HddSpecs,
} from "../constants/types";

function Hdd() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<HddSpecs> | null>(null);

    const params = useMemo<HddParams>(() => {
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
            min_rotation_speed: searchParams.get("min_rotation_speed")
                ? parseInt(searchParams.get("min_rotation_speed")!, 10)
                : undefined,
            max_rotation_speed: searchParams.get("max_rotation_speed")
                ? parseInt(searchParams.get("max_rotation_speed")!, 10)
                : undefined,
            form_factor: searchParams.getAll("form_factor"),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchHdd(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        FetchComponentFilters<HddSpecs>("hdd")
            .then((data) => {
                setFilters(data);
            })
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="HDD Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="hdd"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Hdd;
