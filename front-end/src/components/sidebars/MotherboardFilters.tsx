import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import type { MotherboardSpecs } from "../../constants/types";

type MotherboardFiltersProps = {
    specs: MotherboardSpecs;
};

function MotherboardFilters({ specs }: MotherboardFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        chipset: boolean;
        socket: boolean;
        formFactor: boolean;
    }>({
        chipset: false,
        socket: false,
        formFactor: false,
    });

    const toggleSection = (section: "chipset" | "socket" | "formFactor") => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Chipset Section */}
            {specs.Chipset && specs.Chipset.length > 0 && (
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
                            items={specs.Chipset}
                            paramName="chipset"
                            label="Chipset"
                        />
                    )}
                </div>
            )}

            {/* Socket Section */}
            {specs.Socket && specs.Socket.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("socket")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Socket
                        </h4>
                        <span className="text-gray-500">
                            {openSections.socket ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.socket && (
                        <ScrollableCheckboxList
                            items={specs.Socket}
                            paramName="socket"
                            label="Socket"
                        />
                    )}
                </div>
            )}

            {/* Form Factor Section */}
            {specs.FormFactor && specs.FormFactor.length > 0 && (
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
                            items={specs.FormFactor}
                            paramName="form_factor"
                            label="Form Factor"
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default MotherboardFilters;
