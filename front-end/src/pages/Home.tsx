import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchProducts } from "../api/products";
import { FetchProductFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductParams,
    ProductOrder,
    ProductResponse,
    DefaultSpecs,
} from "../constants/types";

function Home() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] = useState<DefaultSpecs | null>(null);

    const params = useMemo<ProductParams>(() => {
        return {
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
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchProducts(params),
    );

    useEffect(() => {
        execute();
        // Fetch filters
        FetchProductFilters().then(setFilters).catch(console.error);
    }, [params]);

    return (
        <ProductListLayout
            title="Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            totalCount={data?.count ?? null}
            currentPage={params.page || 1}
            filters={filters}
        />
    );
}

export default Home;
