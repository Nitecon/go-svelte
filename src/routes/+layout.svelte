<script>
    import Header from './Header.svelte';
    import '../styles.css';
    import {userEvents} from "../events/main.js";
    import {onMount} from 'svelte';
    import {fly} from 'svelte/transition';
    import {backOut} from "svelte/easing";
</script>

<Header/>
<!--<CookieBanner/>
<Growl/>
<InfoModal/>-->

<div id="app-container" class="app-container">
    <div class="page" in:fly={{y: 100,delay: 250,easing: backOut}}>
        <slot />
    </div>
</div>

<style>
    .app-container {
        min-height: 50vh;
    }

</style>