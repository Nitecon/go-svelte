<script>
    import {onDestroy, onMount} from "svelte";

    import {userEvents} from "../events/main.js";

    let loggedIn = false;
    let userData = null;
    let isMobile = false;
    let name = "unknown"
    onMount(() => {
        const handleResize = () => {
            isMobile = window.innerWidth <= 1350;
        };

        window.addEventListener('resize', handleResize);
        handleResize(); // Call once on mount to initialize the value

        return () => {
            window.removeEventListener('resize', handleResize);
        };
    });
    let unsubUser = userEvents.subscribe((data) => {
        loggedIn = data.isLoggedIn;
        userData = data.userData;
        if (userData.username === null || userData.username === undefined || userData.username === "") {
            name = userData.email;
        } else {
            name = userData.username;
        }
    });
    let menuOpen = false;
    let showProfile = true;
    const toggleMenu = () => {
        menuOpen = !menuOpen;
        showProfile = !menuOpen;
    };
    onDestroy(() => {
        unsubUser();
    });
</script>

<nav>
    {#if isMobile}
        <div class="mobile-menu">
            <div class="hamburger-container">
                <div class="hamburger action" on:click="{toggleMenu}" on:keypress={toggleMenu}>&#9776;</div>
                <div class:open={menuOpen} class="mobile-link-box">
                    <ul class="mobile-list">
                        <li><a href="/">Home</a></li>
                        <li><a href="/about">About Us</a></li>
                        {#if loggedIn}
                            <li><a href="/api/auth/logout">Log Out</a></li>
                        {/if}
                    </ul>
                </div>
            </div>
            <ul class:showLogin={showProfile} class="mobile-list login-box ra">
                {#if loggedIn}
                    <li><a href="/profile">{name}</a></li>
                {:else}
                    <li><a href="/login">Login</a></li>
                {/if}
            </ul>
        </div>
    {:else }

        <div class="menu ">
            <ul class="main-links">
                <li><a href="/">Home</a></li>
                <li><a href="/about">About Us</a></li>
            </ul>
            <ul class="main-links ra">
                {#if loggedIn}
                    <li><a href="/api/auth/logout">Logout</a></li>
                    <li><a href="/profile">{name}</a></li>
                {:else}
                    <li><a href="/login">Login</a></li>
                {/if}
            </ul>
        </div>
    {/if}
</nav>

<style>
    .login-box {
        display: none;
        list-style: none;
        margin: 0;
        padding: 0;
        background-color: transparent;
        color: var(--color-text);
        text-decoration: none;
        border: 0;
        cursor: pointer;
        justify-content: right;
    }

    .login-box li {
        margin: 5px;
        border: 0;
    }

    .login-box.showLogin {
        display: flex;
        flex-flow: row;
        justify-content: right;
    }

    .hamburger-container {
        display: flex;
        align-items: center;
    }
</style>