export const replaceWithCurrentYear = (input: string, match: string): string => {
    return input.replace(match, new Date().getFullYear().toString());
};

export const replaceCopyright = () => {
    const copyright = document.getElementById('copyright');
    if (copyright) {
        copyright.innerHTML = replaceWithCurrentYear(copyright.innerHTML, '{current year}');
    }
};
