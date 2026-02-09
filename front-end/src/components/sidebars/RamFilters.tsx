import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import RangeSlider from "../RangeSlider";
import type { RamSpecs } from "../../constants/types";

type RamFiltersProps = {
    specs: RamSpecs;
};

function RamFilters({ specs }: RamFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        capacity: boolean;
        speed: boolean;
        type: boolean;
        compatibility: boolean;
        configuration: boolean;
    }>({
        capacity: false,
        speed: false,
        type: false,
        compatibility: false,
        configuration: false,
    });

    const toggleSection = (
        section:
            | "capacity"
            | "speed"
            | "type"
            | "compatibility"
            | "configuration",
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
                            Capacity
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
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

            {/* Speed Section */}
            {specs.speed && specs.speed.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("speed")}
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
                            Speed
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.speed ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.speed && (
                        <RangeSlider
                            values={specs.speed}
                            label="Speed"
                            text="Speed"
                            minParamKey="min_speed"
                            maxParamKey="max_speed"
                            formatValue={(value) => `${value.toFixed(0)} MHz`}
                        />
                    )}
                </div>
            )}

            {/* Type Section */}
            {specs.type && specs.type.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("type")}
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
                            Type
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.type ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.type && (
                        <ScrollableCheckboxList
                            items={specs.type}
                            paramName="type"
                            label="Type"
                        />
                    )}
                </div>
            )}

            {/* Compatibility Section */}
            {specs.compatibility && specs.compatibility.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("compatibility")}
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
                            Compatibility
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.compatibility ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.compatibility && (
                        <ScrollableCheckboxList
                            items={specs.compatibility}
                            paramName="compatibility"
                            label="Compatibility"
                        />
                    )}
                </div>
            )}

            {/* Configuration Section */}
            {specs.configuration && specs.configuration.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("configuration")}
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
                            Configuration
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.configuration ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.configuration && (
                        <ScrollableCheckboxList
                            items={specs.configuration.map((c) => c.toString())}
                            paramName="configuration"
                            label="Configuration"
                        />
                    )}
                </div>
            )}
        </>
    );
}

export default RamFilters;
