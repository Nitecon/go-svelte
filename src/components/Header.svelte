<script>
    export let type = "h1";
    export let title = "Home";
    export let image = "/favicon.png";
</script>

{#if type === "h1" }
    <h1 class="header-container">
        <img src="{image}" alt="{title} Logo"/>{title}
    </h1>
{:else if type === "h2" }
    <h2 class="header-container">
        <img src="{image}" alt="{title} Logo"/>{title}
    </h2>
{:else if type === "h3" }
    <h3 class="header-container">
        <img src="{image}" alt="{title} Logo"/>{title}
    </h3>
{:else if type === "h4" }
    <h4 class="header-container">
        <img src="{image}" alt="{title} Logo"/>{title}
    </h4>
{:else if type === "h5" }
    <h5 class="header-container">
        <img src="{image}" alt="{title} Logo"/>{title}
    </h5>
{:else if type === "h6" }
    <h6 class="header-container">
        <img src="{image}" alt="{title} Logo"/>{title}
    </h6>
{/if}