# My Personal Website

![GitHub Workflow Status (Continuous Integration)][ci-badge]
![Github-Netlify deployments][cd-badge]

The source code for my personal website, built with Go, Htmx and Hyperscript, and deployed with Netlify.

## Tools

### Built with

![Go][go-badge]
![Htmx][htmx-badge]
![TailwindCSS][tailwind-badge]

### Deployed with

![Netlify][netlify-badge]

## External

### Components

-   [Htmx][htmx] - [Big Sky Software][big-sky-software]
-   [Hyperscript][hyperscript] - [Big Sky Software][big-sky-software]
-   [\<zero-md\>][zero-md] - [zerodevx][zerodevx]
-   [Theme Switch][theme-switch] - [web.dev][web-dev]

### Fonts

-   JosefinSans.woff2 - [Josefin Sans Project]() &copy; 2010
-   FragmentMono.woff2 - [Wei Huang][wei-huang], [URW Type][urw-type] &copy; 2022

### Inspiration

-   [brittanychiang.com][brittanychiang.com] - [Brittany Chiang][brittany-chiang]
-   [jakelazaroff.com][jakelazaroff.com] - [Jake Lazaroff][jake-lazaroff]
-   [Realtime Colors][realtime-colors] - [Juxtopossed][juxtopossed]

## Usage

### Requirements

-   Go `>= 1.22.5`

### Dev dependencies

- Parallel: `brew install parallel`
- Watchexec: `brew install watchexec`

### Environment Variables

-   Cults3D project rendering requires access to the [Cults3D graphql API][cults-graphql].
    -   `CULTS3D_USERNAME`: _string_ - Username of a Cults3D account.
    -   `CULTS3D_API_KEY`: _string_ - API key from Cults3D account matching username.

[ci-badge]: https://img.shields.io/github/actions/workflow/status/NDoolan360/NDoolan360-Site/ci.yml?logo=github&logoColor=white&label=CI
[cd-badge]: https://img.shields.io/github/deployments/NDoolan360/NDoolan360-Site/production?logo=netlify&label=CD
[go-badge]: https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=FFF
[htmx-badge]: https://img.shields.io/badge/Htmx-333?logo=htmx&logoColor=FFF
[tailwind-badge]: https://img.shields.io/badge/Tailwind%20CSS-0f172a?logo=tailwindcss&logoColor=06B6D4
[netlify-badge]: https://img.shields.io/badge/Netlify-FFF?logo=netlify&logoColor=004846&link=https%3A%2F%2Fnetlify.com
[htmx]: https://github.com/bigskysoftware/htmx
[hyperscript]: https://github.com/bigskysoftware/_hyperscript
[big-sky-software]: https://github.com/bigskysoftware
[zero-md]: https://github.com/zerodevx/zero-md
[zerodevx]: https://github.com/zerodevx
[theme-switch]: https://web.dev/patterns/theming/theme-switch
[web-dev]: https://web.dev
[josefin-sans]: https://github.com/ThomasJockin/JosefinSansFont-master
[wei-huang]: https://weiweihuanghuang.github.io/
[urw-type]: https://www.urwtype.com
[brittanychiang.com]: https://brittanychiang.com
[brittany-chiang]: https://github.com/bchiang7
[jakelazaroff.com]: https://jakelazaroff.com
[jake-lazaroff]: https://github.com/jakelazaroff
[realtime-colors]: https://www.realtimecolors.com
[juxtopossed]: https://github.com/juxtopposed
[cults-graphql]: https://cults3d.com/en/pages/graphql
