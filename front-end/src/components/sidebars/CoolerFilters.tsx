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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            Type
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            Fan RPM
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            Noise Level (dB)
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
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
                        className="flex items-center justify-between w-full text-left mb-2 p-2 transition-colors"
                        style={{ backgroundColor: 'transparent' }}
                        onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                        onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
                    >
                        <h4 className="text-sm font-semibold" style={{ color: '#000000' }}>
                            Compatibility
                        </h4>
                        <span style={{ color: '#8A8A8A' }}>
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
