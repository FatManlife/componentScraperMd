import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import RangeSlider from "../RangeSlider";
import type { SsdSpecs } from "../../constants/types";

type SsdFiltersProps = {
    specs: SsdSpecs;
};

function SsdFilters({ specs }: SsdFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        capacity: boolean;
        readingSpeed: boolean;
        writingSpeed: boolean;
        formFactor: boolean;
    }>({
        capacity: false,
        readingSpeed: false,
        writingSpeed: false,
        formFactor: false,
    });

    const toggleSection = (
        section: "capacity" | "readingSpeed" | "writingSpeed" | "formFactor",
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

            {/* Reading Speed Section */}
            {specs.reading_speed && specs.reading_speed.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("readingSpeed")}
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
                            Reading Speed
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.readingSpeed ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.readingSpeed && (
                        <RangeSlider
                            values={specs.reading_speed}
                            label="Reading Speed"
                            text="Reading Speed"
                            minParamKey="min_reading_speed"
                            maxParamKey="max_reading_speed"
                            formatValue={(value) => `${value.toFixed(0)} MB/s`}
                        />
                    )}
                </div>
            )}

            {/* Writing Speed Section */}
            {specs.writing_speed && specs.writing_speed.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("writingSpeed")}
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
                            Writing Speed
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
                            {openSections.writingSpeed ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.writingSpeed && (
                        <RangeSlider
                            values={specs.writing_speed}
                            label="Writing Speed"
                            text="Writing Speed"
                            minParamKey="min_writing_speed"
                            maxParamKey="max_writing_speed"
                            formatValue={(value) => `${value.toFixed(0)} MB/s`}
                        />
                    )}
                </div>
            )}

            {/* Form Factor Section */}
            {specs.form_factor && specs.form_factor.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("formFactor")}
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
                            Form Factor
                        </h4>
                        <span style={{ color: "#8A8A8A" }}>
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

export default SsdFilters;
