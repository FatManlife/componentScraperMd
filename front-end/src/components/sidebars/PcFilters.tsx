import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import type { PcSpecs } from "../../constants/types";

type PcFiltersProps = {
    specs: PcSpecs;
};

function PcFilters({ specs }: PcFiltersProps) {
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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: "transparent" }}
                        onMouseEnter={(e) =>
                            (e.currentTarget.style.backgroundColor = "#F4F4F4")
                        }
                        onMouseLeave={(e) =>
                            (e.currentTarget.style.backgroundColor =
                                "transparent")
                        }
                    >
                        <h4
                            className="text-sm font-semibold"
                            style={{ color: "#000000" }}
                        >
                            CPU
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: "transparent" }}
                        onMouseEnter={(e) =>
                            (e.currentTarget.style.backgroundColor = "#F4F4F4")
                        }
                        onMouseLeave={(e) =>
                            (e.currentTarget.style.backgroundColor =
                                "transparent")
                        }
                    >
                        <h4
                            className="text-sm font-semibold"
                            style={{ color: "#000000" }}
                        >
                            GPU
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: "transparent" }}
                        onMouseEnter={(e) =>
                            (e.currentTarget.style.backgroundColor = "#F4F4F4")
                        }
                        onMouseLeave={(e) =>
                            (e.currentTarget.style.backgroundColor =
                                "transparent")
                        }
                    >
                        <h4
                            className="text-sm font-semibold"
                            style={{ color: "#000000" }}
                        >
                            RAM (GB)
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: "transparent" }}
                        onMouseEnter={(e) =>
                            (e.currentTarget.style.backgroundColor = "#F4F4F4")
                        }
                        onMouseLeave={(e) =>
                            (e.currentTarget.style.backgroundColor =
                                "transparent")
                        }
                    >
                        <h4
                            className="text-sm font-semibold"
                            style={{ color: "#000000" }}
                        >
                            Storage (GB)
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
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

export default PcFilters;
