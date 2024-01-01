import { describe, expect, test } from "bun:test";
import { file } from "bun";
import {
    GithubRepo,
    githubRepoToProject,
    projectIntoTemplate,
    scrapeBgg,
    scrapeCults3d,
    upgradeBggData,
} from "projects";
import { fetchData } from "utils";

describe("Projects", () => {
    test("Github project into Template", async () => {
        const indexData = await file("index.html").text();
        const indexMockDoc = await fetchData<Document>(indexData);
        const githubMockData = await file("test/data/githubRepos.json").text();
        const githubMockJson = await fetchData<GithubRepo[]>(githubMockData, "json");

        const template = indexMockDoc.getElementById("project-template") as HTMLTemplateElement;
        const githubProjects = githubRepoToProject(githubMockJson);
        const project = githubProjects.at(0);

        expect(project).not.toBeUndefined();
        if (project) {
            expect(template).not.toBeUndefined();
            const fragment = projectIntoTemplate(project, template);

            expect(
                fragment.querySelector<HTMLHeadingElement>('[class*="card-title"]')?.textContent,
            ).toBe("NDoolan360 Site");
            expect(
                fragment.querySelector<HTMLParagraphElement>('[class*="card-description"]')
                    ?.textContent,
            ).toBe("My hand crafted personal website 🎨🌝");
            expect(fragment.querySelector<HTMLAnchorElement>('[class*="card-link"]')?.href).toBe(
                "https://github.com/NDoolan360/NDoolan360-Site",
            );
            expect(
                fragment.querySelector<HTMLParagraphElement>('[class*="card-language-name"]')
                    ?.textContent,
            ).toBe("TypeScript");
            expect(
                fragment
                    .querySelector<HTMLSpanElement>('[class*="card-language-colour"]')
                    ?.getAttribute("style"),
            ).toBe("background-color: #3178c6");
            expect(fragment.querySelector('[class*="card-logo"]')?.ariaLabel).toBe("Github Logo");
            expect(
                fragment.querySelector<HTMLImageElement>('[class*="card-feature-image"]')?.src,
            ).toBeUndefined();
        }
    });

    test("Cults3d project into Template", async () => {
        const indexData = await file("index.html").text();
        const indexMockDoc = await fetchData<Document>(indexData);
        const cults3dMockData = await file("test/data/cults3dProjects.html").text();
        const cults3dMockDoc = await fetchData<Document>(cults3dMockData);

        const template = indexMockDoc.getElementById("project-template") as HTMLTemplateElement;
        const cults3dProjects = scrapeCults3d(cults3dMockDoc);
        const project = cults3dProjects.at(0);

        expect(project).not.toBeUndefined();
        if (project) {
            expect(template).not.toBeUndefined();
            const fragment = projectIntoTemplate(project, template);

            expect(
                fragment.querySelector<HTMLHeadingElement>('[class*="card-title"]')?.textContent,
            ).toBe("Reciprocating Rack and Pinion Fidget V2");
            expect(
                fragment.querySelector<HTMLParagraphElement>('[class*="card-description"]')
                    ?.textContent,
            ).toBeUndefined();
            expect(fragment.querySelector<HTMLAnchorElement>('[class*="card-link"]')?.href).toBe(
                "https://cults3d.com/en/3d-model/gadget/reciprocating-rack-and-pinion-fidget-v2",
            );
            expect(
                fragment.querySelector<HTMLParagraphElement>('[class*="card-language-name"]')
                    ?.textContent,
            ).toBeUndefined();
            expect(
                fragment
                    .querySelector<HTMLSpanElement>('[class*="card-language-colour"]')
                    ?.getAttribute("style"),
            ).toBeUndefined();
            expect(fragment.querySelector('[class*="card-logo"]')?.ariaLabel).toBe("Cults 3D Logo");
            const featureImage = fragment.querySelector<HTMLImageElement>(
                '[class*="card-feature-image"]',
            );
            expect(featureImage).not.toBeUndefined();
            if (featureImage) {
                expect(featureImage.src).toBe("/images/default.webp");
                if (featureImage.onload) {
                    featureImage.onload(new Event("load"));
                    expect(featureImage.src).toBe(
                        "https://images.cults3d.com/{RRaP Image Link}/https://files.cults3d.com/{RRaP High-res Image Link}",
                    );

                    featureImage.onload(new Event("load"));
                    expect(featureImage.src).toBe(
                        "https://files.cults3d.com/{RRaP High-res Image Link}",
                    );
                }
            }
        }
    });

    test("Bgg project into Template", async () => {
        const indexData = await file("index.html").text();
        const indexMockDoc = await fetchData<Document>(indexData);
        const bggMockData = await file("test/data/bggProjects.html").text();
        const bggMockDoc = await fetchData<Document>(bggMockData);
        const bggMockXmlData = await file("test/data/bggImage.xml").text();
        const bggMockXml = await fetchData<Document>(bggMockXmlData, "text/xml");

        const template = indexMockDoc.getElementById("project-template") as HTMLTemplateElement;
        const bggProjects = scrapeBgg(bggMockDoc);
        const project = bggProjects.at(0);

        expect(project).not.toBeUndefined();
        if (project) {
            upgradeBggData(project, bggMockXml);

            expect(template).not.toBeUndefined();
            const fragment = projectIntoTemplate(project, template);

            expect(
                fragment.querySelector<HTMLHeadingElement>('[class*="card-title"]')?.textContent,
            ).toBe("Cake Toppers");
            expect(
                fragment.querySelector<HTMLParagraphElement>('[class*="card-description"]')
                    ?.textContent,
            ).toBe("Bakers assemble the most outrageous cakes to top each other.");
            expect(fragment.querySelector<HTMLAnchorElement>('[class*="card-link"]')?.href).toBe(
                "https://boardgamegeek.com/boardgame/330653/cake-toppers",
            );
            expect(
                fragment.querySelector<HTMLParagraphElement>('[class*="card-language-name"]')
                    ?.textContent,
            ).toBeUndefined();
            expect(
                fragment
                    .querySelector<HTMLSpanElement>('[class*="card-language-colour"]')
                    ?.getAttribute("style"),
            ).toBeUndefined();
            expect(fragment.querySelector('[class*="card-logo"]')?.ariaLabel).toBe(
                "Board Game Geek Logo",
            );
            const featureImage = fragment.querySelector<HTMLImageElement>(
                '[class*="card-feature-image"]',
            );
            expect(featureImage).not.toBeUndefined();
            if (featureImage) {
                expect(featureImage.src).toBe("/images/default.webp");
            }

            if (featureImage?.onload) {
                featureImage.onload(new Event("load"));
                expect(featureImage.src).toBe("{Cake Toppers Image Link}");

                featureImage.onload(new Event("load"));
                expect(featureImage.src).toBe("{Cake Toppers High-res Image Link}");
            }
        }
    });
});
