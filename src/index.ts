import { inject } from "@vercel/analytics";
import "./index.css";
import {
    GithubRepo,
    githubRepoToProject,
    projectIntoTemplate,
    scrapeBgg,
    scrapeCults3d,
    upgradeBggData,
} from "./projects";
import { fetchData, fetchJson, replaceWithCurrentYear } from "./utils";

// Insert Analytics
inject();

// Update copyright year
const copyright = document.getElementById("copyright");
if (copyright) {
    copyright.innerHTML = replaceWithCurrentYear(copyright.innerHTML, "2023");
}

const loadProjects = async () => {
    // Create project-gallery loader
    const loader = document.createElement("span");
    loader.classList.add("loader");
    document.getElementById("projects")?.append(loader);

    // Load items into gallery
    const gallery = document.getElementById("project-gallery");
    const template = document.getElementById("project-template") as HTMLTemplateElement | undefined;

    if (gallery && template) {
        const githubPage = await fetchJson<GithubRepo[]>("/proxy/api/github");
        const githubProjects = githubRepoToProject(githubPage).map((p) =>
            projectIntoTemplate(p, template),
        );
        gallery.append(...githubProjects);

        const bggPage = await fetchData("/proxy/boardgamegeek");
        const bggRawProjects = scrapeBgg(bggPage);

        // Get higher resolution image from bgg xmlapi
        for (const project of bggRawProjects) {
            const id = project.url
                ?.toString()
                ?.split("/")
                .find((v) => v.match(/\d+/g));
            const gameXml = await fetchData(`/proxy/xmlapi/boardgamegeek/${id}`, "text/xml");
            upgradeBggData(project, gameXml);
        }

        const bggProjects = bggRawProjects.map((p) => projectIntoTemplate(p, template));

        gallery.append(...bggProjects);

        const cults3dPage = await fetchData("/proxy/cults3d");
        const cults3dProjects = scrapeCults3d(cults3dPage).map((p) =>
            projectIntoTemplate(p, template),
        );
        gallery.append(...cults3dProjects);
    }

    // remove project-gallery loader
    loader.remove();
};

loadProjects();
