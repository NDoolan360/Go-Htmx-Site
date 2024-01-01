export const replaceWithCurrentYear = (input: string, match: string): string => {
    return input.replace(match, new Date().getFullYear().toString());
};

// Fetch data from sites profile
export const fetchData = async <T>(
    source: string,
    parserType: DOMParserSupportedType | "json" = "text/html",
): Promise<T> => {
    let data: string;

    if (
        source.startsWith("/proxy") ||
        source.startsWith("http://") ||
        source.startsWith("https://")
    ) {
        const response = await fetch(source);
        data = await (response as Response).text();
    } else {
        // Read source as data if not a url (used for testing)
        data = source;
    }

    if (parserType === "json") {
        return JSON.parse(data);
    }
    const parser = new DOMParser();
    return parser.parseFromString(data, parserType) as T;
};
