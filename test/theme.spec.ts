import { describe, expect, test } from "bun:test";
import { onClick, setPreference } from "theme";

describe("Theme", () => {
    test("Toggle Theme (light -> dark)", () => {
        setPreference("light");
        onClick();

        expect(localStorage.getItem("theme")).toEqual("dark");
        expect(document.firstElementChild?.getAttribute("data-theme")).toEqual("dark");
    });
    test("Toggle Theme (dark -> light)", () => {
        setPreference("dark");
        onClick();

        expect(localStorage.getItem("theme")).toEqual("light");
        expect(document.firstElementChild?.getAttribute("data-theme")).toEqual("light");
    });
});
