import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchPcMini } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    PcMiniParams,
    ProductOrder,
    ComponentFiltersResponse,
    PcMiniSpecs,
} from "../constants/types";

function PcMini() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<PcMiniSpecs> | null>(null);

    const params = useMemo<PcMiniParams>(() => {
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
            cpu: searchParams.getAll("cpu"),
            gpu: searchParams.getAll("gpu"),
            ram: searchParams
                .getAll("ram")
                .map((r) => parseInt(r, 10))
                .filter((r) => !isNaN(r)),
            storage: searchParams
                .getAll("storage")
                .map((s) => parseInt(s, 10))
                .filter((s) => !isNaN(s)),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchPcMini(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        FetchComponentFilters<PcMiniSpecs>("pcmini")
            .then((data) => {
                setFilters(data);
            })
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="PC Mini Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="pcmini"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default PcMini;
