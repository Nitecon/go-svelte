<script>
    import { fade, fly } from 'svelte/transition';
    import {growlEvents} from "../events/main.js";
    import {onDestroy} from "svelte";

    export let message = '';
    export let growlerClass = 'success';

    export let duration = 3000;
    export let y = -50;

    export let visible = false;

    let unsubscribe = growlEvents.subscribe((event) => {
        message = event.message;
        growlerClass = event.type;
        if (event.type === "info"){
            growlerClass = "success";
        }
        visible = event.isShown;
        if (visible){
            setTimeout(close, duration);
        }
    });

    onDestroy(() => {
        unsubscribe();
    });

    function close() {
        visible = false;
    }

    function handleKeydown(event) {
        if (event.key === 'Enter' || event.key === ' ') {
            close();
        }
    }
</script>

<div class="growler {growlerClass}" class:visible on:click={close} in:fly={{ y }} on:keydown={handleKeydown} out:fade>
    {message}
</div>

<style>


    div.visible {
        opacity: 1;
        pointer-events: all;
    }

    .success {
        background-color: #4caf50;
        color: #fff;
    }

    .error {
        background-color: #f44336;
        color: #fff;
    }
</style>
