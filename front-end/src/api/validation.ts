const MAX_STRING_LENGTH = 500;
const MAX_ARRAY_LENGTH = 100;
const MAX_NUMBER_VALUE = 1000000;

export const sanitizeString = (input: string): string => {
    if (typeof input !== "string") return "";
    return input.slice(0, MAX_STRING_LENGTH).trim();
};

export const sanitizeNumber = (input: number): number => {
    if (typeof input !== "number" || isNaN(input)) return 0;
    return Math.min(Math.abs(input), MAX_NUMBER_VALUE);
};

export const sanitizeArray = <T>(
    input: T[],
    maxLength: number = MAX_ARRAY_LENGTH,
): T[] => {
    if (!Array.isArray(input)) return [];
    return input.slice(0, maxLength);
};

export const sanitizeStringArray = (input: string[]): string[] => {
    return sanitizeArray(input).map(sanitizeString).filter(Boolean);
};

export const sanitizeNumberArray = (input: number[]): number[] => {
    return sanitizeArray(input)
        .map(sanitizeNumber)
        .filter((n) => !isNaN(n));
};
