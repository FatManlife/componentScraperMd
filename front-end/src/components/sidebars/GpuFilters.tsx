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
    }>({
        chipset: false,
    });

    const toggleSection = (section: "chipset") => {
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
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Chipset
                        </h4>
                        <span className="text-gray-500">
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
                <RangeSlider
                    values={specs.vram}
                    label="VRAM"
                    text="VRAM"
                    minParamKey="min_vram"
                    maxParamKey="max_vram"
                    formatValue={(value) => `${value.toFixed(0)} GB`}
                />
            )}

            {/* GPU Frequency Section */}
            {specs.gpu_frequency && specs.gpu_frequency.length > 0 && (
                <RangeSlider
                    values={specs.gpu_frequency}
                    label="GPU Frequency"
                    text="GPU Frequency"
                    minParamKey="min_gpu_frequency"
                    maxParamKey="max_gpu_frequency"
                    formatValue={(value) => `${value.toFixed(0)} MHz`}
                />
            )}

            {/* VRAM Frequency Section */}
            {specs.vram_frequency && specs.vram_frequency.length > 0 && (
                <RangeSlider
                    values={specs.vram_frequency}
                    label="VRAM Frequency"
                    text="VRAM Frequency"
                    minParamKey="min_vram_frequency"
                    maxParamKey="max_vram_frequency"
                    formatValue={(value) => `${value.toFixed(0)} MHz`}
                />
            )}
        </>
    );
}

export default GpuFilters;
