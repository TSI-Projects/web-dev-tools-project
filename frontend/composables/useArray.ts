export default function () {
    const wrap = <T extends unknown>(value: any): T[] => {
        if (!value) {
            return [];
        }
    
        return !Array.isArray(value)
            ? [value]
            : value;
    };

    return {
        wrap,
    };
}
