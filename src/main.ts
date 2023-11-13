import sections, { type SectionValue } from './sections';
import socials, { type SocialLink } from './socials';

for (const [key, value] of Object.entries<SocialLink>(socials)) {
  addSocialLink(key, value);
}

for (const [key, values] of Object.entries<SectionValue[]>(sections)) {
  addRelativeLink(key);
  addSection(key, values);
}


function addSocialLink(name: string, info: SocialLink) {
  const listItem = document.createElement("li");
  listItem.innerHTML = `
    <label class="sr-only" for="${info.title}-link"> ${name} </label>
    <a id="${name}-link" href="${info.link}" target="_blank" rel="noreferrer">
      <svg xmlns="http://www.w3.org/2000/svg" height="32px" width="32px">
        <use href="/images/logos.svg#${name}" />
      </svg>
    </a>
  `;
  document.getElementById('social-links')?.append(listItem);
}

function addRelativeLink(title: string) {
  const listItem = document.createElement("li");
  listItem.innerHTML = `<a href="#${title}">${title}</a>`;
  document.getElementById('relative-links')?.append(listItem);
}

function addSection(sectionName: string, values: SectionValue[]) {
  let section = document.createElement("section");
  section.id = sectionName;
  section.innerHTML = `
  <hgroup class="separator">
    <h2> ${sectionName} </h2>
    <hr />
  </hgroup>`;
  let contentWrapper = document.createElement("div");
  contentWrapper.append(...values.map((value) => {
    switch (value.type) {
      case "chip":
        return createChip(value.value);
      case "card":
        return createCard(value.value);
    }
    return "";
  }))
  section.append(contentWrapper);
  document.getElementById('main')?.append(section);
}

function createChip(text: string): HTMLParagraphElement {
  const chip = document.createElement("p");
  chip.classList.add('chip');
  chip.innerHTML = text;
  return chip;
}

function createCard(url: string): HTMLElement {
  const chip = document.createElement("a");
  chip.classList.add('card');
  chip.href = url;
  chip.innerHTML = url;
  return chip;
}
