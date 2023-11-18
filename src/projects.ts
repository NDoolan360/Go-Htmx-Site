import DOMPurify from 'dompurify';

type Site = 'cults3d' | 'github' | 'boardgamegeek';
type Language = { name: string; style: string };
type Image = { src: string | null; srcBackup: string | null; alt: string | null };

type Project = {
    host?: Site;
    title?: string;
    description?: string;
    url?: URL;
    image?: Image;
    programmingLanguage?: Language;
};

// Fetch data from sites profile
export const fetchData = async (site: string, parserType: DOMParserSupportedType = 'text/html'): Promise<Document> => {
    const response = await fetch(site);
    const data = await response.text();

    const parser = new DOMParser();
    return parser.parseFromString(data, parserType);
};

export const scrapeGithub = (doc: Document): Project[] => {
    const githubProjects: Project[] = [];

    const projectElements = doc.querySelectorAll('div[class*="Box pinned-item-list-item"]:not(div[class*="fork"])');

    for (const projectElement of projectElements) {
        const titleElement = projectElement.querySelector('span[class*="repo"]');
        const descriptionElement = projectElement.querySelector('p[class*="pinned-item-desc"]');
        const urlElement = projectElement.querySelector('a[class*="Link"]');
        const langaugeNameElement = projectElement.querySelector('span[itemprop*="programmingLanguage"]');
        const langaugeColourElement = projectElement.querySelector('span[class*="repo-language-color"]');

        let title, description, url, programmingLanguage;
        if (titleElement) {
            title = titleElement.innerHTML.trim();
        }
        if (descriptionElement) {
            description = descriptionElement.innerHTML.trim();
        }
        if (urlElement) {
            url = new URL(urlElement.getAttribute('href')!, 'https://github.com');
        }
        if (langaugeNameElement?.innerHTML && langaugeColourElement?.getAttribute('style')) {
            programmingLanguage = {
                name: langaugeNameElement.innerHTML,
                style: langaugeColourElement.getAttribute('style')!,
            };
        }

        githubProjects.push({
            host: 'github',
            title,
            description,
            image: { src: '/images/github.png', srcBackup: null, alt: 'Github Logo' },
            url,
            programmingLanguage,
        });
    }

    return githubProjects;
};

export const scrapeCults3d = (doc: Document): Project[] => {
    const cults3dProjects: Project[] = [];

    const projectElements = doc.querySelectorAll('article[class*="crea"]');

    for (const projectElement of projectElements) {
        const titleElement = projectElement.querySelector('a[class*="drawer-contents"]');
        const urlElement = projectElement.querySelector('a[class*="drawer-contents"]');
        const imageElement = projectElement.querySelector('img[class*="painting-image"]');

        let title, url, image;
        if (titleElement?.getAttribute('title')) {
            title = titleElement.getAttribute('title')!.trim();
        }
        if (urlElement?.getAttribute('href')) {
            url = new URL(urlElement.getAttribute('href')!, 'https://cults3d.com');
        }
        if (imageElement?.getAttribute('data-src')) {
            let source = imageElement?.getAttribute('data-src');
            let sourceBackup = null;

            // extract full size file rather than thumbnail image if possible
            const regex = /https:\/\/files\.cults3d\.com[^'"]+/;
            const match = source?.match(regex);

            if (match && match[0]) {
                sourceBackup = source;
                source = match[0];
            }

            image = {
                src: source,
                srcBackup: sourceBackup,
                alt: imageElement.getAttribute('alt'),
            } as Image;
        }

        cults3dProjects.push({ host: 'cults3d', title, url, image });
    }

    return cults3dProjects;
};

export const scrapeBgg = (doc: Document): Project[] => {
    const bggProjects: Project[] = [];
    const projectElements = doc.querySelectorAll('tr[id*="row_"]');

    for (const projectElement of projectElements) {
        const titleElement = projectElement.querySelector('td[class*="collection_objectname"] > div > a');
        const descriptionElement = projectElement.querySelector('td[class*="collection_objectname"] > p');
        const imageElement = projectElement.querySelector('td[class*="collection_thumbnail"] > a > img');
        const urlElement = projectElement.querySelector('td[class*="collection_thumbnail"] > a');

        let title, description, url, image;
        if (titleElement) {
            title = titleElement.innerHTML.trim();
        }
        if (descriptionElement) {
            description = descriptionElement.innerHTML.trim();
        }
        if (urlElement?.getAttribute('href')) {
            url = new URL(urlElement.getAttribute('href')!, 'https://boardgamegeek.com');
        }
        if (imageElement?.getAttribute('src')) {
            image = {
                src: imageElement.getAttribute('src'),
                srcBackup: null,
                alt: imageElement.getAttribute('alt'),
            } as Image;
        }

        bggProjects.push({ host: 'boardgamegeek', title, description, url, image });
    }

    return bggProjects;
};

export const upgradeBggImage = (project: Project, xmlDoc: XMLDocument) => {
    const imageXmlElement = xmlDoc.getElementsByTagName('image').item(0);
    if (imageXmlElement && project.image) {
        project.image.srcBackup = project.image.src;
        project.image.src = imageXmlElement.innerHTML;
    }
};

export const projectIntoTemplate = (project: Project, template: HTMLTemplateElement): DocumentFragment => {
    const templateClone = document.importNode(template.content, true);

    // Set project feature image
    const imgElement = templateClone.querySelector<HTMLImageElement>('[class="card-feature-image"]')!;
    imgElement.alt = DOMPurify.sanitize(project.image?.alt ?? project.title ?? 'Feature image');
    // Chain loading of progressively higher res images (default -> srcBackup -> src)
    imgElement.src = '/images/default.png';
    if (project.image?.src) {
        const src = DOMPurify.sanitize(project.image!.src!);
        const backup = DOMPurify.sanitize(project.image?.srcBackup ?? src);
        // After loading the default, load the backup
        imgElement.onload = () => {
            imgElement.src = backup;
            // After loading the backup load the high-res
            imgElement.onload = () => {
                imgElement.src = src;
            };
        };
    } else {
        // Omit image if not present
        imgElement.remove();
    }

    // Set project title
    const titleElement = templateClone.querySelector('[class="card-heading"]')!;
    if (project.title) {
        titleElement.textContent = DOMPurify.sanitize(project.title);
    } else {
        // Omit title if not present
        titleElement.remove();
    }

    // Set project description
    const descriptionElement = templateClone.querySelector('[class="card-description"]')!;
    if (project.description) {
        descriptionElement.textContent = DOMPurify.sanitize(project.description);
    } else {
        // Omit description if not present
        descriptionElement.remove();
    }

    // Set project URL
    const linkElement = templateClone.querySelector<HTMLAnchorElement>('[class*="card-link"]')!;
    if (project.url) {
        linkElement.href = DOMPurify.sanitize(project.url.href);
    } else {
        // Omit link if not present
        linkElement.remove();
    }

    // Set project language
    const languageElement = templateClone.querySelector('[class="card-language-colour"]')!;
    const languageTextElement = templateClone.querySelector('[class="card-language"]')!;
    if (project.programmingLanguage && project.programmingLanguage.name) {
        languageElement.setAttribute('style', DOMPurify.sanitize(project.programmingLanguage.style));
        languageTextElement.textContent = DOMPurify.sanitize(project.programmingLanguage.name);
    } else {
        // Omit language if not present
        languageElement.remove();
        languageTextElement.remove();
    }

    // Set logo
    const logoElement = templateClone.querySelector<LogoLink>('[class*="card-logo"]')!;
    if (project.host && project.url) {
        const host = DOMPurify.sanitize(project.host);
        templateClone.firstElementChild?.classList.add(host);
        logoElement.setAttribute('href', project.url.toString());
        logoElement.innerHTML = host;
    } else {
        // Omit logo if not present
        logoElement.remove();
    }
    return templateClone;
};

export const appendRandom = async (parent: HTMLElement, ...elements: DocumentFragment[]) => {
    for (const element of elements) {
        const children = parent.children;
        const randomIndex = Math.floor(Math.random() * (children.length + 1));

        if (randomIndex === children.length) {
            // If the random index is equal to the length, append at the end
            parent.append(element);
        } else {
            // Otherwise, insert at the randomly determined index
            parent.insertBefore(element, children[randomIndex]);
        }
    }
};

export const loadProjects = async () => {
    // Create project-gallery loader
    const loader = document.createElement('span');
    loader.classList.add('loader');
    document.getElementById('projects')?.append(loader);

    // Load items into gallery
    const gallery = document.getElementById('project-gallery')!;
    const template = (document.getElementById('project-template') as HTMLTemplateElement)!;

    const githubPage = await fetchData('/proxy/github');
    const githubProjects = scrapeGithub(githubPage).map((p) => projectIntoTemplate(p, template));
    await appendRandom(gallery, ...githubProjects);

    const bggPage = await fetchData('/proxy/boardgamegeek');
    const bggRawProjects = scrapeBgg(bggPage);

    // Get higher resolution image from bgg xmlapi
    for (const project of bggRawProjects) {
        const id = project.url
            ?.toString()
            ?.split('/')
            .find((v) => v.match(/\d+/g));
        const gameXml = await fetchData(`/xmlapi/boardgamegeek/${id}`, 'text/xml');
        upgradeBggImage(project, gameXml);
    }

    const bggProjects = bggRawProjects.map((p) => projectIntoTemplate(p, template));

    await appendRandom(gallery, ...bggProjects);

    const cults3dPage = await fetchData('/proxy/cults3d');
    const cults3dProjects = scrapeCults3d(cults3dPage).map((p) => projectIntoTemplate(p, template));
    await appendRandom(gallery, ...cults3dProjects);

    // remove loader
    loader.remove();
};
