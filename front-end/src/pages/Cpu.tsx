import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchCpu } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    CpuParams,
    ProductOrder,
    ComponentFiltersResponse,
    CpuSpecs,
} from "../constants/types";

function Cpu() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<CpuSpecs> | null>(null);

    const params = useMemo<CpuParams>(() => {
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
            cores: searchParams
                .getAll("cores")
                .map((c) => parseInt(c, 10))
                .filter((c) => !isNaN(c)),
            threads: searchParams
                .getAll("threads")
                .map((t) => parseInt(t, 10))
                .filter((t) => !isNaN(t)),
            base_clock: searchParams
                .getAll("base_clock")
                .map((b) => parseFloat(b))
                .filter((b) => !isNaN(b)),
            boost_clock: searchParams
                .getAll("boost_clock")
                .map((b) => parseFloat(b))
                .filter((b) => !isNaN(b)),
            socket: searchParams.getAll("socket"),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchCpu(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        // Fetch filters for CPU category (only once on mount)
        FetchComponentFilters<CpuSpecs>("cpu")
            .then(setFilters)
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="CPU Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="cpu"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Cpu;
