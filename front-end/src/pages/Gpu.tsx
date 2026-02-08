import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchGpu } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    GpuParams,
    ProductOrder,
    ComponentFiltersResponse,
    GpuSpecs,
} from "../constants/types";

function Gpu() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<GpuSpecs> | null>(null);

    const params = useMemo<GpuParams>(() => {
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
            chipset: searchParams.getAll("chipset"),
            min_vram: searchParams.get("min_vram")
                ? parseInt(searchParams.get("min_vram")!, 10)
                : undefined,
            max_vram: searchParams.get("max_vram")
                ? parseInt(searchParams.get("max_vram")!, 10)
                : undefined,
            min_gpu_frequency: searchParams.get("min_gpu_frequency")
                ? parseInt(searchParams.get("min_gpu_frequency")!, 10)
                : undefined,
            max_gpu_frequency: searchParams.get("max_gpu_frequency")
                ? parseInt(searchParams.get("max_gpu_frequency")!, 10)
                : undefined,
            min_vram_frequency: searchParams.get("min_vram_frequency")
                ? parseInt(searchParams.get("min_vram_frequency")!, 10)
                : undefined,
            max_vram_frequency: searchParams.get("max_vram_frequency")
                ? parseInt(searchParams.get("max_vram_frequency")!, 10)
                : undefined,
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchGpu(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        FetchComponentFilters<GpuSpecs>("gpu")
            .then((data) => {
                setFilters(data);
            })
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="GPU Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="gpu"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Gpu;
