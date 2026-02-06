import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchAio } from "../api/components";
import { FetchFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    AioParams,
    ProductOrder,
    DefaultFilters,
} from "../constants/types";

function Aio() {
    const [searchParams] = useSearchParams();
    //const [totalCount, setTotalCount] = useState<number | null>(null);
    const [filters, setFilters] = useState<DefaultFilters | null>(null);

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
            diagonal: searchParams.getAll("diagonal"),
            cpu: searchParams.getAll("cpu"),
            ram: searchParams.getAll("ram"),
            storage: searchParams.getAll("storage"),
            gpu: searchParams.getAll("gpu"),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchAio(params),
    );

    useEffect(() => {
        execute();
        // Fetch count for AIO category
        // Fetch filters for AIO category
        FetchFilters("aio").then(setFilters).catch(console.error);
    }, [params]);

    return (
        // <ProductListLayout
        //     title="AIO Products"
        //     loading={loading}
        //     error={error}
        //     data={data}
        //     currentPage={params.defaultParams.page || 1}
        //     totalCount={data?.count ?? null}
        //     filters={filters}
        // />
        <div></div>
    );
}

export default Aio;
