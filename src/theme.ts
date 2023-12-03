import "./theme-switch.css";

// Below script Adapted from:
// https://web.dev/patterns/theming/theme-switch/#js
let theme = "dark";

const storageTheme = localStorage.getItem("theme");
if (storageTheme) {
    theme = storageTheme;
} else if (window.matchMedia) {
    theme = window.matchMedia("(prefers-color-scheme: light)").matches ? "light" : "dark";
}
document.documentElement.setAttribute("data-theme", theme);

export const onClick = () => {
    // flip current value
    theme = theme === "light" ? "dark" : "light";

    setPreference(theme);
};

export const setPreference = (theme: string) => {
    localStorage.setItem("theme", theme);
    reflectPreference(theme);
};

const reflectPreference = (theme: string) => {
    document.documentElement.setAttribute("data-theme", theme);
    document.querySelector("#theme-toggle")?.setAttribute("aria-label", theme);
};

// set early so no page flashes / CSS is made aware
reflectPreference(theme);

window.onload = () => {
    // set on load so screen readers can see latest value on the button
    reflectPreference(theme);

    const themeTemplate = document.getElementById("theme-toggle-template") as
        | HTMLTemplateElement
        | undefined;
    if (themeTemplate) {
        const themeTemplateClone = document.importNode(themeTemplate.content, true);
        themeTemplate.parentElement?.appendChild(themeTemplateClone);
    }
    const themeToggle = document.getElementById("theme-toggle");
    if (themeToggle) {
        themeToggle.addEventListener("click", onClick);
    }
};

// sync with system changes
if (window.matchMedia) {
    window
        .matchMedia("(prefers-color-scheme: light)")
        .addEventListener("change", ({ matches: isLight }) => {
            theme = isLight ? "light" : "dark";
            setPreference(theme);
        });
}
