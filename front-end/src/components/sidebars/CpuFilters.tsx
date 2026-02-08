import { useState } from "react";
import ScrollableCheckboxList from "../ScrollableCheckboxList";
import type { CpuSpecs } from "../../constants/types";

type CpuFiltersProps = {
    specs: CpuSpecs;
};

function CpuFilters({ specs }: CpuFiltersProps) {
    const [openSections, setOpenSections] = useState<{
        cores: boolean;
        threads: boolean;
        baseClock: boolean;
        boostClock: boolean;
        socket: boolean;
    }>({
        cores: false,
        threads: false,
        baseClock: false,
        boostClock: false,
        socket: false,
    });

    const toggleSection = (
        section: "cores" | "threads" | "baseClock" | "boostClock" | "socket",
    ) => {
        setOpenSections((prev) => ({
            ...prev,
            [section]: !prev[section],
        }));
    };

    return (
        <>
            {/* Cores Section */}
            {specs.Cores && specs.Cores.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("cores")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Cores
                        </h4>
                        <span className="text-gray-500">
                            {openSections.cores ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.cores && (
                        <ScrollableCheckboxList
                            items={specs.Cores.map((c) => c.toString())}
                            paramName="cores"
                            label="Cores"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : item
                            }
                        />
                    )}
                </div>
            )}

            {/* Threads Section */}
            {specs.Threads && specs.Threads.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("threads")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Threads
                        </h4>
                        <span className="text-gray-500">
                            {openSections.threads ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.threads && (
                        <ScrollableCheckboxList
                            items={specs.Threads.map((t) => t.toString())}
                            paramName="threads"
                            label="Threads"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : item
                            }
                        />
                    )}
                </div>
            )}

            {/* Base Clock Section */}
            {specs.BaseClock && specs.BaseClock.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("baseClock")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Base Clock (GHz)
                        </h4>
                        <span className="text-gray-500">
                            {openSections.baseClock ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.baseClock && (
                        <ScrollableCheckboxList
                            items={specs.BaseClock.map((b) => b.toString())}
                            paramName="base_clock"
                            label="Base Clock"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item} GHz`
                            }
                        />
                    )}
                </div>
            )}

            {/* Boost Clock Section */}
            {specs.BoostClock && specs.BoostClock.length > 0 && (
                <div className="mb-6">
                    <button
                        onClick={() => toggleSection("boostClock")}
                        className="flex items-center justify-between w-full text-left mb-2 hover:bg-gray-50 p-2 rounded"
                    >
                        <h4 className="text-sm font-semibold text-gray-700">
                            Boost Clock (GHz)
                        </h4>
                        <span className="text-gray-500">
                            {openSections.boostClock ? "−" : "+"}
                        </span>
                    </button>
                    {openSections.boostClock && (
                        <ScrollableCheckboxList
                            items={specs.BoostClock.map((b) => b.toString())}
                            paramName="boost_clock"
                            label="Boost Clock"
                            formatLabel={(item) =>
                                item === "0" ? "Unk" : `${item} GHz`
                            }
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
        </>
    );
}

export default CpuFilters;
