<script lang="ts">
    export let name: string;
</script>

<a href="https://www.{name.toLowerCase()}.com/">
    <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill={!['Cults3d', 'BoardGameGeek'].includes(name)
            ? 'none'
            : 'currentColor'}
        stroke={!['Cults3d', 'BoardGameGeek'].includes(name)
            ? 'currentColor'
            : 'none'}
        stroke-width={!['Cults3d', 'BoardGameGeek'].includes(name) ? '2' : '0'}
        stroke-linecap="round"
        stroke-linejoin="round"
        style="height: 1.5rem; width: 1.5rem;"
    >
        {#if name == 'Github'}
            <path
                d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"
            />
            <path d="M9 18c-4.51 2-5-2-7-2" />
        {:else if name == 'Cults3d'}
            <path
                d="M5.09 8.824C3.23 10.114 2 12.314 2 14.81 2 18.775 5.1 22 8.91 22a6.64 6.64 0 0 0 3.09-.762h-.002l-.012-.007A6.66 6.66 0 0 0 15.09 22C18.9 22 22 18.775 22 14.81c0-2.498-1.231-4.701-3.094-5.99C18.716 5.026 15.69 2 11.999 2 8.312 2 5.29 5.017 5.093 8.802zm1.577-.829C7.2 5.44 9.387 3.52 12 3.52c3.008 0 5.456 2.547 5.456 5.677a5.77 5.77 0 0 1-1.769 4.18l-.76-.454-1.464.873 2.35 1.402a7.231 7.231 0 0 0 2.963-4.562 5.762 5.762 0 0 1 1.765 4.174c0 3.128-2.445 5.672-5.45 5.672-2.61 0-4.8-1.918-5.331-4.47l.777-.467V13.8l-2.35 1.405a7.292 7.292 0 0 0 2.359 5.017l-.001-.002a5.22 5.22 0 0 1-1.635.262c-3.005 0-5.45-2.545-5.45-5.672s2.445-5.672 5.45-5.672c.57 0 1.12.092 1.637.262l.003.92 1.458.867V8.382a6.63 6.63 0 0 0-3.094-.762 6.65 6.65 0 0 0-2.249.391z"
            />
        {:else if name == 'LinkedIn'}
            <path
                d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"
            />
            <rect width="4" height="12" x="2" y="9" />
            <circle cx="4" cy="4" r="2" />
        {:else if name == 'BoardGameGeek'}
            <path
                d="M 16.870973,5.6604271 19.008331,1 5.1505932,6.1006591 5.9080939,12.209966 4.628579,13.443659 8.443343,23 l 8.109201,-2.987661 2.818877,-6.617401 -1.210493,-1.166992 0.892644,-7.1620351 z"
            />
        {/if}
    </svg>
    <span class="sr-only">{name}</span>
</a>

<style>
    a {
        color: var(--on-primary);
    }

    a:is(:hover, :focus-visible) {
        color: var(--on-secondary);
    }

    .sr-only {
        position: absolute;
        left: -10000px;
        top: auto;
        width: 1px;
        height: 1px;
        overflow: hidden;
    }
</style>
