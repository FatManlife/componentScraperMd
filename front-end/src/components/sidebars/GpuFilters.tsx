import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import RangeSlider from "../RangeSlider";
import type { GpuSpecs } from "../../constants/types";

type GpuFiltersProps = {
    specs: GpuSpecs;
};

function GpuFilters({ specs }: GpuFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        chipset: boolean;
        vram: boolean;
        gpu_frequency: boolean;
        vram_frequency: boolean;
    }>({
        chipset: false,
        vram: false,
        gpu_frequency: false,
        vram_frequency: false,
    });

    const toggleSection = (
        section: "chipset" | "vram" | "gpu_frequency" | "vram_frequency",
    ) => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Chipset Section */}
            {specs.chipset && specs.chipset.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("chipset")}
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
                            Chipset
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.chipset ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.chipset && (
                        <ScrollableCheckboxList
                            items={specs.chipset}
                            paramName="chipset"
                            label="Chipset"
                        />
                    )}
                </div>
            )}

            {/* VRAM Section */}
            {specs.vram && specs.vram.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("vram")}
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
                            VRAM
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.vram ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.vram && (
                        <RangeSlider
                            values={specs.vram}
                            label="VRAM"
                            text="VRAM"
                            minParamKey="min_vram"
                            maxParamKey="max_vram"
                            formatValue={(value) => `${value.toFixed(0)} GB`}
                        />
                    )}
                </div>
            )}

            {/* GPU Frequency Section */}
            {specs.gpu_frequency && specs.gpu_frequency.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("gpu_frequency")}
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
                            GPU Frequency
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.gpu_frequency ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.gpu_frequency && (
                        <RangeSlider
                            values={specs.gpu_frequency}
                            label="GPU Frequency"
                            text="GPU Frequency"
                            minParamKey="min_gpu_frequency"
                            maxParamKey="max_gpu_frequency"
                            formatValue={(value) => `${value.toFixed(0)} MHz`}
                        />
                    )}
                </div>
            )}

            {/* VRAM Frequency Section */}
            {specs.vram_frequency && specs.vram_frequency.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("vram_frequency")}
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
                            VRAM Frequency
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.vram_frequency ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.vram_frequency && (
                        <RangeSlider
                            values={specs.vram_frequency}
                            label="VRAM Frequency"
                            text="VRAM Frequency"
                            minParamKey="min_vram_frequency"
                            maxParamKey="max_vram_frequency"
                            formatValue={(value) => `${value.toFixed(0)} MHz`}
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default GpuFilters;
