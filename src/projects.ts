import domPurify from "dompurify";
import * as githubColorsJson from "./github-colors.json";
import { fetchData } from "./utils";

type Site = "Cults 3D" | "Github" | "Board Game Geek";
type Language = { name: string; color: string };
type Image = {
    highResSrc: string | null;
    lowResSrc: string | null;
    alt: string | null;
};

type Project = {
    host?: Site;
    title?: string;
    description?: string;
    url?: URL;
    image?: Image;
    programmingLanguage?: Language;
    chips?: string[];
};

export type GithubRepo = {
    name: string;
    description: string;
    fork: boolean;
    language: string;
    topics: string[];
    // biome-ignore lint/style/useNamingConvention: External API naming convention
    html_url: string;
};

export const githubRepoToProject = (data: GithubRepo[]): Project[] => {
    if (!Array.isArray(data)) {
        return [];
    }
    const githubColors: { [k: string]: { color: string } | undefined } = githubColorsJson;
    return data
        .filter((r) => !r.fork && r.topics.length > 0)
        .map((r): Project => {
            return {
                host: "Github",
                description: r.description,
                title: r.name.replace(/[\_\-\.]/g, " "),
                programmingLanguage: {
                    name: r.language,
                    color: githubColors[r.language]?.color ?? "",
                },
                url: new URL(r.html_url),
                chips: r.topics,
            };
        });
};

export const scrapeCults3d = (doc: Document): Project[] => {
    const cults3dProjects: Project[] = [];

    const projectElements = doc.querySelectorAll('article[class*="crea"]');

    for (const projectElement of projectElements) {
        const titleElement = projectElement.querySelector('a[class*="drawer-contents"]');
        const urlElement = projectElement.querySelector('a[class*="drawer-contents"]');
        const imageElement = projectElement.querySelector('img[class*="painting-image"]');

        let title: string | undefined;
        let url: URL | undefined;
        let image: Image | undefined;

        const titleValue = titleElement?.getAttribute("title");
        if (titleValue) {
            title = titleValue.trim();
        }
        const href = urlElement?.getAttribute("href");
        if (href) {
            url = new URL(href, "https://cults3d.com");
        }
        const dataSrc = imageElement?.getAttribute("data-src");
        if (dataSrc) {
            let source = dataSrc;
            let sourceBackup = null;

            // extract full size file rather than thumbnail image if possible
            const regex = /https:\/\/files\.cults3d\.com[^'"]+/;
            const match = source?.match(regex);

            if (match?.[0]) {
                sourceBackup = source;
                source = match[0];
            }

            image = {
                highResSrc: source,
                lowResSrc: sourceBackup,
                alt: imageElement?.getAttribute("alt"),
            } as Image;
        }

        cults3dProjects.push({ host: "Cults 3D", title, url, image });
    }

    return cults3dProjects;
};

export const scrapeBgg = (doc: Document): Project[] => {
    const bggProjects: Project[] = [];
    const projectElements = doc.querySelectorAll('tr[id*="row_"]');

    for (const projectElement of projectElements) {
        const titleElement = projectElement.querySelector(
            'td[class*="collection_objectname"] > div > a',
        );
        const descriptionElement = projectElement.querySelector(
            'td[class*="collection_objectname"] > p',
        );
        const imageElement = projectElement.querySelector(
            'td[class*="collection_thumbnail"] > a > img',
        );
        const urlElement = projectElement.querySelector('td[class*="collection_thumbnail"] > a');

        let title: string | undefined;
        let description: string | undefined;
        let url: URL | undefined;
        let image: Image | undefined;

        if (titleElement) {
            title = titleElement.textContent?.trim();
        }
        if (descriptionElement) {
            description = descriptionElement.textContent?.trim();
        }
        const href = urlElement?.getAttribute("href");
        if (href) {
            url = new URL(href, "https://boardgamegeek.com");
        }
        if (imageElement?.getAttribute("src")) {
            image = {
                highResSrc: imageElement.getAttribute("src"),
                lowResSrc: null,
                alt: imageElement.getAttribute("alt"),
            };
        }

        bggProjects.push({
            host: "Board Game Geek",
            title,
            description,
            url,
            image,
        });
    }

    return bggProjects;
};

export const upgradeBggData = (project: Project, xmlDoc: XMLDocument) => {
    const imageXmlElement = xmlDoc.getElementsByTagName("image").item(0);
    if (imageXmlElement && project.image) {
        project.image.lowResSrc = project.image.highResSrc;
        project.image.highResSrc = imageXmlElement.textContent;
    }
    const mechanicXmlElements = xmlDoc.getElementsByTagName("boardgamemechanic");
    const mechanics: string[] = [];
    for (const mechanic of mechanicXmlElements) {
        mechanics.push(mechanic.textContent ?? "");
    }

    project.chips = mechanics;
};

export const projectIntoTemplate = (
    project: Project,
    template: HTMLTemplateElement,
): DocumentFragment => {
    const templateClone = document.importNode(template.content, true);

    const setElementContent = <T extends Element, U>(
        selector: string,
        content: U | undefined,
        setter: (element: T, content: U) => void,
    ) => {
        const element = templateClone.querySelector<T>(selector);
        if (element) {
            if (content) {
                setter(element, content);
            } else {
                element.remove();
            }
        }
    };

    // Set project title
    setElementContent('[class*="card-title"]', project.title, (element, content) => {
        element.textContent = domPurify.sanitize(content);
    });

    // Set project URL
    setElementContent<HTMLLinkElement, URL>(
        '[class*="card-link"]',
        project.url,
        (element, content) => {
            element.href = domPurify.sanitize(content.href);
        },
    );

    // Set project link aria-label
    setElementContent<HTMLLinkElement, (string | undefined)[]>(
        '[class*="card-link"]',
        [project.host, project.title],
        (element, [host, title]) => {
            if (host && title) {
                element.ariaLabel = `${domPurify.sanitize(title)} on ${domPurify.sanitize(host)}`;
            }
        },
    );

    // Set project description
    setElementContent('[class*="card-description"]', project.description, (element, content) => {
        element.textContent = domPurify.sanitize(content);
    });

    // Set project language
    setElementContent<Element, Language>(
        '[class*="card-language-name"]',
        project.programmingLanguage,
        (element, content) => {
            element.textContent = domPurify.sanitize(content.name);
        },
    );

    // Set project language colour
    setElementContent<Element, Language>(
        '[class*="card-language-colour"]',
        project.programmingLanguage,
        (element, content) => {
            element.setAttribute("style", `background-color: ${domPurify.sanitize(content.color)}`);
        },
    );

    // Set project language colour
    setElementContent<Element, string[]>(
        '[class*="card-chips"]',
        project.chips,
        (element, content) => {
            for (const chip of content) {
                const chipElement = document.createElement("span");
                chipElement.classList.add("chip");
                chipElement.textContent = chip;
                element.append(chipElement);
            }
        },
    );

    // Set logo image and aria-label
    setElementContent<SVGElement, string>(
        '[class*="card-logo"]',
        project.host,
        (element, content) => {
            element.ariaLabel = `${domPurify.sanitize(content)} Logo`;
            const use = element.children.item(0) as SVGUseElement;
            use.setAttribute(
                "href",
                `/images/logos/${domPurify
                    .sanitize(content)
                    .toLowerCase()
                    .replace(/\s/g, "")}.svg#logo`,
            );
        },
    );

    // Set project feature image
    setElementContent<HTMLImageElement, Image>(
        '[class*="card-feature-image"]',
        project.image,
        (element, content) => {
            element.alt = content.alt ?? "Feature image";

            // Chain loading of progressively higher res images (default -> srcBackup -> src)
            element.src = "/images/default.webp";
            const lowRes = content.lowResSrc;
            const highRes = content.highResSrc;
            if (lowRes) {
                element.onload = () => {
                    element.src = domPurify.sanitize(lowRes);
                    if (highRes) {
                        element.loading = "lazy";
                        element.onload = () => {
                            element.src = domPurify.sanitize(highRes);
                            element.onload = () => {
                                // Prevents infinite loading
                            };
                        };
                    }
                };
            } else if (highRes) {
                element.loading = "lazy";
                element.onload = () => {
                    element.src = domPurify.sanitize(highRes);
                    element.onload = () => {
                        // Prevents infinite loading
                    };
                };
            }
        },
    );

    return templateClone;
};

export const loadProjects = async () => {
    // Create project-gallery loader
    const loader = document.createElement("span");
    loader.classList.add("loader");
    document.getElementById("projects")?.append(loader);

    // Load items into gallery
    const gallery = document.getElementById("project-gallery");
    const template = document.getElementById("project-template") as HTMLTemplateElement | undefined;

    if (gallery && template) {
        const githubPage = await fetchData<GithubRepo[]>("/proxy/api/github", "json");
        const githubProjects = githubRepoToProject(githubPage).map((p) =>
            projectIntoTemplate(p, template),
        );
        gallery.append(...githubProjects);

        const bggPage = await fetchData<Document>("/proxy/boardgamegeek");
        const bggRawProjects = scrapeBgg(bggPage);

        // Get higher resolution image from bgg xmlapi
        for (const project of bggRawProjects) {
            const id = project.url
                ?.toString()
                ?.split("/")
                .find((v) => v.match(/\d+/g));
            const gameXml = await fetchData<Document>(
                `/proxy/xmlapi/boardgamegeek/${id}`,
                "text/xml",
            );
            upgradeBggData(project, gameXml);
        }

        const bggProjects = bggRawProjects.map((p) => projectIntoTemplate(p, template));

        gallery.append(...bggProjects);

        const cults3dPage = await fetchData<Document>("/proxy/cults3d");
        const cults3dProjects = scrapeCults3d(cults3dPage).map((p) =>
            projectIntoTemplate(p, template),
        );
        gallery.append(...cults3dProjects);
    }

    // remove project-gallery loader
    loader.remove();
};
