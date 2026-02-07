import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchAio } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    AioParams,
    ProductOrder,
    ComponentFiltersResponse,
    AioSpecs,
} from "../constants/types";

function Aio() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<AioSpecs> | null>(null);

    const params = useMemo<AioParams>(() => {
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
            diagonal: searchParams
                .getAll("diagonal")
                .map((d) => parseFloat(d))
                .filter((d) => !isNaN(d)),
            cpu: searchParams.getAll("cpu"),
            ram: searchParams
                .getAll("ram")
                .map((r) => parseInt(r, 10))
                .filter((r) => !isNaN(r)),
            storage: searchParams
                .getAll("storage")
                .map((s) => parseInt(s, 10))
                .filter((s) => !isNaN(s)),
            gpu: searchParams.getAll("gpu"),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchAio(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        // Fetch filters for AIO category (only once on mount)
        FetchComponentFilters<AioSpecs>("aio")
            .then((data) => {
                console.log("AIO Filters received:", data);
                setFilters(data);
            })
            .catch(console.error);
    }, []);

    console.log("Current filters state:", filters);
    console.log(
        "Passing to layout - category:",
        "aio",
        "specificSpecs:",
        filters?.specificSpecs,
    );

    return (
        <ProductListLayout
            title="AIO Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="aio"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Aio;
