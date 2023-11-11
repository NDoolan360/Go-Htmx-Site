<script lang="ts">
    export let image: string;
    export let imageAlt: string;
</script>

<div>
    <img src={image} alt={imageAlt} />
    <hgroup>
        <slot name="title" />
        <slot name="description" />
    </hgroup>
</div>

<style>
    div {
        border-radius: 0.375rem;
        overflow: hidden;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        background-color: black;
        width: 80vw;
        max-width: 18rem;
    }

    img {
        padding: 0.5rem;
        width: 100%;
        max-height: 12.5rem;
        object-fit: cover;
    }

    hgroup {
        padding: 1rem;
        width: 100%;
    }

    slot[name='title'],
    slot[name='description'] {
        width: 100%;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }

    slot[name='title'] {
        font-size: 1.5rem;
        font-weight: bold;
        margin-bottom: 1rem;
    }

    slot[name='description'] {
        font-size: 1rem;
    }
</style>
