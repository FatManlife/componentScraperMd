import { useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { FetchMotherboard } from "../api/components";
import { FetchComponentFilters } from "../api/filters";
import { useFetch } from "../hooks/useFetch";
import ProductListLayout from "../components/ProductListLayout";
import type {
    ProductResponse,
    MotherboardParams,
    ProductOrder,
    ComponentFiltersResponse,
    MotherboardSpecs,
} from "../constants/types";

function Motherboard() {
    const [searchParams] = useSearchParams();
    const [filters, setFilters] =
        useState<ComponentFiltersResponse<MotherboardSpecs> | null>(null);

    const params = useMemo<MotherboardParams>(() => {
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
            socket: searchParams.getAll("socket"),
            form_factor: searchParams.getAll("form_factor"),
        };
    }, [searchParams.toString()]);

    const { data, loading, error, execute } = useFetch<ProductResponse>(() =>
        FetchMotherboard(params),
    );

    useEffect(() => {
        execute();
    }, [params]);

    useEffect(() => {
        FetchComponentFilters<MotherboardSpecs>("motherboard")
            .then((data) => {
                setFilters(data);
            })
            .catch(console.error);
    }, []);

    return (
        <ProductListLayout
            title="Motherboard Products"
            loading={loading}
            error={error}
            data={data?.products ?? []}
            currentPage={params.defaultParams.page || 1}
            totalCount={data?.count ?? null}
            filters={filters?.defaultSpecs ?? null}
            category="motherboard"
            specificSpecs={filters?.specificSpecs ?? null}
        />
    );
}

export default Motherboard;
