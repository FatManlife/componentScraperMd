//Product
export type Product = {
    id: number;
    name: string;
    image_url: string;
    brand: string;
    price: number;
    url: string;
    category: string;
    website: string;
    website_image: string;
};

export type Products = {
    id: number;
    name: string;
    image_url: string;
    brand: string;
    price: number;
    category: string;
    website: string;
    website_image: string;
};

export type ProductResponse = {
    count: number;
    products: Product[];
};

export type ProductOrder = "products.id ASC" | "price_asc" | "price_desc";

export type ProductParams = {
    name?: string;
    website?: string[];
    page?: number;
    min?: number;
    max?: number;
    order?: ProductOrder;
};

//Component Responses
export type AioResponse = {
    Product: Product;
    diagonal: number;
    cpu: string;
    ram: number;
    storage: number;
    gpu: string;
};

export type CaseResponse = {
    Product: Product;
    format: string;
    motherboard_form_factor: string;
};

export type CoolerResponse = {
    Product: Product;
    type: string;
    fan_rpm: number;
    noise: number;
    size: string;
    compatibility: string[];
};

export type CpuResponse = {
    Product: Product;
    cores: number;
    threads: number;
    base_clock: number;
    boost_clock: number;
    socket: string;
    tdp: number;
};

export type FanResponse = {
    Product: Product;
    fan_rpm: number;
    noise: number;
    size: string;
};

export type GpuResponse = {
    Product: Product;
    chipset: string;
    vram: number;
    gpu_frequency: number;
    vram_frequency: number;
};

export type HddResponse = {
    Product: Product;
    capacity: number;
    rotation_speed: number;
    form_factor: string;
};

export type LaptopResponse = {
    Product: Product;
    cpu: string;
    gpu: string;
    ram: number;
    storage: number;
    diagonal: number;
    battery: number;
};

export type MotherboardResponse = {
    Product: Product;
    chipset: string;
    socket: string;
    form_factor: string;
    ram_support: string;
    form_factor_ram: string;
};

export type PcResponse = {
    Product: Product;
    cpu: string;
    gpu: string;
    ram: number;
    storage: number;
    motherboard: string;
    psu: string;
    case: string;
};

export type PcMiniResponse = {
    Product: Product;
    cpu: string;
    gpu: string;
    ram: number;
    storage: number;
};

export type PsuResponse = {
    Product: Product;
    power: number;
    efficiency: string;
    form_factor: string;
};

//Params and specs for components
export type CpuParams = {
    defaultParams: ProductParams;
    cores?: number[];
    threads?: number[];
    base_clock?: number[];
    boost_clock?: number[];
    socket?: string[];
};

export type CpuSpecs = {
    Cores: number[];
    Threads: number[];
    BaseClock: number[];
    BoostClock: number[];
    Socket: string[];
};

export type FanParams = {
    defaultParams: ProductParams;
    min_fan_rpm?: number;
    max_fan_rpm?: number;
    min_noise?: number;
    max_noise?: number;
};

export type FanSpecs = {
    FanRPM: number[];
    Noise: number[];
};

export type GpuParams = {
    defaultParams: ProductParams;
    chipset?: string[];
    min_vram?: number;
    max_vram?: number;
    min_gpu_frequency?: number;
    max_gpu_frequency?: number;
    min_vram_frequency?: number;
    max_vram_frequency?: number;
};

export type GpuSpecs = {
    chipset: string[];
    vram: number[];
    gpu_frequency: number[];
    vram_frequency: number[];
};

export type HddParams = {
    defaultParams: ProductParams;
    min_capacity?: number;
    max_capacity?: number;
    min_rotation_speed?: number;
    max_rotation_speed?: number;
    form_factor?: string[];
};

export type HddSpecs = {
    capacity: number[];
    rotation_speed: number[];
    form_factor: string[];
};

export type LaptopParams = {
    defaultParams: ProductParams;
    cpu?: string[];
    gpu?: string[];
    ram?: number[];
    storage?: number[];
    diagonal?: number[];
};

export type LaptopSpecs = {
    cpu: string[];
    gpu: string[];
    ram: number[];
    storage: number[];
    diagonal: number[];
};

export type MotherboardParams = {
    defaultParams: ProductParams;
    chipset?: string[];
    socket?: string[];
    form_factor?: string[];
};

export type MotherboardSpecs = {
    Chipset: string[];
    Socket: string[];
    FormFactor: string[];
};

export type PcParams = {
    defaultParams: ProductParams;
    cpu?: string[];
    gpu?: string[];
    ram?: number[];
    storage?: number[];
};

export type PcSpecs = {
    cpu: string[];
    gpu: string[];
    ram: number[];
    storage: number[];
};

export type PcMiniParams = {
    defaultParams: ProductParams;
    cpu?: string[];
    gpu?: string[];
    ram?: number[];
    storage?: number[];
};

export type PcMiniSpecs = {
    cpu: string[];
    gpu: string[];
    ram: number[];
    storage: number[];
};

export type PsuParams = {
    defaultParams: ProductParams;
    min_power?: number;
    max_power?: number;
    efficiency?: string[];
    form_factor?: string[];
};

export type PsuSpecs = {
    power: number[];
    efficiency: string[];
    form_factor: string[];
};

export type RamResponse = {
    Product: Product;
    capacity: number;
    speed: number;
    type: string;
    compatibility: string;
    configuration: number;
};

export type RamParams = {
    defaultParams: ProductParams;
    min_capacity?: number;
    max_capacity?: number;
    min_speed?: number;
    max_speed?: number;
    type?: string[];
    compatibility?: string[];
    configuration?: number[];
};

export type RamSpecs = {
    capacity: number[];
    speed: number[];
    type: string[];
    compatibility: string[];
    configuration: number[];
};

export type SsdResponse = {
    Product: Product;
    capacity: number;
    reading_speed: number;
    writing_speed: number;
    form_factor: string;
};

export type SsdParams = {
    defaultParams: ProductParams;
    min_capacity?: number;
    max_capacity?: number;
    min_reading_speed?: number;
    max_reading_speed?: number;
    min_writing_speed?: number;
    max_writing_speed?: number;
    form_factor?: string[];
};

export type SsdSpecs = {
    capacity: number[];
    reading_speed: number[];
    writing_speed: number[];
    form_factor: string[];
};

export type CoolerParams = {
    defaultParams: ProductParams;
    type?: string[];
    fan_rpm?: number[];
    noise?: number[];
    compatibility?: string[];
};

export type CoolerSpecs = {
    Type: string[];
    FanRPM: number[];
    Noise: number[];
    Compatibility: string[];
};

export type CaseParams = {
    defaultParams: ProductParams;
    format?: string[];
    motherboard_form_factor?: string[];
};

export type CaseSpecs = {
    Format: string[];
    MotherboardFormFactor: string[];
};

export type AioParams = {
    defaultParams: ProductParams;
    diagonal?: number[];
    cpu?: string[];
    ram?: number[];
    storage?: number[];
    gpu?: string[];
};

export type AioSpecs = {
    Diagonal: number[];
    Cpu: string[];
    Ram: number[];
    Storage: number[];
    Gpu: string[];
};

export type DefaultSpecs = {
    websites: string[];
    prices: number[];
    order: string[];
};

export type ComponentFiltersResponse<T> = {
    defaultSpecs: DefaultSpecs;
    specificSpecs: T;
};
