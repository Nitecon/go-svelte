<script>
    import Navi from '../components/Navigation.svelte';
    import CookieBanner from '../components/CookieBanner.svelte';
    import Growl from '../components/Growl.svelte';
    import '../app.css';
    import {fly} from 'svelte/transition';
    import {backOut} from "svelte/easing";
</script>

<Navi/>
<CookieBanner/>
<Growl/>

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