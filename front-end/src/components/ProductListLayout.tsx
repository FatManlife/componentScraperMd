import type { ReactNode } from "react";
import Pagination from "./Pagination";
import ProductList from "./ProductList";
import Sidebar from "./Sidebar";
import AioFilters from "./sidebars/AioFilters";
import CaseFilters from "./sidebars/CaseFilters";
import CoolerFilters from "./sidebars/CoolerFilters";
import CpuFilters from "./sidebars/CpuFilters";
import FanFilters from "./sidebars/FanFilters";
import GpuFilters from "./sidebars/GpuFilters";
import HddFilters from "./sidebars/HddFilters";
import LaptopFilters from "./sidebars/LaptopFilters";
import MotherboardFilters from "./sidebars/MotherboardFilters";
import PcFilters from "./sidebars/PcFilters";
import PcMiniFilters from "./sidebars/PcMiniFilters";
import PsuFilters from "./sidebars/PsuFilters";
import RamFilters from "./sidebars/RamFilters";
import SsdFilters from "./sidebars/SsdFilters";
import type {
    Products,
    DefaultSpecs,
    AioSpecs,
    CaseSpecs,
    CoolerSpecs,
    CpuSpecs,
    FanSpecs,
    GpuSpecs,
    HddSpecs,
    LaptopSpecs,
    MotherboardSpecs,
    PcSpecs,
    PcMiniSpecs,
    PsuSpecs,
    RamSpecs,
    SsdSpecs,
} from "../constants/types";

type ProductListLayoutProps = {
    title: string;
    loading: boolean;
    error: string | null;
    data: Products[] | null;
    currentPage: number;
    totalCount: number | null;
    filters: DefaultSpecs | null;
    category?: string | null;
    specificSpecs?:
        | AioSpecs
        | CaseSpecs
        | CoolerSpecs
        | CpuSpecs
        | FanSpecs
        | GpuSpecs
        | HddSpecs
        | LaptopSpecs
        | MotherboardSpecs
        | PcSpecs
        | PcMiniSpecs
        | PsuSpecs
        | RamSpecs
        | SsdSpecs
        | null;
    children?: ReactNode;
};

function ProductListLayout({
    title,
    loading,
    error,
    data,
    currentPage,
    totalCount,
    filters,
    category,
    specificSpecs,
    children,
}: ProductListLayoutProps) {
    const renderSpecificFilters = () => {
        if (!category || !specificSpecs) {
            return null;
        }

        switch (category) {
            case "aio":
                return <AioFilters specs={specificSpecs as AioSpecs} />;
            case "case":
                return <CaseFilters specs={specificSpecs as CaseSpecs} />;
            case "cooler":
                return <CoolerFilters specs={specificSpecs as CoolerSpecs} />;
            case "cpu":
                return <CpuFilters specs={specificSpecs as CpuSpecs} />;
            case "fan":
                return <FanFilters specs={specificSpecs as FanSpecs} />;
            case "gpu":
                return <GpuFilters specs={specificSpecs as GpuSpecs} />;
            case "hdd":
                return <HddFilters specs={specificSpecs as HddSpecs} />;
            case "laptop":
                return <LaptopFilters specs={specificSpecs as LaptopSpecs} />;
            case "motherboard":
                return (
                    <MotherboardFilters
                        specs={specificSpecs as MotherboardSpecs}
                    />
                );
            case "pc":
                return <PcFilters specs={specificSpecs as PcSpecs} />;
            case "pcmini":
                return <PcMiniFilters specs={specificSpecs as PcMiniSpecs} />;
            case "psu":
                return <PsuFilters specs={specificSpecs as PsuSpecs} />;
            case "ram":
                return <RamFilters specs={specificSpecs as RamSpecs} />;
            case "ssd":
                return <SsdFilters specs={specificSpecs as SsdSpecs} />;
            default:
                return null;
        }
    };

    return (
        <div className="bg-gray-100 py-8 px-4">
            <div className="max-w-6xl mx-auto">
                <h2 className="text-3xl font-bold text-gray-800 mb-6">
                    {title}
                </h2>

                {children}

                <div className="flex gap-6">
                    <Sidebar
                        filters={filters}
                        specificFilters={renderSpecificFilters()}
                    />

                    <div className="flex-1 flex flex-col">
                        <div className="min-h-200">
                            {loading && (
                                <div className="text-center text-gray-600 py-20">
                                    Loading...
                                </div>
                            )}

                            {error && (
                                <div className="mt-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
                                    {error}
                                </div>
                            )}

                            {data && !loading && (
                                <ProductList products={data} />
                            )}
                        </div>

                        <div className="h-16 flex items-center justify-center mt-6">
                            {totalCount !== null && (
                                <Pagination
                                    currentPage={currentPage}
                                    totalCount={totalCount}
                                    itemsPerPage={24}
                                />
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default ProductListLayout;
