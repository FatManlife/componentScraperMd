import type {
    AioParams,
    CaseParams,
    CoolerParams,
    CpuParams,
    FanParams,
    GpuParams,
    HddParams,
    LaptopParams,
    MotherboardParams,
    PcParams,
    PcMiniParams,
    PsuParams,
    RamParams,
    SsdParams,
    ProductResponse,
} from "../constants/types";
import api from "./axios";
import { appendDefaultParams } from "./utils";

export const FetchAio = async (params: AioParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle AIO-specific params
        if (params.diagonal && params.diagonal.length > 0) {
            params.diagonal.forEach((d) =>
                queryParams.append("diagonal", d.toString()),
            );
        }
        if (params.cpu && params.cpu.length > 0) {
            params.cpu.forEach((c) => queryParams.append("cpu", c));
        }
        if (params.ram && params.ram.length > 0) {
            params.ram.forEach((r) => queryParams.append("ram", r.toString()));
        }
        if (params.storage && params.storage.length > 0) {
            params.storage.forEach((s) =>
                queryParams.append("storage", s.toString()),
            );
        }
        if (params.gpu && params.gpu.length > 0) {
            params.gpu.forEach((g) => queryParams.append("gpu", g));
        }

        const response = await api.get<ProductResponse>(
            `/aio?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchCase = async (
    params: CaseParams,
): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle Case-specific params
        if (params.format && params.format.length > 0) {
            params.format.forEach((f) => queryParams.append("format", f));
        }
        if (
            params.motherboard_form_factor &&
            params.motherboard_form_factor.length > 0
        ) {
            params.motherboard_form_factor.forEach((m) =>
                queryParams.append("motherboard_form_factor", m),
            );
        }

        const response = await api.get<ProductResponse>(
            `/case?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchCooler = async (
    params: CoolerParams,
): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle Cooler-specific params
        if (params.type && params.type.length > 0) {
            params.type.forEach((t) => queryParams.append("type", t));
        }
        if (params.fan_rpm && params.fan_rpm.length > 0) {
            params.fan_rpm.forEach((r) =>
                queryParams.append("fan_rpm", r.toString()),
            );
        }
        if (params.noise && params.noise.length > 0) {
            params.noise.forEach((n) =>
                queryParams.append("noise", n.toString()),
            );
        }
        if (params.compatibility && params.compatibility.length > 0) {
            params.compatibility.forEach((c) =>
                queryParams.append("compatibility", c),
            );
        }

        const response = await api.get<ProductResponse>(
            `/cooler?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchCpu = async (params: CpuParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle CPU-specific params
        if (params.cores && params.cores.length > 0) {
            params.cores.forEach((c) =>
                queryParams.append("cores", c.toString()),
            );
        }
        if (params.threads && params.threads.length > 0) {
            params.threads.forEach((t) =>
                queryParams.append("threads", t.toString()),
            );
        }
        if (params.base_clock && params.base_clock.length > 0) {
            params.base_clock.forEach((b) =>
                queryParams.append("base_clock", b.toString()),
            );
        }
        if (params.boost_clock && params.boost_clock.length > 0) {
            params.boost_clock.forEach((b) =>
                queryParams.append("boost_clock", b.toString()),
            );
        }
        if (params.socket && params.socket.length > 0) {
            params.socket.forEach((s) => queryParams.append("socket", s));
        }

        const response = await api.get<ProductResponse>(
            `/cpu?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchFan = async (params: FanParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle Fan-specific params
        if (params.min_fan_rpm !== undefined) {
            queryParams.append("min_fan_rpm", params.min_fan_rpm.toString());
        }
        if (params.max_fan_rpm !== undefined) {
            queryParams.append("max_fan_rpm", params.max_fan_rpm.toString());
        }
        if (params.min_noise !== undefined) {
            queryParams.append("min_noise", params.min_noise.toString());
        }
        if (params.max_noise !== undefined) {
            queryParams.append("max_noise", params.max_noise.toString());
        }

        const response = await api.get<ProductResponse>(
            `/fan?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchGpu = async (params: GpuParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle GPU-specific params
        if (params.chipset && params.chipset.length > 0) {
            params.chipset.forEach((c) => queryParams.append("chipset", c));
        }
        if (params.min_vram !== undefined) {
            queryParams.append("min_vram", params.min_vram.toString());
        }
        if (params.max_vram !== undefined) {
            queryParams.append("max_vram", params.max_vram.toString());
        }
        if (params.min_gpu_frequency !== undefined) {
            queryParams.append(
                "min_gpu_frequency",
                params.min_gpu_frequency.toString(),
            );
        }
        if (params.max_gpu_frequency !== undefined) {
            queryParams.append(
                "max_gpu_frequency",
                params.max_gpu_frequency.toString(),
            );
        }
        if (params.min_vram_frequency !== undefined) {
            queryParams.append(
                "min_vram_frequency",
                params.min_vram_frequency.toString(),
            );
        }
        if (params.max_vram_frequency !== undefined) {
            queryParams.append(
                "max_vram_frequency",
                params.max_vram_frequency.toString(),
            );
        }

        const response = await api.get<ProductResponse>(
            `/gpu?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchHdd = async (params: HddParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle HDD-specific params
        if (params.min_capacity !== undefined) {
            queryParams.append("min_capacity", params.min_capacity.toString());
        }
        if (params.max_capacity !== undefined) {
            queryParams.append("max_capacity", params.max_capacity.toString());
        }
        if (params.min_rotation_speed !== undefined) {
            queryParams.append(
                "min_rotation_speed",
                params.min_rotation_speed.toString(),
            );
        }
        if (params.max_rotation_speed !== undefined) {
            queryParams.append(
                "max_rotation_speed",
                params.max_rotation_speed.toString(),
            );
        }
        if (params.form_factor && params.form_factor.length > 0) {
            params.form_factor.forEach((f) =>
                queryParams.append("form_factor", f),
            );
        }

        const response = await api.get<ProductResponse>(
            `/hdd?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchLaptop = async (
    params: LaptopParams,
): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle Laptop-specific params
        if (params.cpu && params.cpu.length > 0) {
            params.cpu.forEach((c) => queryParams.append("cpu", c));
        }
        if (params.gpu && params.gpu.length > 0) {
            params.gpu.forEach((g) => queryParams.append("gpu", g));
        }
        if (params.ram && params.ram.length > 0) {
            params.ram.forEach((r) => queryParams.append("ram", r.toString()));
        }
        if (params.storage && params.storage.length > 0) {
            params.storage.forEach((s) =>
                queryParams.append("storage", s.toString()),
            );
        }
        if (params.diagonal && params.diagonal.length > 0) {
            params.diagonal.forEach((d) =>
                queryParams.append("diagonal", d.toString()),
            );
        }

        const response = await api.get<ProductResponse>(
            `/laptop?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchMotherboard = async (
    params: MotherboardParams,
): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle Motherboard-specific params
        if (params.chipset && params.chipset.length > 0) {
            params.chipset.forEach((c) => queryParams.append("chipset", c));
        }
        if (params.socket && params.socket.length > 0) {
            params.socket.forEach((s) => queryParams.append("socket", s));
        }
        if (params.form_factor && params.form_factor.length > 0) {
            params.form_factor.forEach((f) =>
                queryParams.append("form_factor", f),
            );
        }

        const response = await api.get<ProductResponse>(
            `/motherboard?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchPc = async (params: PcParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle PC-specific params
        if (params.cpu && params.cpu.length > 0) {
            params.cpu.forEach((c) => queryParams.append("cpu", c));
        }
        if (params.gpu && params.gpu.length > 0) {
            params.gpu.forEach((g) => queryParams.append("gpu", g));
        }
        if (params.ram && params.ram.length > 0) {
            params.ram.forEach((r) => queryParams.append("ram", r.toString()));
        }
        if (params.storage && params.storage.length > 0) {
            params.storage.forEach((s) =>
                queryParams.append("storage", s.toString()),
            );
        }

        const response = await api.get<ProductResponse>(
            `/pc?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchPcMini = async (
    params: PcMiniParams,
): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle PC Mini-specific params
        if (params.cpu && params.cpu.length > 0) {
            params.cpu.forEach((c) => queryParams.append("cpu", c));
        }
        if (params.gpu && params.gpu.length > 0) {
            params.gpu.forEach((g) => queryParams.append("gpu", g));
        }
        if (params.ram && params.ram.length > 0) {
            params.ram.forEach((r) => queryParams.append("ram", r.toString()));
        }
        if (params.storage && params.storage.length > 0) {
            params.storage.forEach((s) =>
                queryParams.append("storage", s.toString()),
            );
        }

        const response = await api.get<ProductResponse>(
            `/pcmini?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchPsu = async (params: PsuParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle PSU-specific params
        if (params.min_power !== undefined) {
            queryParams.append("min_power", params.min_power.toString());
        }
        if (params.max_power !== undefined) {
            queryParams.append("max_power", params.max_power.toString());
        }
        if (params.efficiency && params.efficiency.length > 0) {
            params.efficiency.forEach((e) =>
                queryParams.append("efficiency", e),
            );
        }
        if (params.form_factor && params.form_factor.length > 0) {
            params.form_factor.forEach((f) =>
                queryParams.append("form_factor", f),
            );
        }

        const response = await api.get<ProductResponse>(
            `/psu?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchRam = async (params: RamParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle RAM-specific params
        if (params.min_capacity !== undefined) {
            queryParams.append("min_capacity", params.min_capacity.toString());
        }
        if (params.max_capacity !== undefined) {
            queryParams.append("max_capacity", params.max_capacity.toString());
        }
        if (params.min_speed !== undefined) {
            queryParams.append("min_speed", params.min_speed.toString());
        }
        if (params.max_speed !== undefined) {
            queryParams.append("max_speed", params.max_speed.toString());
        }
        if (params.type && params.type.length > 0) {
            params.type.forEach((t) => queryParams.append("type", t));
        }
        if (params.compatibility && params.compatibility.length > 0) {
            params.compatibility.forEach((c) =>
                queryParams.append("compatibility", c),
            );
        }
        if (params.configuration && params.configuration.length > 0) {
            params.configuration.forEach((c) =>
                queryParams.append("configuration", c.toString()),
            );
        }

        const response = await api.get<ProductResponse>(
            `/ram?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const FetchSsd = async (params: SsdParams): Promise<ProductResponse> => {
    try {
        const queryParams = new URLSearchParams();

        // Handle default ProductParams
        appendDefaultParams(queryParams, params.defaultParams);

        // Handle SSD-specific params
        if (params.min_capacity !== undefined) {
            queryParams.append("min_capacity", params.min_capacity.toString());
        }
        if (params.max_capacity !== undefined) {
            queryParams.append("max_capacity", params.max_capacity.toString());
        }
        if (params.min_reading_speed !== undefined) {
            queryParams.append(
                "min_reading_speed",
                params.min_reading_speed.toString(),
            );
        }
        if (params.max_reading_speed !== undefined) {
            queryParams.append(
                "max_reading_speed",
                params.max_reading_speed.toString(),
            );
        }
        if (params.min_writing_speed !== undefined) {
            queryParams.append(
                "min_writing_speed",
                params.min_writing_speed.toString(),
            );
        }
        if (params.max_writing_speed !== undefined) {
            queryParams.append(
                "max_writing_speed",
                params.max_writing_speed.toString(),
            );
        }
        if (params.form_factor && params.form_factor.length > 0) {
            params.form_factor.forEach((f) =>
                queryParams.append("form_factor", f),
            );
        }

        const response = await api.get<ProductResponse>(
            `/ssd?${queryParams.toString()}`,
        );
        return response.data;
    } catch (err) {
        console.error(err);
        throw err;
    }
};
