<script>
    import { fly } from 'svelte/transition';
    import { backOut } from 'svelte/easing';

    let showBanner = true;

    function closeBanner() {
        showBanner = false;
    }
</script>

{#if showBanner}
    <div class="cookie-banner" in:fly={{ y: -100, delay: 250, easing: backOut }}>
        <p>
            We only use essential cookies for providing a secure and personalized experience to our logged-in users. No third-party tracking cookies, such as Google Analytics, are used on this website. By continuing to use the site, you agree to our use of these essential cookies.
        </p>
        <button on:click={closeBanner} class="close-banner-btn">Close</button>
    </div>
{/if}

<style>
    .cookie-banner {
        position: fixed;
        bottom: 0;
        left: 0;
        right: 0;
        background-color: #f1f1f1;
        color: #333;
        padding: 16px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        font-size: 14px;
        z-index: 1000;
    }

    .close-banner-btn {
        background-color: #333;
        color: #fff;
        border: none;
        padding: 8px 16px;
        font-size: 14px;
        cursor: pointer;
    }
</style>
