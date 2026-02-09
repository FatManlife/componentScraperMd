import { useState, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

type RangeSliderProps = {
    values: number[];
    label: string;
    text: string;
    minParamKey?: string;
    maxParamKey?: string;
    formatValue?: (value: number) => string;
};

function RangeSlider({
    values,
    minParamKey = "min",
    maxParamKey = "max",
    formatValue = (value) => value.toFixed(0),
}: RangeSliderProps) {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();

    const rangeMin = Math.min(...values);
    const rangeMax = Math.max(...values);

    const [minValue, setMinValue] = useState<number>(rangeMin);
    const [maxValue, setMaxValue] = useState<number>(rangeMax);
    const [isEnabled, setIsEnabled] = useState<boolean>(false);

    useEffect(() => {
        const urlMin = searchParams.get(minParamKey);
        const urlMax = searchParams.get(maxParamKey);

        if (urlMin || urlMax) {
            setIsEnabled(true);
            setMinValue(urlMin ? parseFloat(urlMin) : rangeMin);
            setMaxValue(urlMax ? parseFloat(urlMax) : rangeMax);
        } else {
            setMinValue(rangeMin);
            setMaxValue(rangeMax);
        }
    }, [searchParams, rangeMin, rangeMax, minParamKey, maxParamKey]);

    const handleMinChange = (value: number) => {
        const newMin = Math.min(value, maxValue - 1);
        setMinValue(newMin);
    };

    const handleMaxChange = (value: number) => {
        const newMax = Math.max(value, minValue + 1);
        setMaxValue(newMax);
    };

    const handleMinRelease = () => {
        if (isEnabled) {
            const newParams = new URLSearchParams(searchParams);
            newParams.set(minParamKey, minValue.toString());
            newParams.set(maxParamKey, maxValue.toString());
            newParams.delete("page");
            navigate(`?${newParams.toString()}`, { replace: true });
        }
    };

    const handleMaxRelease = () => {
        if (isEnabled) {
            const newParams = new URLSearchParams(searchParams);
            newParams.set(minParamKey, minValue.toString());
            newParams.set(maxParamKey, maxValue.toString());
            newParams.delete("page");
            navigate(`?${newParams.toString()}`, { replace: true });
        }
    };

    const toggleEnabled = () => {
        const newEnabled = !isEnabled;
        setIsEnabled(newEnabled);

        const newParams = new URLSearchParams(searchParams);
        if (!newEnabled) {
            newParams.delete(minParamKey);
            newParams.delete(maxParamKey);
            setMinValue(rangeMin);
            setMaxValue(rangeMax);
        } else {
            newParams.set(minParamKey, minValue.toString());
            newParams.set(maxParamKey, maxValue.toString());
        }
        navigate(`?${newParams.toString()}`, { replace: true });
    };

    const minPercent = ((minValue - rangeMin) / (rangeMax - rangeMin)) * 100;
    const maxPercent = ((maxValue - rangeMin) / (rangeMax - rangeMin)) * 100;

    return (
        <div className="px-2">
            {/* Enable/Disable Toggle */}
            <label className="flex items-center gap-2 cursor-pointer p-2 mb-3 transition-colors"
                style={{ backgroundColor: 'transparent' }}
                onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#F4F4F4'}
                onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
            >
                <input
                    type="checkbox"
                    checked={isEnabled}
                    onChange={toggleEnabled}
                    className="w-4 h-4"
                />
                <span className="text-sm" style={{ color: '#000000' }}>Allow Filtering</span>
            </label>
            {isEnabled && (
                <div className="space-y-3 px-2">
                    <div className="flex justify-between text-xs" style={{ color: '#8A8A8A' }}>
                        <span>{formatValue(minValue)}</span>
                        <span>{formatValue(maxValue)}</span>
                    </div>

                    <div className="relative h-6">
                        {/* Track background */}
                        <div className="absolute top-1/2 left-0 right-0 h-2 -translate-y-1/2" style={{ backgroundColor: '#D9D9D9', borderRadius: '2px' }} />

                        {/* Active track */}
                        <div
                            className="absolute top-1/2 h-2 -translate-y-1/2"
                            style={{
                                left: `${minPercent}%`,
                                right: `${100 - maxPercent}%`,
                                backgroundColor: '#000000',
                                borderRadius: '2px'
                            }}
                        />

                        {/* Min range input */}
                        <input
                            type="range"
                            min={rangeMin}
                            max={rangeMax}
                            value={minValue}
                            onChange={(e) =>
                                handleMinChange(parseFloat(e.target.value))
                            }
                            onMouseUp={handleMinRelease}
                            onTouchEnd={handleMinRelease}
                            disabled={!isEnabled}
                            className="absolute w-full h-6 appearance-none bg-transparent pointer-events-none [&::-webkit-slider-thumb]:pointer-events-auto [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:w-4 [&::-webkit-slider-thumb]:h-4 [&::-webkit-slider-thumb]:bg-[#000000] [&::-webkit-slider-thumb]:rounded-full [&::-webkit-slider-thumb]:cursor-pointer [&::-webkit-slider-thumb]:shadow [&::-moz-range-thumb]:pointer-events-auto [&::-moz-range-thumb]:appearance-none [&::-moz-range-thumb]:w-4 [&::-moz-range-thumb]:h-4 [&::-moz-range-thumb]:bg-[#000000] [&::-moz-range-thumb]:rounded-full [&::-moz-range-thumb]:cursor-pointer [&::-moz-range-thumb]:border-0 [&::-moz-range-thumb]:shadow"
                            style={{
                                zIndex: minValue > rangeMax - 100 ? 5 : 3,
                            }}
                        />

                        {/* Max range input */}
                        <input
                            type="range"
                            min={rangeMin}
                            max={rangeMax}
                            value={maxValue}
                            onChange={(e) =>
                                handleMaxChange(parseFloat(e.target.value))
                            }
                            onMouseUp={handleMaxRelease}
                            onTouchEnd={handleMaxRelease}
                            disabled={!isEnabled}
                            className="absolute w-full h-6 appearance-none bg-transparent pointer-events-none [&::-webkit-slider-thumb]:pointer-events-auto [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:w-4 [&::-webkit-slider-thumb]:h-4 [&::-webkit-slider-thumb]:bg-[#000000] [&::-webkit-slider-thumb]:rounded-full [&::-webkit-slider-thumb]:cursor-pointer [&::-webkit-slider-thumb]:shadow [&::-moz-range-thumb]:pointer-events-auto [&::-moz-range-thumb]:appearance-none [&::-moz-range-thumb]:w-4 [&::-moz-range-thumb]:h-4 [&::-moz-range-thumb]:bg-[#000000] [&::-moz-range-thumb]:rounded-full [&::-moz-range-thumb]:cursor-pointer [&::-moz-range-thumb]:border-0 [&::-moz-range-thumb]:shadow"
                            style={{ zIndex: 4 }}
                        />
                    </div>
                </div>
            )}
        </div>
    );
}

export default RangeSlider;
