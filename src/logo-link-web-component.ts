class LogoLink extends HTMLAnchorElement {
    constructor() {
        super();
    }

    connectedCallback() {
        const id = this.innerHTML.toLowerCase().replace(/ /g, '');
        const height = this.getAttribute('height');
        const width = this.getAttribute('width');

        // Create label and add to parent element
        const label = document.createElement('label');
        label.classList.add('sr-only');
        label.setAttribute('for', id);
        label.innerHTML = this.innerHTML;
        this.parentElement?.insertBefore(label, this);

        // Add logo
        this.setAttribute('id', id);
        this.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg"
                            height="${height ?? '28px'}" width="${width ?? '28px'}">
                                <use href="/images/logos.svg#${id}" />
                            </svg>`;
    }
}

customElements.define('logo-link', LogoLink, { extends: 'a' });
