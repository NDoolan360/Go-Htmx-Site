export class Logo extends HTMLSpanElement {
    connectedCallback() {
        const id = this.ariaLabel?.toLowerCase().replaceAll(/\s/g, "");

        if (id && this.ariaLabel) {
            this.classList.add(id);
            this.classList.add(`text-${id}`);
            this.ariaLabel = `${this.ariaLabel} Logo`;

            const use = document.createElementNS("http://www.w3.org/2000/svg", "use");
            use.setAttribute("href", `/images/logos.svg#${id}`);

            const svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
            svg.appendChild(use);

            this.innerText = "";
            this.appendChild(svg);
        }
    }
}

customElements.define("custom-logo", Logo, { extends: "span" });
