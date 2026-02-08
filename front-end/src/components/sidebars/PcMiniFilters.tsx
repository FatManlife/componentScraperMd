import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import type { PcMiniSpecs } from "../../constants/types";

type PcMiniFiltersProps = {
    specs: PcMiniSpecs;
};

function PcMiniFilters({ specs }: PcMiniFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        cpu: boolean;
        gpu: boolean;
        ram: boolean;
        storage: boolean;
    }>({
        cpu: false,
        gpu: false,
        ram: false,
        storage: false,
    });

    const toggleSection = (section: "cpu" | "gpu" | "ram" | "storage") => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* CPU Section */}
            {specs.cpu && specs.cpu.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("cpu")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            CPU
                        </h4>
                        <span className="text-gray-500">
                            {openSections.cpu ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.cpu && (
                        <ScrollableCheckboxList
                            items={specs.cpu}
                            paramName="cpu"
                            label="CPU"
                        />
                    )}
                </div>
            )}

            {/* GPU Section */}
            {specs.gpu && specs.gpu.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("gpu")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            GPU
                        </h4>
                        <span className="text-gray-500">
                            {openSections.gpu ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.gpu && (
                        <ScrollableCheckboxList
                            items={specs.gpu}
                            paramName="gpu"
                            label="GPU"
                        />
                    )}
                </div>
            )}

            {/* RAM Section */}
            {specs.ram && specs.ram.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("ram")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            RAM (GB)
                        </h4>
                        <span className="text-gray-500">
                            {openSections.ram ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.ram && (
                        <ScrollableCheckboxList
                            items={specs.ram.map((r) => r.toString())}
                            paramName="ram"
                            label="RAM"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item} GB`
                            }
                        />
                    )}
                </div>
            )}

            {/* Storage Section */}
            {specs.storage && specs.storage.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("storage")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Storage (GB)
                        </h4>
                        <span className="text-gray-500">
                            {openSections.storage ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.storage && (
                        <ScrollableCheckboxList
                            items={specs.storage.map((s) => s.toString())}
                            paramName="storage"
                            label="Storage"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item} GB`
                            }
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default PcMiniFilters;
