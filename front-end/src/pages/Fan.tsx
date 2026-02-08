import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchFan } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    FanParams,
    ProductOrder,
    ComponentFiltersResponse,
    FanSpecs,
} from "../constants/types";

function Fan() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<FanSpecs> | null>(null);

    const params = useMemo<FanParams>(() => {
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
            min_fan_rpm: searchParams.get("min_fan_rpm")
                ? parseInt(searchParams.get("min_fan_rpm")!, 10)
                : undefined,
            max_fan_rpm: searchParams.get("max_fan_rpm")
                ? parseInt(searchParams.get("max_fan_rpm")!, 10)
                : undefined,
            min_noise: searchParams.get("min_noise")
                ? parseFloat(searchParams.get("min_noise")!)
                : undefined,
            max_noise: searchParams.get("max_noise")
                ? parseFloat(searchParams.get("max_noise")!)
                : undefined,
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchFan(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        FetchComponentFilters<FanSpecs>("fan")
            .then((data) => {
                setFilters(data);
            })
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="Fan Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="fan"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Fan;
