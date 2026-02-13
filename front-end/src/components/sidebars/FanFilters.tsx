import { useState } from "react";
import RangeSlider from "../RangeSlider";
import type { FanSpecs } from "../../constants/types";

type FanFiltersProps = {
    specs: FanSpecs;
};

function FanFilters({ specs }: FanFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        fan_rpm: boolean;
        noise: boolean;
    }>({
        fan_rpm: false,
        noise: false,
    });

    const toggleSection = (section: "fan_rpm" | "noise") => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Fan RPM Section */}
            {specs.FanRPM && specs.FanRPM.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("fan_rpm")}
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
                            Fan RPM
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.fan_rpm ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.fan_rpm && (
                        <RangeSlider
                            values={specs.FanRPM}
                            label="Fan RPM"
                            text="Fan RPM"
                            minParamKey="min_fan_rpm"
                            maxParamKey="max_fan_rpm"
                            formatValue={(value) => `${value.toFixed(0)} RPM`}
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
                            Noise Level
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.noise ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.noise && (
                        <RangeSlider
                            values={specs.Noise}
                            label="Noise Level"
                            text="Noise"
                            minParamKey="min_noise"
                            maxParamKey="max_noise"
                            formatValue={(value) => `${value.toFixed(1)} dB`}
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default FanFilters;
