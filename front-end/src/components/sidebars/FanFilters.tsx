import RangeSlider from "../RangeSlider";
import type { FanSpecs } from "../../constants/types";

type FanFiltersProps = {
    specs: FanSpecs;
};

function FanFilters({ specs }: FanFiltersProps) {
    return (
        <>
            {/* Fan RPM Section */}
            {specs.FanRPM && specs.FanRPM.length > 0 && (
                <RangeSlider
                    values={specs.FanRPM}
                    label="Fan RPM"
                    text="Fan RPM"
                    minParamKey="min_fan_rpm"
                    maxParamKey="max_fan_rpm"
                    formatValue={(value) => `${value.toFixed(0)} RPM`}
                />
            )}

            {/* Noise Section */}
            {specs.Noise && specs.Noise.length > 0 && (
                <RangeSlider
                    values={specs.Noise}
                    label="Noise Level"
                    text="Noise"
                    minParamKey="min_noise"
                    maxParamKey="max_noise"
                    formatValue={(value) => `${value.toFixed(1)} dB`}
                />
            )}
        </>
    );
}

export default FanFilters;
