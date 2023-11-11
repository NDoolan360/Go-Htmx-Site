<script lang="ts">
    export let link: string | undefined = undefined;
    export let chip = false;
</script>

<a class="btn round" class:chip {...$$restProps} href={link}>
    <slot />
</a>

<style>
    a {
        display: inline-flex;
        align-items: center;
        justify-content: center;
        padding: 0.75rem 1.25rem;
        border: 1px solid transparent;
        border-radius: 100rem;
        font-size: 1rem;
        color: var(--on-primary);
        background-color: var(--primary);
        margin-top: 1rem;
    }

    a:not(.chip):is(:hover, :focus-visible) {
        color: var(--on-secondary);
        background-color: var(--secondary);
    }

    .chip {
        padding: 0.5rem 0.75rem;
        font-size: 0.875rem;
        background-color: var(--on-secondary);
        color: var(--secondary);
    }
</style>
