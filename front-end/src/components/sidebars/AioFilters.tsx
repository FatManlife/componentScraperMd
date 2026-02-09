import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import type { AioSpecs } from "../../constants/types";

type AioFiltersProps = {
    specs: AioSpecs;
};

function AioFilters({ specs }: AioFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        diagonal: boolean;
        cpu: boolean;
        ram: boolean;
        storage: boolean;
        gpu: boolean;
    }>({
        diagonal: false,
        cpu: false,
        ram: false,
        storage: false,
        gpu: false,
    });

    const toggleSection = (
        section: "diagonal" | "cpu" | "ram" | "storage" | "gpu",
    ) => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Diagonal Section */}
            {specs.Diagonal && specs.Diagonal.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("diagonal")}
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            Diagonal (inches)
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
                            {openSections.diagonal ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.diagonal && (
                        <ScrollableCheckboxList
                            items={specs.Diagonal.map((d) => d.toString())}
                            paramName="diagonal"
                            label="Diagonal"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item}"`
                            }
                        />
                    )}
                </div>
            )}

            {/* CPU Section */}
            {specs.Cpu && specs.Cpu.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("cpu")}
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            CPU
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
                            {openSections.cpu ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.cpu && (
                        <ScrollableCheckboxList
                            items={specs.Cpu}
                            paramName="cpu"
                            label="CPU"
                        />
                    )}
                </div>
            )}

            {/* RAM Section */}
            {specs.Ram && specs.Ram.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("ram")}
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            RAM (GB)
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
                            {openSections.ram ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.ram && (
                        <ScrollableCheckboxList
                            items={specs.Ram.map((r) => r.toString())}
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
            {specs.Storage && specs.Storage.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("storage")}
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            Storage (GB)
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
                            {openSections.storage ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.storage && (
                        <ScrollableCheckboxList
                            items={specs.Storage.map((s) => s.toString())}
                            paramName="storage"
                            label="Storage"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item} GB`
                            }
                        />
                    )}
                </div>
            )}

            {/* GPU Section */}
            {specs.Gpu && specs.Gpu.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("gpu")}
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            GPU
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
                            {openSections.gpu ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.gpu && (
                        <ScrollableCheckboxList
                            items={specs.Gpu}
                            paramName="gpu"
                            label="GPU"
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default AioFilters;
