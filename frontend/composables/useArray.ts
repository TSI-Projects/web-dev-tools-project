export default function () {
    const wrap = <T extends unknown, D extends unknown>(value: any, defaultValue: D): T[] | D => {
        if (!value) {
            return defaultValue;
        }
    
        return !Array.isArray(value)
            ? [value]
            : value;
    };

    return {
        wrap,
    };
}
