<script lang="ts">
    import Placeholder from '/images/placeholder.png';
    import PlaceholderPdf from '/documents/placeholder.pdf';
    import Button from './components/Button.svelte';
    import Card from './components/Card.svelte';
    import Logo from './components/Logo.svelte';
</script>

<section
    id="hero"
    class="flex column"
    style:background-image="url({Placeholder})"
>
    <h1>Hello, I'm Nathan Doolan</h1>
    <p>Software Engineer</p>
    <Button link="#projects">View My Work</Button>
</section>

<section id="bio" class="flex">
    <div class="flex column">
        <img alt="Profile" src={Placeholder} />
        <Button link={PlaceholderPdf} target="_blank" rel="noreferrer">
            Download Resume
        </Button>
    </div>
    <p>
        Hi there! I'm a software engineer specialized in frontend and backend
        development for complex scalable web apps. I write about web development
        on my blog and regularly speak at various web conferences and meetups.
    </p>
</section>

<section id="interests" class="flex column">
    <h2>Interests</h2>
    <div style="flex">
        <Button chip>Web Development</Button>
        <Button chip>3D Printing</Button>
        <Button chip>Board Games</Button>
        <Button chip>Tabletop Games</Button>
        <Button chip>Hiking</Button>
        <Button chip>Cooking</Button>
    </div>
</section>

<section id="projects" class="flex column">
    <h2>Projects</h2>
    <div class="flex">
        <Card image={Placeholder} imageAlt="Project 1">
            <h2 slot="title">Project 1</h2>
            <p slot="description">This is a description of Project 1.</p>
        </Card>
        <Card image={Placeholder} imageAlt="Project 2">
            <h2 slot="title">Project 2</h2>
            <p slot="description">This is a description of Project 2.</p>
        </Card>
        <Card image={Placeholder} imageAlt="Project 3">
            <h2 slot="title">Project 3</h2>
            <p slot="description">This is a description of Project 3.</p>
        </Card>
    </div>
</section>

<footer id="links">
    <div class="flex container">
        <Logo name="Github" />
        <Logo name="LinkedIn" />
        <Logo name="BoardGameGeek" />
        <Logo name="Cults3d" />
    </div>
</footer>

<style>
    * {
        --primary: teal;
        --on-primary: #222;
        --secondary: darkslategrey;
        --on-secondary: lightslategrey;
        --background: #1c1b22;
        --on-background: white;
    }

    section {
        width: 100%;
        margin: 0 auto;
        padding: 3rem 1rem;
    }

    .container {
        width: 100%;
        margin: 0 auto;
        padding: 1rem;
    }

    .flex {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    .column {
        flex-direction: column;
    }

    #hero {
        height: 100vh;
        background-position: center;
        background-size: cover;
    }

    h1 {
        font-size: 2.25rem;
        font-weight: bold;
        color: var(--primary);
        margin-top: 1rem;
    }

    #bio > div > img {
        aspect-ratio: 1;
        object-fit: cover;
        border-radius: 100rem;
    }

    #bio > p {
        max-width: 50vw;
    }

    #interests {
        color: var(--primary);
        background-color: var(--on-primary);
    }

    footer {
        width: 100%;
        padding: 1rem 0;
    }

    #links {
        background-color: var(--primary);
    }
</style>
