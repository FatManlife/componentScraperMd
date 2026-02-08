import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import type { CoolerSpecs } from "../../constants/types";

type CoolerFiltersProps = {
    specs: CoolerSpecs;
};

function CoolerFilters({ specs }: CoolerFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        type: boolean;
        fanRPM: boolean;
        noise: boolean;
        compatibility: boolean;
    }>({
        type: false,
        fanRPM: false,
        noise: false,
        compatibility: false,
    });

    const toggleSection = (
        section: "type" | "fanRPM" | "noise" | "compatibility",
    ) => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Type Section */}
            {specs.Type && specs.Type.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("type")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Type
                        </h4>
                        <span className="text-gray-500">
                            {openSections.type ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.type && (
                        <ScrollableCheckboxList
                            items={specs.Type}
                            paramName="type"
                            label="Type"
                        />
                    )}
                </div>
            )}

            {/* Fan RPM Section */}
            {specs.FanRPM && specs.FanRPM.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("fanRPM")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Fan RPM
                        </h4>
                        <span className="text-gray-500">
                            {openSections.fanRPM ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.fanRPM && (
                        <ScrollableCheckboxList
                            items={specs.FanRPM.map((r) => r.toString())}
                            paramName="fan_rpm"
                            label="Fan RPM"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item} RPM`
                            }
                        />
                    )}
                </div>
            )}

            {/* Noise Section */}
            {specs.Noise && specs.Noise.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("noise")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Noise Level (dB)
                        </h4>
                        <span className="text-gray-500">
                            {openSections.noise ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.noise && (
                        <ScrollableCheckboxList
                            items={specs.Noise.map((n) => n.toString())}
                            paramName="noise"
                            label="Noise"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item} dB`
                            }
                        />
                    )}
                </div>
            )}

            {/* Compatibility Section */}
            {specs.Compatibility && specs.Compatibility.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("compatibility")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Compatibility
                        </h4>
                        <span className="text-gray-500">
                            {openSections.compatibility ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.compatibility && (
                        <ScrollableCheckboxList
                            items={specs.Compatibility}
                            paramName="compatibility"
                            label="Compatibility"
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default CoolerFilters;
