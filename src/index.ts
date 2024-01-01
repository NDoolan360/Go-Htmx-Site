import { inject } from "@vercel/analytics";
import { loadProjects } from "./projects";
import { replaceWithCurrentYear } from "./utils";

// Insert Analytics
inject();

// Update copyright year
const copyright = document.getElementById("copyright");
if (copyright) {
    copyright.innerHTML = replaceWithCurrentYear(copyright.innerHTML, "2023");
}

loadProjects();
