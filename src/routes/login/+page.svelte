<script>
import Header from "../../components/Header.svelte";
import {email} from "svelte-use-form";
import DynamicForm from "../../components/DynamicForm.svelte";
import {onDestroy, onMount} from "svelte";
import {userEvents} from "../../events/main.js";
import {goto} from "$app/navigation";
import { page } from '$app/stores';
let userInfo;
let unsubUser = userEvents.subscribe((data) => {
    userInfo = data;
    if (!!userInfo && userInfo.isLoggedIn) {
        goto('/#profile');
    }
});
onDestroy(() => {
    unsubUser();
});

let formStatus = null;
let formResult = null;
let formConfig = null;
let failures = 0;
onMount(() => {
    formConfig = {
        method: 'POST',
        submitButtonText: 'Log In',
        submitButtonClass: 'login-button',
        location: '/api/auth/login',
        redirLocation: '/#profile',
        fields: {
            email: {label: 'Email', type: 'text', value: '', validators: [email]},
            password: {label: 'Password', type: 'password', value: '', validators: []},
        },
    };
});


export let google_login = "/api/auth/login/google"
//export let github_login = "/api/auth/login/github"
//export let twitter_login = "/api/auth/login/twitter"

let errorMessage = '';
const hash = $page.url.hash.replace(/^#/, '').split('/')[1];

let newAccount = false;
let resetAccount = false;
let forgotAccount = false;
// check if query param created=true
if (hash === '?created=true') {
    newAccount = true;
}
if (hash === '?reset=true') {
    resetAccount = true;
}
if (hash === '?forgot=true') {
    forgotAccount = true;
}
</script>

<svelte:head>
    <title>Login</title>
    <meta name="description" content="Log into the application" />
</svelte:head>

<section class="main-section">
    <main>

        <div class="box">
            <div class="main-form">
                <Header title="Login"></Header>
                {#if newAccount}
                    <div class="create-account">
                        <h4>Your account has been created!</h4>
                        <p>Please check your email on how to validate your account, prior to use.</p>
                    </div>
                {/if}
                {#if forgotAccount}
                    <div class="create-account">
                        <h4>Check your email!</h4>
                        <p>Please check your email, a reset password has been sent to your account.</p>
                    </div>
                {/if}
                {#if resetAccount}
                    <div class="create-account">
                        <h4>Reset Successful</h4>
                        <p>Your reset was successful, please log in with your new password.</p>
                    </div>
                {/if}
                {#if errorMessage}
                    <div class="errorBorder">
                        <p>{errorMessage}</p>
                    </div>
                {/if}
                <div class="login-form">
                    {#if failures >= 3}
                        <div class="create-account">
                            <h4>Too many failed login attempts!</h4>
                            <p>Click <a href="/#forgot-password">here</a> to reset your password.</p>
                        </div>
                    {:else}
                        <DynamicForm bind:numFailures={failures} bind:config={formConfig} bind:submissionStatus="{formStatus}" bind:submissionResult="{formResult}"/>
                        <div class="create-account">
                            <a href="/#create-account">Create Account</a>
                            &nbsp;|&nbsp;
                            <a href="/#forgot-password">Forgot Password</a>
                        </div>
                    {/if}
                </div>

                <div class="social">
                    <a href="{google_login}" class="google-signin">
                        <img src="https://developers.google.com/identity/images/btn_google_signin_light_normal_web.png" alt="Sign in with Google">
                    </a>
                    <!--
                    <hr />
                    <a href="{github_login}" class="github-signin">
                        <img src="https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png" alt="GitHub Logo">
                        Sign in with GitHub
                    </a>
                    <hr />
                    <a href="{twitter_login}" class="twitter-signin">
                        <img src="https://abs.twimg.com/a/1537210120/img/t1/twitter_logo_standard_colors.svg" alt="Twitter Logo">
                        Sign in with Twitter
                    </a>-->
                </div>
            </div>
        </div>
    </main>
</section>

<style>
    .create-account {
        margin-top: 20px;
        text-align: center;
        border: 1px solid #ccc;
        padding: 10px;
        border-radius: 5px;
    }

    h4 {
        font-size: 18px;
        font-weight: 500;
        line-height: 24px;
        text-align: center;
        margin-top: 5px;
        color: #18d4ff;
    }


    .main-form {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
        padding: 10px;
    }

    .login-form {
        width: 100%;
        max-width: 600px;
        padding: 15px;
        margin: 0 auto;
    }

    .create-account {
        margin-top: 20px;
        text-align: center;
    }

    .social {
        margin-top: 10px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
        padding: 10px;
    }

    .google-signin {
        display: inline-block;
        background: none;
        border: none;
        cursor: pointer;
        text-align: center;
        padding: 0;
    }

    .google-signin img {
        width: 191px;
        height: 46px;
        border: none;
    }

</style>