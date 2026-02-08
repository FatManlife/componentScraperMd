import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import RangeSlider from "../RangeSlider";
import type { PsuSpecs } from "../../constants/types";

type PsuFiltersProps = {
    specs: PsuSpecs;
};

function PsuFilters({ specs }: PsuFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        power: boolean;
        efficiency: boolean;
        formFactor: boolean;
    }>({
        power: false,
        efficiency: false,
        formFactor: false,
    });

    const toggleSection = (section: "power" | "efficiency" | "formFactor") => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Power Section */}
            {specs.power && specs.power.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("power")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Power
                        </h4>
                        <span className="text-gray-500">
                            {openSections.power ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.power && (
                        <RangeSlider
                            values={specs.power}
                            label="Power"
                            text="Power"
                            minParamKey="min_power"
                            maxParamKey="max_power"
                            formatValue={(value) => `${value.toFixed(0)} W`}
                        />
                    )}
                </div>
            )}

            {/* Efficiency Section */}
            {specs.efficiency && specs.efficiency.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("efficiency")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Efficiency
                        </h4>
                        <span className="text-gray-500">
                            {openSections.efficiency ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.efficiency && (
                        <ScrollableCheckboxList
                            items={specs.efficiency}
                            paramName="efficiency"
                            label="Efficiency"
                        />
                    )}
                </div>
            )}

            {/* Form Factor Section */}
            {specs.form_factor && specs.form_factor.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("formFactor")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Form Factor
                        </h4>
                        <span className="text-gray-500">
                            {openSections.formFactor ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.formFactor && (
                        <ScrollableCheckboxList
                            items={specs.form_factor}
                            paramName="form_factor"
                            label="Form Factor"
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default PsuFilters;
