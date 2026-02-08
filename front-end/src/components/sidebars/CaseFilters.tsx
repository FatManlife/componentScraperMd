import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import type { CaseSpecs } from "../../constants/types";

type CaseFiltersProps = {
    specs: CaseSpecs;
};

function CaseFilters({ specs }: CaseFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        format: boolean;
        motherboardFormFactor: boolean;
    }>({
        format: false,
        motherboardFormFactor: false,
    });

    const toggleSection = (section: "format" | "motherboardFormFactor") => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Format Section */}
            {specs.Format && specs.Format.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("format")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Format
                        </h4>
                        <span className="text-gray-500">
                            {openSections.format ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.format && (
                        <ScrollableCheckboxList
                            items={specs.Format}
                            paramName="format"
                            label="Format"
                        />
                    )}
                </div>
            )}

            {/* Motherboard Form Factor Section */}
            {specs.MotherboardFormFactor &&
                specs.MotherboardFormFactor.length > 0 && (
                    <div className="mb-6">
                        <button
                            onClick={() =>
                                toggleSection("motherboardFormFactor")
                            }
                            className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                        >
                            <h4 className="text-sm font-semibold text-gray-700">
                                Motherboard Form Factor
                            </h4>
                            <span className="text-gray-500">
                                {openSections.motherboardFormFactor ? "−" : "+"}
                            </span>
                        </button>
                        {openSections.motherboardFormFactor && (
                            <ScrollableCheckboxList
                                items={specs.MotherboardFormFactor}
                                paramName="motherboard_form_factor"
                                label="Motherboard Form Factor"
                            />
                        )}
                    </div>
                )}
        </>
    );
}

export default CaseFilters;
