import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import RangeSlider from "../RangeSlider";
import type { HddSpecs } from "../../constants/types";

type HddFiltersProps = {
    specs: HddSpecs;
};

function HddFilters({ specs }: HddFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        capacity: boolean;
        rotationSpeed: boolean;
        formFactor: boolean;
    }>({
        capacity: false,
        rotationSpeed: false,
        formFactor: false,
    });

    const toggleSection = (
        section: "capacity" | "rotationSpeed" | "formFactor",
    ) => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Capacity Section */}
            {specs.capacity && specs.capacity.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("capacity")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Capacity
                        </h4>
                        <span className="text-gray-500">
                            {openSections.capacity ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.capacity && (
                        <RangeSlider
                            values={specs.capacity}
                            label="Capacity"
                            text="Capacity"
                            minParamKey="min_capacity"
                            maxParamKey="max_capacity"
                            formatValue={(value) => `${value.toFixed(0)} GB`}
                        />
                    )}
                </div>
            )}

            {/* Rotation Speed Section */}
            {specs.rotation_speed && specs.rotation_speed.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("rotationSpeed")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Rotation Speed
                        </h4>
                        <span className="text-gray-500">
                            {openSections.rotationSpeed ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.rotationSpeed && (
                        <RangeSlider
                            values={specs.rotation_speed}
                            label="Rotation Speed"
                            text="Rotation Speed"
                            minParamKey="min_rotation_speed"
                            maxParamKey="max_rotation_speed"
                            formatValue={(value) => `${value.toFixed(0)} RPM`}
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

export default HddFilters;
