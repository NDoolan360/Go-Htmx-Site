# My Personal Website

![GitHub Workflow Status (Continuous Integration)](https://img.shields.io/github/actions/workflow/status/NDoolan360/NDoolan360-Site/ci.yml?logo=github&logoColor=white&label=CI)
![Github-Netlify deployments](https://img.shields.io/github/deployments/NDoolan360/NDoolan360-Site/production?logo=netlify&label=CD)

The source code for my personal website, built with Go, Htmx and Hyperscript, and deployed with Netlify.

## Tools

### Built with

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=FFF)
![Htmx](https://img.shields.io/badge/Htmx-333?logo=htmx&logoColor=FFF)
![TailwindCSS](https://img.shields.io/badge/Tailwind%20CSS-0f172a?logo=tailwindcss&logoColor=06B6D4)

### Deployed with

![Netlify](https://img.shields.io/badge/Netlify-FFF?logo=netlify&logoColor=004846&link=https%3A%2F%2Fnetlify.com)

## External

### Components

-   [Htmx](/public/scripts/htmx.min.js) - [Big Sky Software](https://github.com/bigskysoftware/htmx)
-   [Hyperscript](/public/scripts/hypersript.min.js) - [Big Sky Software](https://github.com/bigskysoftware/_hyperscript)
-   [Markdown rendering](/public/scripts/zero-md.min.js) - [\<zero-md\>](https://github.com/zerodevx/zero-md)
-   [Theme Switch](/templates/theme_switch.svg) - [web.dev](https://web.dev/patterns/theming/theme-switch)
-   [Github Language Colors](/api/projects.go) - [github-langs-go](https://github.com/NDoolan360/github-langs-go)

### Fonts

-   [JosefinSans.woff2](/public/fonts) - [Josefin Sans Project](https://github.com/ThomasJockin/JosefinSansFont-master) &copy; 2010
-   [FragmentMono.woff2](/public/fonts) - [Wei Huang](https://weiweihuanghuang.github.io/), [URW Design Studio](https://www.urwtype.com) &copy; 2022

### Inspiration

-   [brittanychiang.com](https://brittanychiang.com) - [Brittany Chiang](https://github.com/bchiang7)
-   [jakelazaroff.com](https://jakelazaroff.com) - [Jake Lazaroff](https://github.com/jakelazaroff)
-   [Realtime Colors](https://www.realtimecolors.com) - [Juxtopossed](https://github.com/juxtopposed)

## Usage

### Environment Variables

-   Cults3D project rendering requires access to the [Cults3D graphql API](https://cults3d.com/en/pages/graphql).
    -   `CULTS3D_USERNAME`: _string_ - Username of a Cults3D account.
    -   `CULTS3D_API_KEY`: _string_ - API key from Cults3D account matching username.
