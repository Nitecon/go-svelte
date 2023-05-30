<script>
    import {validators, Hint, useForm} from "svelte-use-form";
    import {fade} from "svelte/transition";
    import {ucfirst} from "../tools/strings.js";
    import {goto} from "$app/navigation";

    const form = useForm({});
    export let config = {};
    export let submissionStatus = "";
    export let submissionResult = null;

    export let numFailures = 0;

    let submitButtonClass = "";

    let file = null;
    let fileData = null;

    // whenever fileData changes, update the form
    function handleFileChange(e) {
        file = e.target.files[0];
        const reader = new FileReader();
        reader.onload = function () {
            config.fields.content.value = reader.result;
        };
        reader.readAsText(file);
    }

    async function handleSubmit(event) {
        event.preventDefault();
        if ($form.valid) {
            submissionStatus = "submitting";
            try {
                const jsonData = JSON.stringify($form.values);
                await fetch(config.location, {
                    method: config.method,
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: jsonData,
                }).then((response) => {
                    if (!response.ok) {
                        submissionStatus = "error";
                        submissionResult = `Error: ${response.status} - ${response.statusText}`;
                        numFailures++;
                    }
                    return response.json();
                })
                    .then((data) => {
                        if (!!config.successCallback && typeof config.successCallback === "function") {
                            config.successCallback();
                        }
                        if (config.redirLocation && config.redirLocation.length > 0){
                            goto(config.redirLocation); // Use $app/navigation instead of window.location.href
                        }
                        if (data.result) {
                            submissionStatus = "success";
                            submissionResult = data.result;
                        } else {
                            submissionStatus = "success";
                            submissionResult = data;
                        }
                    })
                    .catch((error) => {
                        numFailures++;
                        submissionStatus = "error";
                        submissionResult = error;
                    });
            } catch (error) {
                numFailures++;
                submissionStatus = "error";
                submissionResult = ucfirst(error);
            }
        }
    }
</script>
{#if !!config}

    {#if submissionStatus === 'loading'}
        <div class="form-status" in:fade="{{duration: 500}}">
            <div class="status-container">
                <h4>{submissionResult}</h4>
                <div class="loader-container">
                    <div class="loader"></div>
                </div>
            </div>
        </div>

    {:else if submissionStatus === 'success'}
        {#if submissionResult && submissionResult.message}
            <div class="status-container">
                <div class="form-status" in:fade="{{duration: 500}}">
                    <h4>{submissionResult.message}</h4>
                </div>
            </div>
        {/if}
    {:else}
        {#if submissionResult && submissionResult.message}
            <div class="errorBorder">
                <div class="form-status" in:fade="{{duration: 500}}">
                    <h4>{submissionResult.message}</h4>
                </div>
            </div>
        {/if}
        <form use:form on:submit|preventDefault="{handleSubmit}" in:fade="{{duration: 500}}">
            {#each Object.entries(config.fields) as [fieldName, field]}
                {#if field.type === "hidden"}
                    <input type="hidden" id="{fieldName}" name="{fieldName}" bind:value={field.value} use:validators={field.validators}/>
                {:else}
                    <div class:input-group-top={field.type ==="textarea"} class:input-group={field.type !== "textarea"}>
                        <label for="{fieldName}">{field.label}</label>
                        {#if field.type === "textarea"}
                            <textarea id="{fieldName}" name="{fieldName}" bind:value={field.value} use:validators={field.validators}></textarea>
                        {:else if field.type === "file"}
                            <input type="file" class="form-control-file file-input" id="file" on:change="{handleFileChange}"/>
                        {:else if field.type === "password"}
                            <input type="password" id="{fieldName}" name="{fieldName}" bind:value={field.value} use:validators={field.validators}/>
                        {:else if field.type === "select"}
                            <select id="{fieldName}" name="{fieldName}" bind:value={field.value} use:validators={field.validators}>
                                {#each field.options as option}
                                    <option value="{option.value}">{option.label}</option>
                                {/each}
                            </select>
                        {:else if field.type === "checkbox"}
                            <input type="checkbox" id="{fieldName}" name="{fieldName}" bind:checked={field.value} use:validators={field.validators}/>
                        {:else if field.type === "radio"}
                            <input type="radio" id="{fieldName}" name="{fieldName}" bind:group={field.value} use:validators={field.validators}/>
                        {:else}
                            <input type="text" id="{fieldName}" name="{fieldName}" bind:value={field.value} use:validators={field.validators}/>
                        {/if}
                    </div>
                    <Hint for="{fieldName}" on="minLength" let:value>
                        {field.label} must be at least {value} characters.
                    </Hint>
                    <Hint for="{fieldName}" on="email">
                        Entry must be a valid email address.
                    </Hint>
                    <Hint for="{fieldName}" on="matchField" let:value>
                        {value}
                    </Hint>
                    <Hint for="{fieldName}" on="minUpperCase" let:value>{value}</Hint>
                    <Hint for="{fieldName}" on="minLowerCase" let:value>{value}</Hint>
                    <Hint for="{fieldName}" on="minNumbers" let:value>{value}</Hint>
                    <Hint for="{fieldName}" on="minSymbols" let:value>{value}</Hint>
                {/if}

            {/each}
            <hr/>
            {#if config.submitButtonClass}
                <button type="submit" class="action {config.submitButtonClass}">{config.submitButtonText}</button>
            {:else}
                <button type="submit" class="action btn">{config.submitButtonText}</button>
            {/if}

        </form>
    {/if}
{/if}

<style>
    .form-status {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 10px;
    }

    .status-container {
        display: block;
        justify-content: center;
        align-items: center;
        margin-top: 20px;
        text-align: center;
        border: 1px solid #ccc;
        padding: 10px;
        border-radius: 5px;
        width: 100%;
        max-width: 400px;
    }

    h4 {
        font-size: 18px;
        font-weight: 500;
        line-height: 24px;
        text-align: center;
        margin-top: 5px;
        color: #18d4ff;
    }

</style>